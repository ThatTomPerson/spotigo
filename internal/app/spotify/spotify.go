package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"ttp.sh/spotigo/internal/pkg/dhet"
)

type ApList struct {
	ApList []string `json:"ap_list"`
}

func (l *ApList) Rand() string {
	t := len(l.ApList)
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(t)
	return l.ApList[i]
}

type Spotify struct {
	conn net.Conn
}

func New() *Spotify {
	return &Spotify{}
}

func (s *Spotify) dialAP(ctx context.Context) (net.Conn, error) {
	aps, err := s.findAPs(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't find Spotify APs: %v", err)
	}

	ap := aps.Rand()

	c, err := net.Dial("tcp", ap)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to Spotify AP %s: %v", ap, err)
	}
	logrus.Infof("Connected to %s", ap)

	c, err = dhet.New(c)

	if err != nil {
		return nil, fmt.Errorf("couldn't excahnge keys with spotify: %v", err)
	}

	return c, nil
}

func (s *Spotify) Connect(ctx context.Context) error {
	logrus.Info("Looking for access points")
	c, err := s.dialAP(ctx)
	if err != nil {
		return fmt.Errorf("couldn't find Spotify APs: %v", err)
	}

	s.conn = c
	return nil
}

func (s *Spotify) findAPs(ctx context.Context) (*ApList, error) {
	resp, err := http.Get("https://apresolve.spotify.com")
	if err != nil {
		logrus.Warnf("http.Get https://apresolve.spotify.com : %v", err)
		// try again with a tcp message to vv
		resp, err = http.Get("http://ap.spotify.com")
		if err != nil {
			return nil, err
		}
	}
	// close the body here so we don't forget later
	defer resp.Body.Close()

	aps := &ApList{}

	json.NewDecoder(resp.Body).Decode(aps)

	return aps, nil
}
