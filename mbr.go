package mbr

import (
	"github.com/intdxdt/math"
)

type MBR [4]float64
var nan = math.NaN()
const (
	x1 = iota
	y1
	x2
	y2
)

//Create new bounding box
func CreateMBR(minx, miny, maxx, maxy float64) MBR {
	return MBR{
		minf64(minx, maxx),
		minf64(miny, maxy),
		maxf64(minx, maxx),
		maxf64(miny, maxy),
	}
}

//Make a copy of mbr
func (mbr *MBR) Clone() MBR {
	return *mbr
}

//Bounding Box interface
func (mbr *MBR) BBox() *MBR {
	return mbr
}

//Bounding Box interface
func (mbr *MBR) Bounds() MBR {
	return *mbr
}

//Checks if is null
func (mbr *MBR) IsNull() bool {
	return (mbr[x2] < mbr[x1]) || (mbr[y2] < mbr[y1])
}

