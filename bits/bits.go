package bits

type Chess struct{}

func (c *Chess) GetKingBitboardMoves(pos int) uint64 {
	var K uint64 = uint64(1) << pos
	var noA uint64 = 0xfefefefefefefefe
	var noH uint64 = 0x7f7f7f7f7f7f7f7f
	var Ka uint64 = K & noA
	var Kh uint64 = K & noH

	var mask uint64 = (Ka << 7) | (K << 8) | (Kh << 9) |
		(Ka >> 1) | (Kh << 1) |
		(Ka >> 9) | (K >> 8) | (Kh >> 7)

	return mask
}

func (c *Chess) GetHorseBitboardMoves(pos int) uint64 {
	var K uint64 = uint64(1) << pos
	var nA uint64 = 0xFeFeFeFeFeFeFeFe
	var nAB uint64 = 0xFcFcFcFcFcFcFcFc
	var nH uint64 = 0x7f7f7f7f7f7f7f7f
	var nGH uint64 = 0x3f3f3f3f3f3f3f3f

	var mask uint64 = nGH&(K<<6|K>>10) |
		nH&(K<<15|K>>17) |
		nA&(K<<17|K>>15) |
		nAB&(K<<10|K>>6)

	return mask
}

func (c *Chess) Popcnt(mask uint64) int {
	count := 0

	for mask > 0 {
		if (mask & uint64(1)) == 1 {
			count++
		}

		mask >>= 1
	}

	return count
}

func (c *Chess) PopcntOpt(mask uint64) int {
	count := 0

	for mask > 0 {
		count++

		mask &= mask - 1
	}

	return count
}
