package mbr

//Expand to include other mbr
func (mbr *MBR) ExpandIncludeMBR(other *MBR) *MBR {
	if other[x1] < mbr[x1] {
		mbr[x1] = other[x1]
	}

	if other[x2] > mbr[x2] {
		mbr[x2] = other[x2]
	}

	if other[y1] < mbr[y1] {
		mbr[y1] = other[y1]
	}

	if other[y2] > mbr[y2] {
		mbr[y2] = other[y2]
	}
	return mbr
}

//ExpandBy expands mbr by change in x and y
func (mbr *MBR) ExpandByDelta(dx, dy float64) *MBR {
	var minx, miny = mbr[x1]-dx, mbr[y1]-dy
	var maxx, maxy = mbr[x2]+dx, mbr[y2]+dy

	minx, maxx = minf64(minx, maxx), maxf64(minx, maxx)
	miny, maxy = minf64(miny, maxy), maxf64(miny, maxy)

	mbr[x1], mbr[y1] = minx, miny
	mbr[x2], mbr[y2] = maxx, maxy

	return mbr
}

//ExpandXY expands mbr to include x and y
func (mbr *MBR) ExpandIncludeXY(x, y float64) *MBR {
	if x < mbr[x1] {
		mbr[x1] = x
	} else if x > mbr[x2] {
		mbr[x2] = x
	}

	if y < mbr[y1] {
		mbr[y1] = y
	} else if y > mbr[y2] {
		mbr[y2] = y
	}

	return mbr
}
