package dsd

var (
	deinterleaveMatrixW = []uint8{
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01, 0x00, 0x02,
		0x00, 0x02, 0x00, 0x02, 0x00, 0x02,
		0x00, 0x02, 0x00, 0x02, 0x00, 0x02,
	}
	deinterleaveMatrixX = []uint8{
		0x17, 0x0a, 0x16, 0x09, 0x15, 0x08,
		0x14, 0x07, 0x13, 0x06, 0x12, 0x05,
		0x11, 0x04, 0x10, 0x03, 0x0f, 0x02,
		0x0e, 0x01, 0x0d, 0x00, 0x0c, 0x0a,
		0x0b, 0x09, 0x0a, 0x08, 0x09, 0x07,
		0x08, 0x06, 0x07, 0x05, 0x06, 0x04,
	}
	deinterleaveMatrixY = []uint8{
		0x00, 0x02, 0x00, 0x02, 0x00, 0x02,
		0x00, 0x02, 0x00, 0x03, 0x00, 0x03,
		0x01, 0x03, 0x01, 0x03, 0x01, 0x03,
		0x01, 0x03, 0x01, 0x03, 0x01, 0x03,
		0x01, 0x03, 0x01, 0x03, 0x01, 0x03,
		0x01, 0x03, 0x01, 0x03, 0x01, 0x03,
	}
	deinterleaveMatrixZ = []uint8{
		0x05, 0x03, 0x04, 0x02, 0x03, 0x01,
		0x02, 0x00, 0x01, 0x0d, 0x00, 0x0c,
		0x16, 0x0b, 0x15, 0x0a, 0x14, 0x09,
		0x13, 0x08, 0x12, 0x07, 0x11, 0x06,
		0x10, 0x05, 0x0f, 0x04, 0x0e, 0x03,
		0x0d, 0x02, 0x0c, 0x01, 0x0b, 0x00,
	}
)

type Deinterleaver struct {
	W, X, Y, Z []uint8
}

func NewDeinterleaver() *Deinterleaver {
	d := &Deinterleaver{
		W: make([]uint8, 36),
		X: make([]uint8, 36),
		Y: make([]uint8, 36),
		Z: make([]uint8, 36),
	}
	copy(d.W, deinterleaveMatrixW)
	copy(d.X, deinterleaveMatrixX)
	copy(d.Y, deinterleaveMatrixY)
	copy(d.Z, deinterleaveMatrixZ)
	return d
}
