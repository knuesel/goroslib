// Autogenerated with msg-import, do not edit.
package rosgraph_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Log struct {
	msgs.Package     `ros:"rosgraph_msgs"`
	msgs.Definitions `ros:"byte DEBUG=1,byte INFO=2,byte WARN=4,byte ERROR=8,byte FATAL=16"`
	Header           std_msgs.Header
	Level            int8 `ros:"byte"`
	Name             string
	Msg              string
	File             string
	Function         string
	Line             uint32
	Topics           []string
}
