package cancelmiddleware

import (
	"fmt"
	"net/http"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&cancelMiddlewareExtension{})
}

type cancelMiddlewareExtension struct {
}

func (e *cancelMiddlewareExtension) Name() string {
	return "github.com/peterj/fn-extensions/cancelmiddleware"
}

func (e *cancelMiddlewareExtension) Setup(s fnext.ExtServer) error {
	s.AddRootMiddleware(&CancelMiddleware{})
	return nil
}

type CancelMiddleware struct {
}

func (h *CancelMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CancelMiddleware invoked")

		cancelHeader := r.Header.Get("fn-cancel-call")
		if cancelHeader == "1" {
			fmt.Println("Call to function cancelled")
			return
		}
		next.ServeHTTP(w, r)
	})
}
