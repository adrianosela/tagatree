package ctrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adrianosela/tagatree/api/auth"
	"github.com/adrianosela/tagatree/api/objects"
	"github.com/adrianosela/tagatree/api/store"
	"github.com/gorilla/mux"
)

// listTreesHandler lists trees as per a set of filters
func (c *Controller) listTreesHandler(w http.ResponseWriter, r *http.Request) {
	var opts *store.ListOpts

	// get options from request body if its a POST
	if r.Method == http.MethodPost {
		if err := unmarshalRequestBody(r, &opts); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("could not unmarshal request body"))
			return
		}
	}

	list, err := c.trees.ListTrees(opts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an unknown error occurred retrieving trees"))
		return
	}

	respBytes, err := json.Marshal(&list)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize tree list"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

// getTreeHandler gets a tree by id
// note: the id parameter must be included
// in the request URL. (?id=tree_id)
func (c *Controller) getTreeHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	if id = mux.Vars(r)["id"]; id == "" {
		http.Error(w, "no tree id in request URL", http.StatusBadRequest)
		return
	}

	tree, err := c.trees.GetTree(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("tree not found"))
		return
	}

	respBytes, err := json.Marshal(&tree)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize tree"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

// tagTreeHandler posts a tree to the system
func (c *Controller) tagTreeHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var tree objects.Tree
	if err := unmarshalRequestBody(r, &tree); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not unmarshal request body"))
		return
	}

	// include tree tagger username
	tree.TaggedBy = claims.Subject
	if tree.TaggedAt == nil {
		*tree.TaggedAt = time.Now()
	}

	if err := tree.Validate(false); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) // propagate err
		return
	}

	id, err := c.trees.PutTree(&tree)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) // propagate error
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
	return
}
