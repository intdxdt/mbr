package mbr

// AsArray - bounding box as Array
func (mbr *MBR) AsArray() []float64 {
	return []float64{mbr.MinX, mbr.MinY, mbr.MaxX, mbr.MaxY}
}

// AsPolyArray - as polygon array
func (mbr *MBR) AsPolyArray() [][]float64 {
	var lx, ly = mbr.MinX, mbr.MinY
	var ux, uy = mbr.MaxX, mbr.MaxY
	return [][]float64{
		{lx, ly},
		{lx, uy},
		{ux, uy},
		{ux, ly},
		{lx, ly},
	}
}

// Width of bounding box
func (mbr *MBR) Width() float64 {
	return mbr.MaxX - mbr.MinX
}

// Height of bounding box
func (mbr *MBR) Height() float64 {
	return mbr.MaxY - mbr.MinY
}

// Area  of polygon
func (mbr *MBR) Area() float64 {
	return mbr.Width() * mbr.Height()
}

// IsPoint - is bounding dimensionless like a point: width & height is 0 ?
func (mbr *MBR) IsPoint() bool {
	return mbr.Height() == 0.0 && mbr.Width() == 0.0
}

// Translate mbr  by change in x and y
func (mbr *MBR) Translate(dx, dy float64) MBR {
	return CreateMBR(
		mbr.MinX+dx, mbr.MinY+dy,
		mbr.MaxX+dx, mbr.MaxY+dy,
	)
}

// Center of bounding box
func (mbr *MBR) Center() []float64 {
	return []float64{
		(mbr.MinX + mbr.MaxX) / 2.0,
		(mbr.MinY + mbr.MaxY) / 2.0,
	}
}

// max
func maxF64(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

// min
func minF64(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}
