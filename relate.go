package mbr

import "github.com/intdxdt/math"

// Equals - checks equality of two bounding box
func (mbr *MBR) Equals(other *MBR) bool {
	var minEq = math.FloatEqual(mbr.MinX, other.MinX) && math.FloatEqual(mbr.MinY, other.MinY)
	var maxEq = math.FloatEqual(mbr.MaxX, other.MaxX) && math.FloatEqual(mbr.MaxY, other.MaxY)
	return minEq && maxEq
}

// Intersection  of two bounding box
func (mbr *MBR) Intersection(other *MBR) (MBR, bool) {
	var minx, miny = nan, nan
	var maxx, maxy = nan, nan
	var intersects = mbr.Intersects(other)

	if intersects {
		if mbr.MinX > other.MinX {
			minx = mbr.MinX
		} else {
			minx = other.MinX
		}

		if mbr.MinY > other.MinY {
			miny = mbr.MinY
		} else {
			miny = other.MinY
		}

		if mbr.MaxX < other.MaxX {
			maxx = mbr.MaxX
		} else {
			maxx = other.MaxX
		}

		if mbr.MaxY < other.MaxY {
			maxy = mbr.MaxY
		} else {
			maxy = other.MaxY
		}
	}

	return MBR{minx, miny, maxx, maxy}, intersects
}

// Checks if two bounding boxes intesect
func (mbr *MBR) Intersects(other *MBR) bool {
	//not disjoint
	return !(other.MinX > mbr.MaxX ||
		other.MaxX < mbr.MinX ||
		other.MinY > mbr.MaxY ||
		other.MaxY < mbr.MinY)
}

// Checks if bounding box intersects point
func (mbr *MBR) IntersectsPoint(pt []float64) bool {
	if len(pt) < 2 {
		return false
	}
	return mbr.ContainsXY(pt[0], pt[1])
}

// Intersects bounding box defined by two points a & b
func (mbr *MBR) IntersectsBounds(ax, ay, bx, by float64) bool {
	if mbr.MinX > maxF64(ax, bx) || mbr.MaxX < minF64(ax, bx) {
		return false
	}
	// not disjoint
	return !(mbr.MinY > maxF64(ay, by) || mbr.MaxY < minF64(ay, by))
}

// Contains other bounding box
func (mbr *MBR) Contains(other *MBR) bool {
	return ((other.MinX >= mbr.MinX) &&
		(other.MaxX <= mbr.MaxX) &&
		(other.MinY >= mbr.MinY) &&
		(other.MaxY <= mbr.MaxY))
}

func (mbr *MBR) ContainsXY(x, y float64) bool {
	return (x >= mbr.MinX) &&
		(x <= mbr.MaxX) &&
		(y >= mbr.MinY) &&
		(y <= mbr.MaxY)
}

// CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (mbr *MBR) CompletelyContainsXY(x, y float64) bool {
	return ((x > mbr.MinX) &&
		(x < mbr.MaxX) &&
		(y > mbr.MinY) &&
		(y < mbr.MaxY))
}

// CompletelyContainsMBR is true if mbr completely contains other
func (mbr *MBR) CompletelyContainsMBR(other *MBR) bool {
	return ((other.MinX > mbr.MinX) &&
		(other.MaxX < mbr.MaxX) &&
		(other.MinY > mbr.MinY) &&
		(other.MaxY < mbr.MaxY))
}

// Disjoint of mbr do not intersect
func (mbr *MBR) Disjoint(m *MBR) bool {
	return !mbr.Intersects(m)
}
