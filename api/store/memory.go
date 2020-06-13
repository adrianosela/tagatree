package store

import (
	"fmt"
	"sync"

	"github.com/adrianosela/tagatree/api/objects"
	"go.mongodb.org/mongo-driver/bson/primitive"
	geo "gopkg.in/Billups/golang-geo.v2"
)

// Memory implements the Store
// interface in-memory (mock)
type Memory struct {
	sync.RWMutex

	trees map[string]*objects.Tree
}

// NewMemory is the constructor for a Memory storage
func NewMemory() *Memory {
	return &Memory{
		trees: make(map[string]*objects.Tree),
	}
}

// PutTree writes a new tree to the store
func (m *Memory) PutTree(tree *objects.Tree) (string, error) {
	m.Lock()
	defer m.Unlock()

	// using mongo object ID format for consistency
	tree.ID = primitive.NewObjectID()
	m.trees[tree.ID.String()] = tree
	return tree.ID.String(), nil
}

// GetTree reads a tree from the store
func (m *Memory) GetTree(id string) (*objects.Tree, error) {
	m.RLock()
	defer m.RUnlock()

	if tree, ok := m.trees[id]; ok {
		return tree, nil
	}
	return nil, fmt.Errorf("tree %s not in store", id)
}

// ListTrees gets a list of trees from the db
func (m *Memory) ListTrees(opts *ListOpts) ([]*objects.Tree, error) {
	trees := []*objects.Tree{}

	// if no options specified, return all
	if opts == nil {
		for _, tree := range m.trees {
			trees = append(trees, tree)
		}
		return trees, nil
	}

	if err := opts.validate(); err != nil {
		return nil, err
	}

	var location *geo.Point
	if opts.RadiusMeters > 0 {
		location = geo.NewPoint(opts.Location.Coordinates[1], opts.Location.Coordinates[0])
	}

	for _, tree := range m.trees {
		if location != nil {
			p2 := geo.NewPoint(tree.Location.Coordinates[1], tree.Location.Coordinates[0])
			if location.GreatCircleDistance(p2) > float64(opts.RadiusMeters)/1000 {
				continue
			}
		}
		if opts.Species != nil && len(opts.Species) > 0 {
			if !containsString(opts.Species, tree.Species) {
				continue
			}
		}
		if opts.TaggedBy != nil && len(opts.TaggedBy) > 0 {
			if !containsString(opts.TaggedBy, tree.Species) {
				continue
			}
		}
		trees = append(trees, tree)
	}

	return trees, nil
}

// containsString returns true if the container contains the item
func containsString(container []string, item string) (contained bool) {
	for _, i := range container {
		if i == item {
			return true
		}
	}
	return false
}

// UpdateTree updates a tree in the store
func (m *Memory) UpdateTree(tree *objects.Tree) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.trees[tree.ID.String()]; !ok {
		return fmt.Errorf("tree %s not in store", tree.ID)
	}
	m.trees[tree.ID.String()] = tree
	return nil
}

// DeleteTree deletes a tree from the store
func (m *Memory) DeleteTree(id string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.trees, id)
	return nil
}
