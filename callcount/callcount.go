package callcount

import (
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

var callCountMap map[string]int

func init() {
	server.RegisterExtension(&callCountExtension{})
}

type callCountExtension struct {
}

func (e *callCountExtension) Name() string {
	return "github.com/peterj/fn-extensions/callcount"
}

func (e *callCountExtension) Setup(s fnext.ExtServer) error {
	cc := &CallCount{}

	// Add a Call listener
	s.AddCallListener(cc)
	return nil
}

type CallCount struct {
}
