// Autogenerated with msg-import, do not edit.
package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type TwistWithCovariance struct {
	msgs.Package `ros:"geometry_msgs"`
	Twist        Twist
	Covariance   [36]float64
}
