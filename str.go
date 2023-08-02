package mbr

import (
	"fmt"
)

// String converts mbr to wkt string
func (mbr *MBR) String() string {
	var lx, ly = mbr.MinX, mbr.MinY
	var ux, uy = mbr.MaxX, mbr.MaxY

	return fmt.Sprintf(
		"POLYGON ((%v %v, %v %v, %v %v, %v %v, %v %v))", lx, ly, lx, uy, ux, uy, ux, ly, lx, ly,
	)
}
