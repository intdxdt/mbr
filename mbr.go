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

// CreateMBR - bounding box
func CreateMBR(minx, miny, maxx, maxy float64) MBR {
	return MBR{
		minF64(minx, maxx),
		minF64(miny, maxy),
		maxF64(minx, maxx),
		maxF64(miny, maxy),
	}
}

// CreateNullMBR - null bounding box
func CreateNullMBR() MBR {
	return MBR{1, 1, -1, -1}
}

// Clone - make a copy of mbr
func (mbr *MBR) Clone() MBR {
	return *mbr
}

// BBox - Bounding Box interface
func (mbr *MBR) BBox() *MBR {
	return mbr
}

// IsNull - Checks if is null
func (mbr *MBR) IsNull() bool {
	return (mbr.MaxX < mbr.MinX) || (mbr.MaxY < mbr.MinY)
}
