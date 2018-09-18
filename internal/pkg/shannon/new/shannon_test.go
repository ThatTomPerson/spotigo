package shannon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestKey []byte

func init() {
	TestKey = []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}
}

func TestInitState(t *testing.T) {
	c := &Shannon{}
	c.initState()

	assert.Equal(t, [16]uint32{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}, c.R)
	assert.Equal(t, InitKonst, c.konst)
}

func TestByteToWord(t *testing.T) {
	assert.Equal(t, uint32(0x4030201), byteToWord(TestKey))
	assert.Equal(t, uint32(0x8070605), byteToWord(TestKey[4:]))
	assert.Equal(t, uint32(0x4030201), byteToWord(TestKey[8:]))
	assert.Equal(t, uint32(0x8070605), byteToWord(TestKey[12:]))
}

func TestLoadCycle(t *testing.T) {
	c := &Shannon{}
	c.initState()

	keylen := len(TestKey)

	assert.Equal(t, 32, (keylen & ^3))

	for i := 0; i < (keylen & ^3); i += 4 {
		k := byteToWord(TestKey[i:])
		if i%8 == 0 {
			assert.Equal(t, uint32(0x4030201), k)
		} else {
			assert.Equal(t, uint32(0x8070605), k)
		}

		c.R[KeyP] = c.R[KeyP] ^ k
		assert.Equal(t, uint32(67306360), c.R[KeyP], "i: %d", i)

		c.cycle()
	}
}

func TestLoadKey(t *testing.T) {
	c := &Shannon{}
	c.initState()
	c.loadKey(TestKey)

	assert.Equal(t, [16]uint32{1097138062, 2986595836, 1867816895, 2696735745, 50829331, 4124109027, 1732489910, 3515261258, 2104992913, 95552162, 775257507, 1378231108, 445674099, 2961241956, 1820712884, 1046446811}, c.R, "Register doesn't match expected")
	assert.Equal(t, [16]uint32{12730674, 89, 144, 233, 67306360, 134677607, 67305946, 1087172263, 2577041922, 373320610, 4010104634, 2645023638, 1106318883, 140292518, 1555427227, 1161874008}, c.CRC, "CRC doesn't match expected")
}
