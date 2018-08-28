package mbr

import "github.com/intdxdt/math"


//Checks equallity of two bounding box
func (mbr *MBR) Equals(other *MBR) bool {
	return (
		math.FloatEqual(mbr.minx, other.minx) &&
			math.FloatEqual(mbr.miny, other.miny) &&
			math.FloatEqual(mbr.maxx, other.maxx) &&
			math.FloatEqual(mbr.maxy, other.maxy))
}

//Insersection of two bounding box
func (mbr *MBR) Intersection(other *MBR) (MBR, bool) {
	var minx, miny = nan, nan
	var maxx, maxy = nan, nan
	var intersects = mbr.Intersects(other)

	if intersects {
		if mbr.minx > other.minx {
			minx = mbr.minx
		} else {
			minx = other.minx
		}

		if mbr.miny > other.miny {
			miny = mbr.miny
		} else {
			miny = other.miny
		}

		if mbr.maxx < other.maxx {
			maxx = mbr.maxx
		} else {
			maxx = other.maxx
		}

		if mbr.maxy < other.maxy {
			maxy = mbr.maxy
		} else {
			maxy = other.maxy
		}
	}

	return MBR{minx, miny, maxx, maxy}, intersects
}

//Checks if two bounding boxes intesect
func (mbr *MBR) Intersects(other *MBR) bool {
	//not disjoint
	return !(
		other.minx > mbr.maxx ||
		other.maxx < mbr.minx ||
		other.miny > mbr.maxy ||
		other.maxy < mbr.miny)
}

//Checks if bounding box intersects point
func (mbr *MBR) IntersectsPoint(pt []float64) bool {
	if len(pt) < 2 {
		return false
	}
	return mbr.ContainsXY(pt[0], pt[1])
}

//Intersects bounding box defined by two points a & b
func (mbr *MBR) IntersectsBounds(ax, ay, bx, by float64) bool {
	if mbr.minx > maxf64(ax, bx) || mbr.maxx < minf64(ax, bx) {
		return false
	}
	// not disjoint
	return !(mbr.miny > maxf64(ay, by) || mbr.maxy < minf64(ay, by))
}

//Contains other bounding box
func (mbr *MBR) Contains(other *MBR) bool {
	return (
		(other.minx >= mbr.minx) &&
		(other.maxx <= mbr.maxx) &&
		(other.miny >= mbr.miny) &&
		(other.maxy <= mbr.maxy))
}

func (mbr *MBR) ContainsXY(x, y float64) bool {
	return (x >= mbr.minx) &&
		(x <= mbr.maxx) &&
		(y >= mbr.miny) &&
		(y <= mbr.maxy)
}

//CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (mbr *MBR) CompletelyContainsXY(x, y float64) bool {
	return ((x > mbr.minx) &&
			(x < mbr.maxx) &&
			(y > mbr.miny) &&
			(y < mbr.maxy))
}

//CompletelyContainsMBR is true if mbr completely contains other
func (mbr *MBR) CompletelyContainsMBR(other *MBR) bool {
	return (
		(other.minx > mbr.minx) &&
		(other.maxx < mbr.maxx) &&
		(other.miny > mbr.miny) &&
		(other.maxy < mbr.maxy))
}

//Disjoint of mbr do not intersect
func (mbr *MBR) Disjoint(m *MBR) bool {
	return !mbr.Intersects(m)
}
