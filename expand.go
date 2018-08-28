package mbr

//Expand to include other mbr
func (mbr *MBR) ExpandIncludeMBR(other *MBR) *MBR {
	if other.MinX < mbr.MinX {
		mbr.MinX = other.MinX
	}

	if other.MaxX > mbr.MaxX {
		mbr.MaxX = other.MaxX
	}

	if other.MinY < mbr.MinY {
		mbr.MinY = other.MinY
	}

	if other.MaxY > mbr.MaxY {
		mbr.MaxY = other.MaxY
	}
	return mbr
}

//ExpandBy expands mbr by change in x and y
func (mbr *MBR) ExpandByDelta(dx, dy float64) *MBR {
	var minx, miny = mbr.MinX-dx, mbr.MinY-dy
	var maxx, maxy = mbr.MaxX+dx, mbr.MaxY+dy

	minx, maxx = minf64(minx, maxx), maxf64(minx, maxx)
	miny, maxy = minf64(miny, maxy), maxf64(miny, maxy)

	mbr.MinX, mbr.MinY = minx, miny
	mbr.MaxX, mbr.MaxY = maxx, maxy

	return mbr
}

//ExpandXY expands mbr to include x and y
func (mbr *MBR) ExpandIncludeXY(x, y float64) *MBR {
	if x < mbr.MinX {
		mbr.MinX = x
	} else if x > mbr.MaxX {
		mbr.MaxX = x
	}

	if y < mbr.MinY {
		mbr.MinY = y
	} else if y > mbr.MaxY {
		mbr.MaxY = y
	}

	return mbr
}
