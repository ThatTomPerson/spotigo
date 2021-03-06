package shannon

import (
	"encoding/binary"
	"unsafe"
)

const (
	InitKonst uint32 = 0x6996c53a
	KeyP             = 13
	N                = 16
)

type Shannon struct {
	R     [N]uint32 // Working storage for the shift register */
	CRC   [N]uint32 // Working storage for CRC accumulation */
	initR [N]uint32 // saved register contents */
	konst uint32    // key dependent semi-constant */
	sbuf  uint32    // encryption buffer */
	mbuf  uint32    // partial word MAC buffer */
	nbuf  uint32    // number of part-word stream bits buffered */
}

func New(key []byte) *Shannon {
	c := &Shannon{}

	c.initState()
	c.loadKey(key)
	c.genkonst()
	c.savestate()

	return c
}

func (c *Shannon) initState() {
	c.R[0] = 1
	c.R[1] = 1
	for i := 2; i < N; i++ {
		c.R[i] = c.R[i-1] + c.R[i-2]
	}
	c.konst = InitKonst
}

func splitKey(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}

	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}

	return chunks
}

// sbox1 - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/internal/pkg/shannon/ShannonRef.c:58
func sbox1(w uint32) uint32 {
	w ^= rotl(w, 5) | rotl(w, 7)
	w ^= rotl(w, 19) | rotl(w, 22)
	return w
}

// sbox2 - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/internal/pkg/shannon/ShannonRef.c:66
func sbox2(w uint32) uint32 {
	w ^= rotl(w, 7) | rotl(w, 22)
	w ^= rotl(w, 5) | rotl(w, 19)
	return uint32(w)
}

func rotl(w uint32, x uint32) uint32 {
	return (w << x) | (w >> (32 - x))
}

func rotr(w uint32, x uint32) uint32 {
	return (w >> x) | (w << (32 - x))
}

// Save the current register state
func (c *Shannon) savestate() {
	c.initR = c.R
}

// initialise to previously saved register state
func (c *Shannon) reloadstate() {
	c.R = c.initR
}

// Initialise "konst"
func (c *Shannon) genkonst() {
	c.konst = c.R[0]
}

func (c *Shannon) cycle() {
	t := c.R[12] ^ c.R[13] ^ c.konst
	t = sbox1(t) ^ rotl(c.R[0], 1)
	for i, r := range c.R[1:] {
		c.R[i-1] = r
	}
	c.R[15] = t
	t = sbox2(c.R[2] ^ c.R[15])
	c.R[0] ^= t
	c.sbuf = t ^ c.R[8] ^ c.R[12]
}

// extra nonlinear diffusion of register for key and MAC
func (c *Shannon) diffuse() {
	for i := 0; i < N; i++ {
		c.cycle()
	}
}

func (c *Shannon) nonce(nonce []byte) {
	c.reloadstate()
	c.konst = InitKonst
	c.loadKey(nonce)
	c.genkonst()
	c.nbuf = 0
}

func (c *Shannon) NonceUint32(nonce uint32) {
	var b []byte
	binary.BigEndian.PutUint32(b, nonce)
	c.nonce(b)
}

// copied from ./ref/ShannonRef.go
func byteToWord(key []byte) uint32 {
	i := 0
	return uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(3)]))&uint32((uint32(int32(255))))<<uint64(int32(24)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(2)]))&uint32((uint32(int32(255))))<<uint64(int32(16)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(1)]))&uint32((uint32(int32(255))))<<uint64(int32(8)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(0)]))&uint32((uint32(int32(255))))
	// return uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&b[i]))[:])[int32(3)]))&uint32((uint32(int32(255))))<<uint64(int32(24)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(2)]))&uint32((uint32(int32(255))))<<uint64(int32(16)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(1)]))&uint32((uint32(int32(255))))<<uint64(int32(8)) | uint32(uint8(((*[1000000000]uint8)(unsafe.Pointer(&key[i]))[:])[int32(0)]))&uint32((uint32(int32(255))))
}

func (c *Shannon) loadKey(key []byte) {
	keylen := len(key)
	a := 0x3
	for i := 0; i < (keylen & ^a); i += 4 {
		k := byteToWord(key[i:])
		c.R[KeyP] ^= k
		c.cycle()
	}

	c.R[KeyP] ^= uint32(keylen)
	c.cycle()

	for i, v := range c.R {
		c.CRC[i] = v
	}

	c.diffuse()

	for i, v := range c.CRC {
		c.R[i] ^= v
	}

	// for _, word := range splitKey(key, 4) {
	// 	if len(word) == 4 {
	// 		c.R[KeyP] ^= binary.LittleEndian.Uint32(word)
	// 	} else {
	// 		for i := len(word); i < 4; i++ {
	// 			word = append(word, byte(0))
	// 		}
	// 		c.R[KeyP] ^= binary.LittleEndian.Uint32(word)
	// 	}

	// 	c.cycle()
	// }
}

// Accumulate a CRC of input words, later to be fed into MAC.
// This is actually 32 parallel CRC-16s, using the IBM CRC-16
// polynomial x^16 + x^15 + x^2 + 1.
func (c *Shannon) crcfunc(i uint32) {
	t := c.CRC[0] ^ c.CRC[2] ^ c.CRC[15] ^ i
	for j := range c.CRC[1:] {
		c.CRC[j-1] = c.CRC[j]
	}
	c.CRC[N-1] = t
}

func (c *Shannon) macfunc(word uint32) {
	c.crcfunc(word)
	c.R[KeyP] ^= word
}

func (c *Shannon) Decrypt(buf []byte) []byte {
	return c.process(buf, c.decryptWord, c.decryptByte)
}

func (c *Shannon) decryptByte(b byte) byte {
	b = ^byte((c.sbuf >> (32 - c.nbuf)) & 0xFF)
	c.mbuf ^= uint32(b) << (32 - c.nbuf)
	return b
}

func (c *Shannon) decryptWord(word uint32) uint32 {
	word ^= c.sbuf
	c.macfunc(word)
	return word
}

func (c *Shannon) Encrypt(buf []byte) []byte {
	return c.process(buf, c.encryptWord, c.encryptByte)
}

func (c *Shannon) encryptByte(b byte) byte {
	c.mbuf ^= uint32(b) << (32 - c.nbuf)
	return b ^ byte((c.sbuf>>(32-c.nbuf))&0xFF)
}

func (c *Shannon) encryptWord(word uint32) uint32 {
	c.macfunc(word)
	return word ^ c.sbuf
}

func (c *Shannon) process(buf []byte, wordFunc func(uint32) uint32, byteFunc func(byte) byte) []byte {
	//todo: Process
	return []byte{}
}

// TODO finish
