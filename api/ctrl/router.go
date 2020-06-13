package ctrl

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router returns the HTTP routes handler
func (c *Controller) Router() http.Handler {
	r := mux.NewRouter()

	// auth endpoints
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(c.loginHandler)

	// tree operations
	r.Methods(http.MethodGet).Path("/tree/{id}").Handler(c.auth.Wrap(c.getTreeHandler))
	r.Methods(http.MethodGet, http.MethodPost).Path("/trees").Handler(c.auth.Wrap(c.listTreesHandler))
	r.Methods(http.MethodPost).Path("/tree").Handler(c.auth.Wrap(c.tagTreeHandler))

	return cors(r) // enable CORS because... javascript
}
