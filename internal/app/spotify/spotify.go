package spotify

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	_ "gitlab.com/ThatTomPerson/spotify-go/internal/pkg/api/authentication"
	"gitlab.com/ThatTomPerson/spotify-go/internal/pkg/api/keyexchange"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/monnand/dhkx"
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

	tcpAddr, err := net.ResolveTCPAddr("tcp", ap)
	log.Printf("Connecting to %s (%s:%d)\n", ap, tcpAddr.IP, tcpAddr.Port)

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("couldn't connect to Spotify AP %s: %v", ap, err)
	}

	log.Printf("Connected to %s (%s:%d)\n", ap, tcpAddr.IP, tcpAddr.Port)
	log.Printf("Starting handshake\n")

	g, err := dhkx.GetGroup(1)
	if err != nil {
		return fmt.Errorf("couldn't get a dhkx group: %v", err)
	}

	priv, _ := g.GeneratePrivateKey(nil)
	log.Printf("Generated private key\n")

	// Get the public key from the private key.
	pub := priv.Bytes()

	spew.Dump(pub)

	var version uint64 = 0x10800000000
	var serverKeysKnown uint32 = 1

	hello := &keyexchange.ClientHello{
		BuildInfo: &keyexchange.BuildInfo{
			Product:      keyexchange.Product_PRODUCT_PARTNER.Enum(),
			ProductFlags: []keyexchange.ProductFlags{keyexchange.ProductFlags_PRODUCT_FLAG_DEV_BUILD},
			Platform:     keyexchange.Platform_PLATFORM_LINUX_X86.Enum(),
			Version:      &version,
		},
		CryptosuitesSupported: []keyexchange.Cryptosuite{keyexchange.Cryptosuite_CRYPTO_SUITE_SHANNON},
		LoginCryptoHello: &keyexchange.LoginCryptoHelloUnion{
			DiffieHellman: &keyexchange.LoginCryptoDiffieHellmanHello{
				Gc:              pub,
				ServerKeysKnown: &serverKeysKnown,
			},
		},
		ClientNonce: nonce(0x10),
		Padding:     []byte{0x1e},
	}

	// spew.Dump(hello)
	// return nil

	b, err := proto.Marshal(hello)
	if err != nil {
		return fmt.Errorf("couldn't marshal ClientHello: %v", err)
	}

	size := 2 + 4 + len(b)

	header := make([]byte, 6)

	header[0] = 0
	header[1] = 4
	binary.BigEndian.PutUint32(header[2:], uint32(size))

	packet := bytes.Join([][]byte{header, b}, []byte{})

	spew.Dump(packet)

	spew.Dump(size)
	log.Printf("Sending ClientHello")

	n, err := c.Write(packet)
	if err != nil {
		return fmt.Errorf("couldn't write ClientHello: %v", err)
	}
	spew.Dump(n)

	header = make([]byte, 4)

	n, err = c.Read(header)
	if err != nil {
		return fmt.Errorf("couldn't read response header: %v", err)
	}

	spew.Dump(header)

	size = int(binary.BigEndian.Uint32(header))
	// 0, 1, 2,3,4,5
	spew.Dump(size)

	body := make([]byte, size)

	n, err = c.Read(body)
	if err != nil {
		return fmt.Errorf("couldn't read response header: %v", err)
	}

	spew.Dump(body)
	res := &keyexchange.APResponseMessage{}
	proto.Unmarshal(body, res)

	spew.Dump(res)
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
	spew.Dump(aps)
	return aps, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func nonce(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}
