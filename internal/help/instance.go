package help

import (
	"app/internal/bootstrap"
	"github.com/bwmarrin/snowflake"
)

var (
	Dir    = bootstrap.Dir
	Env    = bootstrap.Env
	Config = bootstrap.Config
	Libs   *libraries
	Token  *tokenHelp
)

type libraries struct {
	Snow *snowflake.Node
}

func newLibs() *libraries {
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(Config.Snowflake.Node)
	return &libraries{
		Snow: sn,
	}
}

func init() {
	Libs = newLibs()
	Token = newTokenHelp()
}
