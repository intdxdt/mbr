package mbr

import (
	"github.com/intdxdt/math"
)

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

	minx, miny := mbr[x1]-dx, mbr[y1]-dy
	maxx, maxy := mbr[x2]+dx, mbr[y2]+dy

	minx, maxx = math.MinF64(minx, maxx), math.MaxF64(minx, maxx)
	miny, maxy = math.MinF64(miny, maxy), math.MaxF64(miny, maxy)

	mbr[x1], mbr[y1] = minx, miny
	mbr[x2], mbr[y2] = maxx, maxy

	return mbr
}

//ExpandXY expands mbr to include x and y
func (mbr *MBR) ExpandIncludeXY(xCoord, yCoord float64) *MBR {

	if xCoord < mbr[x1] {
		mbr[x1] = xCoord
	} else if xCoord > mbr[x2] {
		mbr[x2] = xCoord
	}

	if yCoord < mbr[y1] {
		mbr[y1] = yCoord
	} else if yCoord > mbr[y2] {
		mbr[y2] = yCoord
	}

	return mbr
}
