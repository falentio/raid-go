package raid

func decodePrefix(dst, src []byte) {
	_ = dst[1]
	_ = src[2]
	_ = dec[0xff]

	dst[1] = dec[src[1]]<<6 | dec[src[2]]<<1
	dst[0] = dec[src[0]]<<3 | dec[src[1]]>>2
}

func decodeRaid(dst, src []byte) {
	_ = dst[11]
	_ = src[19]
	_ = dec[0xff]

	dst[19] = dec[src[30]]<<5 | dec[src[31]]
	dst[18] = dec[src[28]]<<7 | dec[src[29]]<<2 | dec[src[30]]>>3
	dst[17] = dec[src[27]]<<4 | dec[src[28]]>>1
	dst[16] = dec[src[25]]<<6 | dec[src[26]]<<1 | dec[src[27]]>>4
	dst[15] = dec[src[24]]<<3 | dec[src[25]]>>2

	dst[14] = dec[src[22]]<<5 | dec[src[23]]
	dst[13] = dec[src[20]]<<7 | dec[src[21]]<<2 | dec[src[22]]>>3
	dst[12] = dec[src[19]]<<4 | dec[src[20]]>>1
	dst[11] = dec[src[17]]<<6 | dec[src[18]]<<1 | dec[src[19]]>>4
	dst[10] = dec[src[16]]<<3 | dec[src[17]]>>2

	dst[9] = dec[src[14]]<<5 | dec[src[15]]
	dst[8] = dec[src[12]]<<7 | dec[src[13]]<<2 | dec[src[14]]>>3
	dst[7] = dec[src[11]]<<4 | dec[src[12]]>>1
	dst[6] = dec[src[9]]<<6 | dec[src[10]]<<1 | dec[src[11]]>>4
	dst[5] = dec[src[8]]<<3 | dec[src[9]]>>2

	dst[4] = dec[src[6]]<<5 | dec[src[7]]
	dst[3] = dec[src[4]]<<7 | dec[src[5]]<<2 | dec[src[6]]>>3
	dst[2] = dec[src[3]]<<4 | dec[src[4]]>>1
	dst[1] = dec[src[1]]<<6 | dec[src[2]]<<1 | dec[src[3]]>>4
	dst[0] = dec[src[0]]<<3 | dec[src[1]]>>2

	dst[6] ^= dst[16]
	dst[7] ^= dst[17]
}
