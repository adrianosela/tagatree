package store

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/adrianosela/tagatree/api/objects"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	// index on the 2dsphere type field "location"
	treeIndexModel = mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{locationFieldKey: bsonx.String("2dsphere")},
	}
)

// Mongo implements the DB
// interface in mongodb
type Mongo struct {
	trees *mongo.Collection
}

// NewMongo is the constructor for a Mongo type Store
//
// the format of the mongo connection string is:
// mongodb://<user>:<pass>@<url>:<port>/<dbname>
func NewMongo(connStr, db string) (*Mongo, error) {
	// mongo setup timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("connected to mongodb")

	trees := client.Database(db).Collection(treesCollectionName)
	ciopts := options.CreateIndexes().SetMaxTime(time.Second * 10)

	if _, err := trees.Indexes().CreateOne(ctx, treeIndexModel, ciopts); err != nil {
		return nil, err
	}

	return &Mongo{
		trees: trees,
	}, nil
}

// PutTree writes a new tree to the db
func (m *Mongo) PutTree(tree *objects.Tree) (string, error) {
	res, err := m.trees.InsertOne(context.TODO(), tree)
	if err != nil {
		return "", err
	}
	id := (res.InsertedID).(primitive.ObjectID).String()
	return id, err
}

// GetTree reads a tree from the db
func (m *Mongo) GetTree(id string) (*objects.Tree, error) {
	idobj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	var tree objects.Tree
	err = m.trees.FindOne(context.TODO(), queryEqual(idFieldKey, idobj)).Decode(&tree)

	return &tree, err
}

// ListTrees gets a list of trees from the db
// note: opts must be validated before passing it this function
func (m *Mongo) ListTrees(opts *ListOpts) ([]*objects.Tree, error) {

	query := bson.D{}

	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		matches := []bson.D{}
		if opts.RadiusMeters > 0 {
			matches = append(matches, queryNear(opts.Location, opts.RadiusMeters))
		}
		if opts.Species != nil && len(opts.Species) > 0 {
			matches = append(matches, queryIn(speciesFieldKey, opts.Species))
		}
		if opts.TaggedBy != nil && len(opts.TaggedBy) > 0 {
			matches = append(matches, queryIn(taggedByFieldKey, opts.TaggedBy))
		}
		if len(matches) == 1 {
			query = matches[0]
		}
		if len(matches) > 1 {
			query = bson.D{{Key: "$and", Value: matches}}
		}
	}

	cur, err := m.trees.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	trees := []*objects.Tree{}
	for cur.Next(context.TODO()) {
		var tree objects.Tree
		err := cur.Decode(&tree)
		if err == nil {
			trees = append(trees, &tree)
		}
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return trees, nil
}

// UpdateTree updates a tree in the db
func (m *Mongo) UpdateTree(tree *objects.Tree) error {
	update := bson.M{
		"$set": bson.M{
			// TODO: set any fields that may be updated
		},
	}
	_, err := m.trees.UpdateOne(context.TODO(), queryEqual(idFieldKey, tree.ID), update)
	return err
}

// DeleteTree deletes a tree from the db
func (m *Mongo) DeleteTree(id string) error {
	idobj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	_, err = m.trees.DeleteOne(context.TODO(), queryEqual(idFieldKey, idobj))
	return err
}

func queryEqual(field string, value interface{}) bson.D {
	return bson.D{
		{
			Key:   idFieldKey,
			Value: value,
		},
	}
}

func queryIn(field string, set []string) bson.D {
	return bson.D{
		{
			Key:   field,
			Value: bson.M{"$in": set},
		},
	}
}

func queryNear(loc *objects.Location, radius int) bson.D {
	return bson.D{
		{
			Key:   locationFieldKey,
			Value: bson.M{"$near": bson.M{"$geometry": loc, "$maxDistance": radius}},
		},
	}
}
