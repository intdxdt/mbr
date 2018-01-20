package mbr

import (
	"github.com/intdxdt/math"
)

type MBR [4]float64

const (
	x1 = iota
	y1
	x2
	y2
)

//Create new bounding box
func NewMBR(minx, miny, maxx, maxy float64) *MBR {
	minx, maxx = math.MinF64(minx, maxx), math.MaxF64(minx, maxx)
	miny, maxy = math.MinF64(miny, maxy), math.MaxF64(miny, maxy)
	return &MBR{minx, miny, maxx, maxy}
}

//Clone bounding box
func (mbr *MBR) Clone() *MBR {
	return &MBR{mbr[x1], mbr[y1], mbr[x2], mbr[y2]}
}

//Bounding Box interface
func (mbr *MBR) BBox() *MBR {
	return mbr
}

//Minimum X value
func (mbr *MBR) MinX() float64 {
	return mbr[x1]
}

//Minimum Y value
func (mbr *MBR) MinY() float64 {
	return mbr[y1]
}

//Maximum X value
func (mbr *MBR) MaxX() float64 {
	return mbr[x2]
}

//Maximum Y value
func (mbr *MBR) MaxY() float64 {
	return mbr[y2]
}
