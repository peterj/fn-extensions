package calllog

import (
	"fmt"
	"net/http"

	"github.com/fnproject/fn/api/models"
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&callLogAPIExtension{})
}

type callLogAPIExtension struct {
}

func (e *callLogAPIExtension) Name() string {
	return "github.com/peterj/fn-extensions/calllog"
}

func (e *callLogAPIExtension) Setup(s fnext.ExtServer) error {
	apiExt := &CallLogAPIExtension{
		ds: s.Datastore(),
	}
	s.AddAppEndpoint("GET", "/logs", apiExt)
	return nil
}

type CallLogAPIExtension struct {
	ds models.Datastore
}

func (h *CallLogAPIExtension) ServeHTTP(w http.ResponseWriter, r *http.Request, app *models.App) {
	ctx := r.Context()
	filter := &models.CallFilter{
		AppID:   app.ID,
		PerPage: 100,
	}
	calls, err := h.ds.GetCalls(ctx, filter)
	if err != nil {
		server.HandleErrorResponse(ctx, w, err)
	}

	for _, c := range calls {
		fmt.Fprint(w, "------------------------------------------")
		fmt.Fprintf(w, "\nid: %s\nstatus: %s\npath: %s\nstarted at: %s\n", c.ID, c.Status, c.Path, c.StartedAt.String())
	}
	fmt.Fprint(w, "------------------------------------------")
}
