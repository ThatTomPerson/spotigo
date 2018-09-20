package shannon

import "ttp.sh/spotigo/internal/pkg/shannon/fast"

type Shannon interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
	Finish([]byte) []byte
	Nonce(int)
}

func New(key []byte) Shannon {
	return fast.New(key)
}
