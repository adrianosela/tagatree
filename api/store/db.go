package store

import (
	"errors"
	"fmt"

	"github.com/adrianosela/tagatree/api/objects"
)

const (
	idFieldKey       = "_id"
	locationFieldKey = "location"
	taggedByFieldKey = "tagged_by"
	speciesFieldKey  = "species"

	usersCollectionName = "users"
	treesCollectionName = "trees"
)

// DB represents the storage needs
// of a harvest games' manager
type DB interface {
	// TODO: auth related functions

	PutTree(*objects.Tree) (string, error)
	GetTree(string) (*objects.Tree, error)
	ListTrees(*ListOpts) ([]*objects.Tree, error)
	UpdateTree(*objects.Tree) error
	DeleteTree(string) error
}

// ListOpts represents applicable filters
// when querying the db for a list of trees
type ListOpts struct {
	RadiusMeters int               `json:"radius,omitempty"`
	Location     *objects.Location `json:"location,omitempty"`
	TaggedBy     []string          `json:"tagged_by,omitempty"`
	Species      []string          `json:"species,omitempty"`
}

// validate validates the filters specified in ListOpts
func (opts *ListOpts) validate() error {
	if opts.RadiusMeters > 0 {
		if opts.Location == nil {
			return errors.New("radius specified, but location nil")
		}
		if err := opts.Location.Validate(); err != nil {
			return fmt.Errorf("invalid location provided: %s", err)
		}
	}
	return nil
}
