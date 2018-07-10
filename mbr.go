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
	minx, maxx = math.MinF64(minx, maxx), math.MaxF64(minx, maxx)
	miny, maxy = math.MinF64(miny, maxy), math.MaxF64(miny, maxy)
	return MBR{minx, miny, maxx, maxy}
}


//Bounding Box interface
func (mbr MBR) BBox() MBR {
	return mbr
}

