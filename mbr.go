package mbr

import (
	"github.com/intdxdt/math"
	"golang.org/x/exp/constraints"
)

var nan = math.NaN()

type Num interface {
	constraints.Signed | constraints.Float
}

type MBR[T Num] struct {
	MinX T
	MinY T
	MaxX T
	MaxY T
}

// CreateMBR - bounding box
func CreateMBR[T Num](minx, miny, maxx, maxy T) MBR[T] {
	return MBR[T]{
		min(minx, maxx),
		min(miny, maxy),
		max(minx, maxx),
		max(miny, maxy),
	}
}

// CreateNullMBR - null bounding box
func CreateNullMBR[T Num]() MBR[T] {
	return MBR[T]{1, 1, -1, -1}
}

// Clone - make a copy of mbr
func (mbr *MBR[T]) Clone() MBR[T] {
	return *mbr
}

// BBox - Bounding Box interface
func (mbr *MBR[T]) BBox() *MBR[T] {
	return mbr
}

// IsNull - Checks if is null
func (mbr *MBR[T]) IsNull() bool {
	return (mbr.MaxX < mbr.MinX) || (mbr.MaxY < mbr.MinY)
}
