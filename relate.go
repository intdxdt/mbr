package mbr

import "github.com/intdxdt/math"


//Checks equallity of two bounding box
func (mbr MBR) Equals(other MBR) bool {
	return (
		math.FloatEqual(mbr[x1], other[x1]) &&
			math.FloatEqual(mbr[y1], other[y1]) &&
			math.FloatEqual(mbr[x2], other[x2]) &&
			math.FloatEqual(mbr[y2], other[y2]))
}

//Insersection of two bounding box
func (mbr MBR) Intersection(other MBR) (MBR, bool) {
	var minx, miny = nan, nan
	var maxx, maxy = nan, nan
	var inters = mbr.Intersects(other)

	if inters {
		if mbr[x1] > other[x1] {
			minx = mbr[x1]
		} else {
			minx = other[x1]
		}

		if mbr[y1] > other[y1] {
			miny = mbr[y1]
		} else {
			miny = other[y1]
		}

		if mbr[x2] < other[x2] {
			maxx = mbr[x2]
		} else {
			maxx = other[x2]
		}

		if mbr[y2] < other[y2] {
			maxy = mbr[y2]
		} else {
			maxy = other[y2]
		}

	}

	return CreateMBR(minx, miny, maxx, maxy), inters
}

//Checks if two bounding boxes intesect
func (mbr MBR) Intersects(other MBR) bool {
	//not disjoint
	return !(other[x1] > mbr[x2] ||
		other[x2] < mbr[x1] ||
		other[y1] > mbr[y2] ||
		other[y2] < mbr[y1])
}

//Checks if bounding box intersects point
func (mbr MBR) IntersectsPoint(pt []float64) bool {
	if len(pt) < 2 {
		return false
	}
	return mbr.ContainsXY(pt[x1], pt[y1])
}

//Intersects bounding box defined by two points pt1 & pt2
func (mbr MBR) IntersectsBounds(pt1, pt2 []float64) bool {
	if len(pt1) < 2 || len(pt2) < 2 {
		return false
	}
	var minq = math.MinF64(pt1[x1], pt2[x1])
	var maxq = math.MaxF64(pt1[x1], pt2[x1])

	if mbr[x1] > maxq || mbr[x2] < minq {
		return false
	}

	minq = math.MinF64(pt1[y1], pt2[y1])
	maxq = math.MaxF64(pt1[y1], pt2[y1])

	// not disjoint
	return !(mbr[y1] > maxq || mbr[y2] < minq)
}

//Contains other bounding box
func (mbr MBR) Contains(other MBR) bool {
	return ((other[x1] >= mbr[x1]) &&
		(other[x2] <= mbr[x2]) &&
		(other[y1] >= mbr[y1]) &&
		(other[y2] <= mbr[y2]))
}

func (mbr MBR) ContainsXY(x, y float64) bool {
	return (x >= mbr[x1]) &&
		(x <= mbr[x2]) &&
		(y >= mbr[y1]) &&
		(y <= mbr[y2])
}

//CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (mbr MBR) CompletelyContainsXY(x, y float64) bool {
	return (
		(x > mbr[x1]) &&
			(x < mbr[x2]) &&
			(y > mbr[y1]) &&
			(y < mbr[y2]))
}

//CompletelyContainsMBR is true if mbr completely contains other
func (mbr MBR) CompletelyContainsMBR(other MBR) bool {
	return ((other[x1] > mbr[x1]) &&
		(other[x2] < mbr[x2]) &&
		(other[y1] > mbr[y1]) &&
		(other[y2] < mbr[y2]))
}

//Disjoint of mbr do not intersect
func (mbr MBR) Disjoint(m MBR) bool {
	return !mbr.Intersects(m)
}
