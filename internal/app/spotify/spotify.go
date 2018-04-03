package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/ThatTomPerson/spotigo/internal/pkg/dhet"
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
}

func New() *Spotify {
	return &Spotify{}
}

func (s *Spotify) Connect(ctx context.Context) error {
	log.Printf("Looking for access points\n")

	aps, err := s.findAPs(ctx)
	if err != nil {
		return fmt.Errorf("couldn't find Spotify APs: %v", err)
	}

	ap := aps.Rand()

	c, err := net.Dial("tcp", ap)
	if err != nil {
		return fmt.Errorf("couldn't connect to Spotify AP %s: %v", ap, err)
	}
	log.Printf("Connected to %s\n", ap)

	c, err = dhet.New(c)
	if err != nil {
		return fmt.Errorf("couldn't excahnge keys with spotify: %v", err)
	}

	return nil
}

func (s *Spotify) findAPs(ctx context.Context) (*ApList, error) {
	resp, err := http.Get("https://apresolve.spotify.com")
	if err != nil {
		// try again with a tcp message to vv
		// resp, err = http.Get("http://ap.spotify.com")
		return nil, err
	}
	// close the body here so we don't forget later
	defer resp.Body.Close()

	aps := &ApList{}

	json.NewDecoder(resp.Body).Decode(aps)

	return aps, nil
}
