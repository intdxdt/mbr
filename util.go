package mbr

//Bounding box as Array
func (mbr *MBR) AsArray() []float64 {
	return []float64{mbr.minx, mbr.miny, mbr.maxx, mbr.maxy}
}

//As polygon array
func (mbr *MBR) AsPolyArray() [][]float64 {
	var lx, ly = mbr.minx, mbr.miny
	var ux, uy = mbr.maxx, mbr.maxy
	return [][]float64{
		{lx, ly},
		{lx, uy},
		{ux, uy},
		{ux, ly},
		{lx, ly},
	}
}

//Width of bounding box
func (mbr *MBR) Width() float64 {
	return mbr.maxx - mbr.minx
}

//Height of bounding box
func (mbr *MBR) Height() float64 {
	return mbr.maxy - mbr.miny
}

//Area  of polygon
func (mbr *MBR) Area() float64 {
	return mbr.Width() * mbr.Height()
}

//Is bounding dimensionless like a point: width & height is 0
func (mbr *MBR) IsPoint() bool {
	return mbr.Height() == 0.0 && mbr.Width() == 0.0
}

//Translate mbr  by change in x and y
func (mbr *MBR) Translate(dx, dy float64) MBR {
	return CreateMBR(
		mbr.minx+dx, mbr.miny+dy,
		mbr.maxx+dx, mbr.maxy+dy,
	)
}

//Center of bounding box
func (mbr *MBR) Center() []float64 {
	return []float64{
		(mbr.minx + mbr.maxx) / 2.0,
		(mbr.miny + mbr.maxy) / 2.0,
	}
}

//max
func maxf64(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

//min
func minf64(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}

