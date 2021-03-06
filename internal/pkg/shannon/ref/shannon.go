package ref

//go:generate bash -c "rm shannon_ref.go || true"
//go:generate bash -c "c2go transpile -V -o ./shannon_ref.go -p ref ../../../../docs/shannon/{Shannon.h,ShannonRef.c,ShannonInternal.h}"

type Shannon struct {
	ctx []shn_ctx
}

func New(key []byte) *Shannon {
	c := &Shannon{
		ctx: []shn_ctx{shn_ctx{}},
	}

	c.new(key)

	return c
}

func (c *Shannon) initstate() {

}
func (c *Shannon) new(key []byte) {
	shn_initstate(c.ctx)

	var _key []uint8_t
	for _, v := range key {
		_key = append(_key, uint8_t(v))
	}
	shn_loadkey(c.ctx, _key, int32(len(_key)))
	shn_genkonst(c.ctx)
	shn_savestate(c.ctx)
}

func (c *Shannon) Encrypt(packet []byte) []byte {
	var _packet []uint8_t
	for _, v := range packet {
		_packet = append(_packet, uint8_t(v))
	}

	shn_encrypt(c.ctx, _packet, int32(len(_packet)-1))

	packet = []byte{}
	for _, v := range _packet {
		packet = append(packet, byte(v))
	}

	return packet
}

func (c *Shannon) Decrypt(packet []byte) []byte {
	var _packet []uint8_t
	for _, v := range packet {
		_packet = append(_packet, uint8_t(v))
	}

	shn_decrypt(c.ctx, _packet, int32(len(_packet)-1))

	packet = []byte{}
	for _, v := range _packet {
		packet = append(packet, byte(v))
	}

	return packet
}

func (c *Shannon) Finish(packet []byte) []byte {
	var _packet []uint8_t
	for _, v := range packet {
		_packet = append(_packet, uint8_t(v))
	}

	shn_finish(c.ctx, _packet, int32(len(_packet)-1))

	packet = []byte{}
	for _, v := range _packet {
		packet = append(packet, byte(v))
	}

	return packet
}
