package mbr

// AsArray - bounding box as Array
func (mbr *MBR[T]) AsArray() []T {
	return []T{mbr.MinX, mbr.MinY, mbr.MaxX, mbr.MaxY}
}

// AsPolyArray - as polygon array
func (mbr *MBR[T]) AsPolyArray() [][]T {
	var lx, ly = mbr.MinX, mbr.MinY
	var ux, uy = mbr.MaxX, mbr.MaxY
	return [][]T{
		{lx, ly},
		{lx, uy},
		{ux, uy},
		{ux, ly},
		{lx, ly},
	}
}

// Width of bounding box
func (mbr *MBR[T]) Width() T {
	return mbr.MaxX - mbr.MinX
}

// Height of bounding box
func (mbr *MBR[T]) Height() T {
	return mbr.MaxY - mbr.MinY
}

// Area  of polygon
func (mbr *MBR[T]) Area() T {
	return mbr.Width() * mbr.Height()
}

// IsPoint - is bounding dimensionless like a point: width & height is 0 ?
func (mbr *MBR[T]) IsPoint() bool {
	return mbr.Height() == 0.0 && mbr.Width() == 0.0
}

// Translate mbr  by change in x and y
func (mbr *MBR[T]) Translate(dx, dy T) MBR[T] {
	return CreateMBR[T](
		mbr.MinX+dx, mbr.MinY+dy,
		mbr.MaxX+dx, mbr.MaxY+dy,
	)
}

// Center of bounding box
func (mbr *MBR[T]) Center() []T {
	return []T{
		(mbr.MinX + mbr.MaxX) / T(2.0),
		(mbr.MinY + mbr.MaxY) / T(2.0),
	}
}
