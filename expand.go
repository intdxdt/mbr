package mbr

// ExpandIncludeMBR - expand to include other mbr
func (mbr *MBR[T]) ExpandIncludeMBR(other *MBR[T]) *MBR[T] {
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

// ExpandByDelta - expands mbr by change in x and y
func (mbr *MBR[T]) ExpandByDelta(dx, dy T) *MBR[T] {
	var minx, miny = mbr.MinX - dx, mbr.MinY - dy
	var maxx, maxy = mbr.MaxX + dx, mbr.MaxY + dy

	minx, maxx = min(minx, maxx), max(minx, maxx)
	miny, maxy = min(miny, maxy), max(miny, maxy)

	mbr.MinX, mbr.MinY = minx, miny
	mbr.MaxX, mbr.MaxY = maxx, maxy

	return mbr
}

// ExpandIncludeXY - expands mbr to include x and y
func (mbr *MBR[T]) ExpandIncludeXY(x, y T) *MBR[T] {
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
