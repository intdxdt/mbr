package mbr

import "github.com/intdxdt/math"

// Distance - computes the distance between two MBRs
func (mbr *MBR[T]) Distance(other *MBR[T]) float64 {
	if mbr.Intersects(other) {
		return 0
	}
	var sminx, sminy = mbr.MinX, mbr.MinY
	var smaxx, smaxy = mbr.MaxX, mbr.MaxY
	var ominx, ominy = other.MinX, other.MinY
	var omaxx, omaxy = other.MaxX, other.MaxY

	var dx, dy T
	//find the closest edge by x
	if smaxx < ominx {
		dx = ominx - smaxx
	} else if sminx > omaxx {
		dx = sminx - omaxx
	}

	//find the closest edge by y
	if smaxy < ominy {
		dy = ominy - smaxy
	} else if sminy > omaxy {
		dy = sminy - omaxy
	}
	return math.Hypot(float64(dx), float64(dy))
}
