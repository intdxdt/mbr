package mbr

//Bounding box as Array
func (mbr MBR) AsArray() []float64 {
	return []float64{mbr[x1], mbr[y1], mbr[x2], mbr[y2]}
}

//As polygon array
func (mbr MBR) AsPolyArray() [][]float64 {
	var lx, ly = mbr[x1], mbr[y1]
	var ux, uy = mbr[x2], mbr[y2]
	return [][]float64{{lx, ly}, {lx, uy}, {ux, uy}, {ux, ly}, {lx, ly}}
}

//Width of bounding box
func (mbr MBR) Width() float64 {
	return mbr[x2] - mbr[x1]
}

//Height of bounding box
func (mbr MBR) Height() float64 {
	return mbr[y2] - mbr[y1]
}

//Area  of polygon
func (mbr MBR) Area() float64 {
	return mbr.Height() * mbr.Width()
}

//Is bounding dimensionless like a point: width & height is 0
func (mbr MBR) IsPoint() bool {
	return mbr.Height() == 0.0 && mbr.Width() == 0.0
}

//Translate mbr  by change in x and y
func (mbr MBR) Translate(dx, dy float64) MBR {
	return CreateMBR(
		mbr[x1]+dx, mbr[y1]+dy,
		mbr[x2]+dx, mbr[y2]+dy,
	)
}

//Center of bounding box
func (mbr MBR) Center() []float64 {
	return []float64{
		(mbr[x1] + mbr[x2]) / 2.0,
		(mbr[y1] + mbr[y2]) / 2.0,
	}
}
