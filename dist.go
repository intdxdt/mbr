package mbr

import "github.com/intdxdt/math"

//Distance computes the distance between two mbrs
func (mbr *MBR) Distance(other *MBR) float64 {
	if mbr.Intersects(other) {
		return 0
	}
	var sminx, sminy =  mbr.minx,    mbr.miny
	var smaxx, smaxy =  mbr.maxx,    mbr.maxy
	var ominx, ominy =  other.minx,  other.miny
	var omaxx, omaxy =  other.maxx,  other.maxy

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
