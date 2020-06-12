package objects

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	errTreeHasNoID       = errors.New("tree has no id")
	errTreeHasNoTaggedBy = errors.New("tree has no tagger")
	errTreeHasNoSpecies  = errors.New("tree has no species")
)

// Tree represents the attributes of a tree
type Tree struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaggedBy string             `json:"tagged_by" bson:"tagged_by"`
	Species  string             `json:"species" bson:"species"`
	Location Location           `json:"location" bson:"location"`
}

// Validate validates a tree has all
// mandatory fields populated
func (t *Tree) Validate(checkID, checkTaggedBy bool) error {
	if checkID && t.ID.String() == "" {
		return errTreeHasNoID
	}
	if checkTaggedBy && t.TaggedBy == "" {
		return errTreeHasNoTaggedBy
	}
	if t.Species == "" {
		return errTreeHasNoSpecies
	}
	return t.Location.Validate()
}
