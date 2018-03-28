package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	_ "gitlab.com/ThatTomPerson/spotify-go/internal/pkg/api/authentication"
	"gitlab.com/ThatTomPerson/spotify-go/internal/pkg/api/keyexchange"
)

type ApList struct {
	ApList []string `json:"ap_list"`
}

func (l *ApList) Rand() string {
	t := len(l.ApList)
	i := rand.Intn(t)
	return l.ApList[i]
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

	ap := aps.Rand()

	c, err := net.Dial("tcp", ap)
	if err != nil {
		return fmt.Errorf("couldn't connect to Spotify AP %s: %v", ap, err)
	}

	var version uint64 = 1

	hello := keyexchange.ClientHello{
		BuildInfo: &keyexchange.BuildInfo{
			Product:      keyexchange.Product_PRODUCT_LIBSPOTIFY.Enum(),
			ProductFlags: []keyexchange.ProductFlags{keyexchange.ProductFlags_PRODUCT_FLAG_DEV_BUILD},
			Platform:     keyexchange.Platform_PLATFORM_OSX_X86_64.Enum(),
			Version:      &version,
		},
		FingerprintsSupported: []keyexchange.Fingerprint{Fingerprint_FINGERPRINT_GRAIN},
	}

	// net.Dial()
	spew.Dump(ap)
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
