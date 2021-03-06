// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
)

type InteractiveMarkerControl struct {
	msgs.Package                 `ros:"visualization_msgs"`
	msgs.Definitions             `ros:"uint8 INHERIT=0,uint8 FIXED=1,uint8 VIEW_FACING=2,uint8 NONE=0,uint8 MENU=1,uint8 BUTTON=2,uint8 MOVE_AXIS=3,uint8 MOVE_PLANE=4,uint8 ROTATE_AXIS=5,uint8 MOVE_ROTATE=6,uint8 MOVE_3D=7,uint8 ROTATE_3D=8,uint8 MOVE_ROTATE_3D=9"`
	Name                         string
	Orientation                  geometry_msgs.Quaternion
	OrientationMode              uint8
	InteractionMode              uint8
	AlwaysVisible                bool
	Markers                      []Marker
	IndependentMarkerOrientation bool
	Description                  string
}
