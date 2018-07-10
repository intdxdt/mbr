package mbr

import "github.com/intdxdt/math"

//Distance computes the distance between two mbrs
func (mbr *MBR) Distance(other MBR) float64 {
	if mbr.Intersects(other) {
		return 0.0
	}
	var sminx, sminy =  mbr[x1], mbr[y1]
	var smaxx, smaxy =  mbr[x2], mbr[y2]
	var ominx, ominy = other[x1], other[y1]
	var omaxx, omaxy = other[x2], other[y2]

	var dx, dy float64
	//find closest edge by x
	if smaxx < ominx {
		dx = ominx - smaxx
	} else if sminx > omaxx {
		dx = sminx - omaxx
	}

	//find closest edge by y
	if smaxy < ominy {
		dy = ominy - smaxy
	} else if sminy > omaxy {
		dy = sminy - omaxy
	}
	return math.Hypot(dx, dy)
}
