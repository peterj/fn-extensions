package callcount

import (
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

var callCountMap map[string]int

func init() {
	callCountMap = make(map[string]int)
	server.RegisterExtension(&callCountExtension{})
}

type callCountExtension struct {
}

func (e *callCountExtension) Name() string {
	return "github.com/peterj/fn-extensions/callcount"
}

func (e *callCountExtension) Setup(s fnext.ExtServer) error {
	s.AddCallListener(&CallCount{})

	return nil
}

type CallCount struct {
}
