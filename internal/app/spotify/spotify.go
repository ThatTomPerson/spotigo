package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	_ "gitlab.com/ThatTomPerson/spotify-go/internal/pkg/api/authentication"
)

type ApList struct {
	ApList []string `json:"ap_list"`
}

type Spotify struct {
}

func New() *Spotify {
	return &Spotify{}
}

func (s *Spotify) Connect(ctx context.Context) error {
	aps, err := s.findAPs(ctx)
	if err != nil {
		return fmt.Errorf("couldn't find Spotify APs: %v", err)
	}

	// net.Dial()
	spew.Dump(aps)
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
