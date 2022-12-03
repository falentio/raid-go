package raid

const (
	// prefer truncate the number to be able use full alphabet prefix
	enc = "012345abcdefghijklmnopqrstuvwxyz"
)

var (
	dec = [256]byte{}
)

func init() {
	for i := range dec {
		dec[i] = 0xff
	}
	for i, c := range []byte(enc) {
		dec[byte(c)] = byte(i)
	}
}

func encodePrefix(src, dst []byte) {
	_ = dst[2]
	_ = src[1]
	_ = enc[0x1f]

	dst[2] = enc[src[1]>>1&0x1f]
	dst[1] = enc[(src[1]>>6|src[0]<<2)&0x1f]
	dst[0] = enc[src[0]>>3&0x1f]
}

func encodeRaid(dst, src []byte) {
	_ = dst[31]
	_ = src[19]
	_ = enc[0x1f]

	src[7] ^= src[17]
	src[6] ^= src[16]

	dst[31] = enc[src[19]&0x1f]
	dst[30] = enc[(src[19]>>5|src[18]<<3)&0x1f]
	dst[29] = enc[src[18]>>2&0x1f]
	dst[28] = enc[(src[18]>>7|src[17]<<1)&0x1f]
	dst[27] = enc[(src[17]>>4|src[16]<<4)&0x1f]
	dst[26] = enc[src[16]>>1&0x1f]
	dst[25] = enc[(src[16]>>6|src[15]<<2)&0x1f]
	dst[24] = enc[src[15]>>3&0x1f]

	dst[23] = enc[src[14]&0x1f]
	dst[22] = enc[(src[14]>>5|src[13]<<3)&0x1f]
	dst[21] = enc[src[13]>>2&0x1f]
	dst[20] = enc[(src[13]>>7|src[12]<<1)&0x1f]
	dst[19] = enc[(src[12]>>4|src[11]<<4)&0x1f]
	dst[18] = enc[src[11]>>1&0x1f]
	dst[17] = enc[(src[11]>>6|src[10]<<2)&0x1f]
	dst[16] = enc[src[10]>>3&0x1f]

	dst[15] = enc[src[9]&0x1f]
	dst[14] = enc[(src[9]>>5|src[8]<<3)&0x1f]
	dst[13] = enc[src[8]>>2&0x1f]
	dst[12] = enc[(src[8]>>7|src[7]<<1)&0x1f]
	dst[11] = enc[(src[7]>>4|src[6]<<4)&0x1f]
	dst[10] = enc[src[6]>>1&0x1f]
	dst[9] = enc[(src[6]>>6|src[5]<<2)&0x1f]
	dst[8] = enc[src[5]>>3&0x1f]

	dst[7] = enc[src[4]&0x1f]
	dst[6] = enc[(src[4]>>5|src[3]<<3)&0x1f]
	dst[5] = enc[src[3]>>2&0x1f]
	dst[4] = enc[(src[3]>>7|src[2]<<1)&0x1f]
	dst[3] = enc[(src[2]>>4|src[1]<<4)&0x1f]
	dst[2] = enc[src[1]>>1&0x1f]
	dst[1] = enc[(src[1]>>6|src[0]<<2)&0x1f]
	dst[0] = enc[src[0]>>3&0x1f]
}
