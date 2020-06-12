package store

import (
	"github.com/adrianosela/tagatree/api/objects"
)

const (
	genericPrimaryKey = "_id"

	usersCollectionName = "users"
	treesCollectionName = "trees"
)

// DB represents the storage needs
// of a harvest games' manager
type DB interface {
	PutTree(*objects.Tree) (string, error)
	GetTree(string) (*objects.Tree, error)
	ListTrees(*ListOpts) ([]*objects.Tree, error)
	UpdateTree(*objects.Tree) error
	DeleteTree(string) error
}

// ListOpts represents applicable filters
// when querying the db for a list of trees
type ListOpts struct {
	// TODO: provide some filtering options
	// i.e. geolocation proximity, tree type, ...
}
