package mbr

import (
	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Integer | constraints.Float
}

type MBR[T Num] struct {
	MinX T
	MinY T
	MaxX T
	MaxY T
}

// NewMBR - bounding box
func NewMBR[T Num](minx, miny, maxx, maxy T) *MBR[T] {
	return &MBR[T]{
		min(minx, maxx),
		min(miny, maxy),
		max(minx, maxx),
		max(miny, maxy),
	}
}

// NewMBRFromArray - bounding box from 2D array
func NewMBRFromArray[T Num](a [2]T, b [2]T) *MBR[T] {
	var minx, miny = a[0], a[1]
	var maxx, maxy = b[0], b[1]
	return NewMBR(minx, miny, maxx, maxy)
}

// NewMBRFromSlice - bounding box from 2D array
func NewMBRFromSlice[T Num](a []T, b []T) *MBR[T] {
	var minx, miny = a[0], a[1]
	var maxx, maxy = b[0], b[1]
	return NewMBR(minx, miny, maxx, maxy)
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
	return MBR[T]{1, 1, 0, 0}
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
