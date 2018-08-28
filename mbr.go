package mbr

import (
	"github.com/intdxdt/math"
)

var nan = math.NaN()

type MBR struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

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

//Checks if is null
func (mbr *MBR) IsNull() bool {
	return (mbr.MaxX < mbr.MinX) || (mbr.MaxY < mbr.MinY)
}
