package dhet

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"

	"github.com/ThatTomPerson/spotigo/internal/pkg/api/keyexchange"
	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
	"github.com/monnand/dhkx"
)

type Conn struct {
	net.Conn
	sendKey []byte
	recvKey []byte
}

func (c *Conn) Write(b []byte) (int, error) {
	b, err := c.encrypt(b)
	if err != nil {
		return 0, err
	}

	return c.Conn.Write(b)
}

func (c *Conn) Read(b []byte) (int, error) {
	return c.Conn.Read(b)
}

func (c *Conn) encrypt(b []byte) ([]byte, error) {
	return b, nil
}

func New(c net.Conn) (net.Conn, error) {
	return exchange(c)
}

func exchange(c net.Conn) (*Conn, error) {
	log.Printf("Starting handshake\n")

	g, err := dhkx.GetGroup(1)
	if err != nil {
		return nil, fmt.Errorf("couldn't get a dhkx group: %v", err)
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
		return nil, fmt.Errorf("couldn't marshal ClientHello: %v", err)
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
		return nil, fmt.Errorf("couldn't write ClientHello: %v", err)
	}
	spew.Dump(n)

	header = make([]byte, 4)

	n, err = c.Read(header)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response header: %v", err)
	}

	spew.Dump(header)
	// sub 4 because of the length part of the header
	size = int(binary.BigEndian.Uint32(header)) - 4
	// 0, 1, 2,3,4,5
	spew.Dump(size)

	body := make([]byte, size)

	n, err = c.Read(body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response header: %v", err)
	}

	spew.Dump(body)
	res := &keyexchange.APResponseMessage{}
	proto.Unmarshal(body, res)

	spew.Dump(res)

	fmt.Print(proto.MarshalTextString(res))

	serverPubKey := dhkx.NewPublicKey(res.Challenge.LoginCryptoChallenge.DiffieHellman.Gs)

	// Compute the key
	k, _ := g.ComputeKey(serverPubKey, priv)

	// Get the key in the form of []byte
	key := k.Bytes()

	spew.Dump(key)
	packets := serverPubKey.Bytes()
	mac := hmac.New(sha1.New, key)
	data := []byte{}
	for i := 0; i <= 6; i++ {
		mac.Write(packets)
		sum := mac.Sum([]byte{byte(i)})
		mac.Reset()

		spew.Dump(sum)
		data = append(data, sum...)
	}

	challenge := data[:0x14]
	sendKey := data[0x14:0x34]
	recvKey := data[0x34:0x54]

	spew.Dump("Keys")
	spew.Dump(challenge)
	spew.Dump(sendKey)
	spew.Dump(recvKey)

	clientResponse := &keyexchange.ClientResponsePlaintext{
		LoginCryptoResponse: &keyexchange.LoginCryptoResponseUnion{
			DiffieHellman: &keyexchange.LoginCryptoDiffieHellmanResponse{
				Hmac: challenge,
			},
		},
	}

	b, err = proto.Marshal(clientResponse)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal ClientHello: %v", err)
	}
	size = 4 + len(b)

	header = make([]byte, 4)

	binary.BigEndian.PutUint32(header, uint32(size))

	packet = bytes.Join([][]byte{header, b}, []byte{})

	c.Write(packet)

	return &Conn{
		Conn:    c,
		sendKey: sendKey,
		recvKey: recvKey,
	}, errors.New("not implimented")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func nonce(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}
