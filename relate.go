package mbr

import "github.com/intdxdt/math"

// Equals - checks equality of two bounding box
func (mbr *MBR[T]) Equals(other *MBR[T]) bool {
	var minEq = math.Equals(mbr.MinX, other.MinX) && math.Equals(mbr.MinY, other.MinY)
	var maxEq = math.Equals(mbr.MaxX, other.MaxX) && math.Equals(mbr.MaxY, other.MaxY)
	return minEq && maxEq
}

// Intersection  of two bounding box
func (mbr *MBR[T]) Intersection(other *MBR[T]) (MBR[T], bool) {
	var minx, miny = T(nan), T(nan)
	var maxx, maxy = T(nan), T(nan)
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

	return MBR[T]{minx, miny, maxx, maxy}, intersects
}

// Intersects - Checks if two bounding boxes intesect
func (mbr *MBR[T]) Intersects(other *MBR[T]) bool {
	//not disjoint
	return !(other.MinX > mbr.MaxX ||
		other.MaxX < mbr.MinX ||
		other.MinY > mbr.MaxY ||
		other.MaxY < mbr.MinY)
}

// IntersectsPoint - Checks if bounding box intersects point
func (mbr *MBR[T]) IntersectsPoint(pt []T) bool {
	if len(pt) < 2 {
		return false
	}
	return mbr.ContainsXY(pt[0], pt[1])
}

// IntersectsBounds - Intersects bounding box defined by two points a & b
func (mbr *MBR[T]) IntersectsBounds(ax, ay, bx, by T) bool {
	if mbr.MinX > max(ax, bx) || mbr.MaxX < min(ax, bx) {
		return false
	}
	// not disjoint
	return !(mbr.MinY > max(ay, by) || mbr.MaxY < min(ay, by))
}

// Contains other bounding box
func (mbr *MBR[T]) Contains(other *MBR[T]) bool {
	return (other.MinX >= mbr.MinX) &&
		(other.MaxX <= mbr.MaxX) &&
		(other.MinY >= mbr.MinY) &&
		(other.MaxY <= mbr.MaxY)
}

// ContainsXY - contains x, y
func (mbr *MBR[T]) ContainsXY(x, y T) bool {
	return (x >= mbr.MinX) &&
		(x <= mbr.MaxX) &&
		(y >= mbr.MinY) &&
		(y <= mbr.MaxY)
}

// CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (mbr *MBR[T]) CompletelyContainsXY(x, y T) bool {
	return (x > mbr.MinX) &&
		(x < mbr.MaxX) &&
		(y > mbr.MinY) &&
		(y < mbr.MaxY)
}

// CompletelyContainsMBR is true if mbr completely contains other
func (mbr *MBR[T]) CompletelyContainsMBR(other *MBR[T]) bool {
	return (other.MinX > mbr.MinX) &&
		(other.MaxX < mbr.MaxX) &&
		(other.MinY > mbr.MinY) &&
		(other.MaxY < mbr.MaxY)
}

// Disjoint of mbr do not intersect
func (mbr *MBR[T]) Disjoint(m *MBR[T]) bool {
	return !mbr.Intersects(m)
}
