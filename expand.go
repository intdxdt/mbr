package mbr

//Expand to include other mbr
func (mbr *MBR) ExpandIncludeMBR(other *MBR) *MBR {
	if other.minx < mbr.minx {
		mbr.minx = other.minx
	}

	if other.maxx > mbr.maxx {
		mbr.maxx = other.maxx
	}

	if other.miny < mbr.miny {
		mbr.miny = other.miny
	}

	if other.maxy > mbr.maxy {
		mbr.maxy = other.maxy
	}
	return mbr
}

//ExpandBy expands mbr by change in x and y
func (mbr *MBR) ExpandByDelta(dx, dy float64) *MBR {
	var minx, miny = mbr.minx-dx, mbr.miny-dy
	var maxx, maxy = mbr.maxx+dx, mbr.maxy+dy

	minx, maxx = minf64(minx, maxx), maxf64(minx, maxx)
	miny, maxy = minf64(miny, maxy), maxf64(miny, maxy)

	mbr.minx, mbr.miny = minx, miny
	mbr.maxx, mbr.maxy = maxx, maxy

	return mbr
}

//ExpandXY expands mbr to include x and y
func (mbr *MBR) ExpandIncludeXY(x, y float64) *MBR {
	if x < mbr.minx {
		mbr.minx = x
	} else if x > mbr.maxx {
		mbr.maxx = x
	}

	if y < mbr.miny {
		mbr.miny = y
	} else if y > mbr.maxy {
		mbr.maxy = y
	}

	return mbr
}
