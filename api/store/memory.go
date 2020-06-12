package store

import (
	"fmt"
	"sync"

	"github.com/adrianosela/tagatree/api/objects"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	for _, tree := range m.trees {
		if opts != nil {
			// TODO: filter using opts
		}
		trees = append(trees, tree)
	}

	return trees, nil
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
