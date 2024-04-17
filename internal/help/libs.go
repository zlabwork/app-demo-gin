package help

import (
	"github.com/bwmarrin/snowflake"
	"os"
	"strconv"
)

var (
	Libs *libraries
)

type libraries struct {
	Snow *snowflake.Node
}

func newLibs() *libraries {
	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libraries{
		Snow: sn,
	}
}

func init() {
	Libs = newLibs()
}
