package mbr

import (
	"fmt"
)

//String converts mbr to wkt string
func (mbr *MBR) String() string {
	var lx, ly = mbr.minx, mbr.miny
	var ux, uy = mbr.maxx, mbr.maxy

	return fmt.Sprintf(
		"POLYGON ((%v %v, %v %v, %v %v, %v %v, %v %v))",
		lx, ly, lx, uy, ux, uy, ux, ly, lx, ly,
	)
}
