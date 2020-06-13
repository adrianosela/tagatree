package objects

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	errTreeHasNoID       = errors.New("tree has no id")
	errTreeHasNoTaggedBy = errors.New("tree has no tagged_by")
	errTreeHasNoTaggedAt = errors.New("tree has no tagged_at")
	errTreeHasNoSpecies  = errors.New("tree has no species")
)

// Tree represents the attributes of a tree
type Tree struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	TaggedBy string             `json:"tagged_by,omitempty" bson:"tagged_by"`
	TaggedAt *time.Time         `json:"tagged_at,omitempty" bson:"tagged_at"`
	Species  string             `json:"species" bson:"species"`
	Location *Location          `json:"location" bson:"location"`
}

// Validate validates a tree has all
// mandatory fields populated
func (t *Tree) Validate(checkID bool) error {
	if checkID && t.ID.String() == "" {
		return errTreeHasNoID
	}
	if t.TaggedBy == "" {
		return errTreeHasNoTaggedBy
	}
	if t.TaggedAt == nil {
		return errTreeHasNoTaggedAt
	}
	if t.Species == "" {
		return errTreeHasNoSpecies
	}
	return t.Location.Validate()
}
