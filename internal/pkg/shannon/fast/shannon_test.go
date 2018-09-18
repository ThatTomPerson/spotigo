package fast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestKey = "test key 128bits"
const (
	TESTSIZE   = 20
	INPUTSIZE  = 100
	STREAMTEST = 1000000
	ITERATIONS = 999999
)

var _TestKey []uint8_t
var TestOut []uint8_t
var StreamOut []uint8_t

func init() {
	for _, v := range TestKey {
		_TestKey = append(_TestKey, uint8_t(v))
	}
	TestOut = []uint8_t{0x4d, 0x7e, 0xd3, 0x9c, 0xb6, 0x95, 0xd9, 0x6a, 0xcf, 0x52, 0x97, 0x70, 0xec, 0x7d, 0xcc, 0xbe, 0xae, 0x2b, 0x6f, 0x8c}
	StreamOut = []uint8_t{0x27, 0x01, 0x9f, 0xc8, 0x84, 0xbb, 0x09, 0x05, 0xea, 0x08, 0xc9, 0xb5, 0x5f, 0x20, 0x7b, 0x5d, 0x34, 0x80, 0xb4, 0xa3}
}

func new() *Shannon {
	c := &Shannon{
		ctx: []shn_ctx{shn_ctx{}},
	}

	shn_key(c.ctx, _TestKey, int32(len(_TestKey)))
	return c
}

func TestBasic(t *testing.T) {
	c := new()

	assert.NotEqual(t, 0x55bf8df5, c.ctx[0].initR[0], "It is probable that byte ordering is incorrect.")

	testframe := []uint8_t{0x0, 0x0, 0x0, 0x0}
	shn_nonce(c.ctx, testframe, 4)

	testbuf := make([]uint8_t, INPUTSIZE+STREAMTEST)
	shn_stream(c.ctx, testbuf, INPUTSIZE)
	assert.Equal(t, TestOut, testbuf[:TESTSIZE], "one chunk")
}

func TestBasicByteByByte(t *testing.T) {
	c := new()
	testframe := []uint8_t{0x0, 0x0, 0x0, 0x0}
	shn_nonce(c.ctx, testframe, 4)

	testbuf := make([]uint8_t, INPUTSIZE+STREAMTEST)
	for i := 0; i < INPUTSIZE; i++ {
		shn_stream(c.ctx, testbuf[i:], 1)
	}
	assert.Equal(t, TestOut, testbuf[:TESTSIZE], "single bytes")

	shn_stream(c.ctx, testbuf[INPUTSIZE:], STREAMTEST)
	assert.Equal(t, StreamOut, testbuf[STREAMTEST:STREAMTEST+TESTSIZE])
}
