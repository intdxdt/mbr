package mbr

import (
	"fmt"
)

//String converts mbr to wkt string
func (mbr *MBR) String() string {
	var lx, ly = mbr[x1], mbr[y1]
	var ux, uy = mbr[x2], mbr[y2]

	return fmt.Sprintf(
		"POLYGON ((%v %v, %v %v, %v %v, %v %v, %v %v))",
		lx, ly, lx, uy, ux, uy, ux, ly, lx, ly,
	)
}
