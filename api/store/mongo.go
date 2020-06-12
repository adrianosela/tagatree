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
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
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
	err = m.trees.FindOne(context.TODO(), querySingle(idobj)).Decode(&tree)

	return &tree, err
}

// ListTrees gets a list of trees from the db
func (m *Mongo) ListTrees(opts *ListOpts) ([]*objects.Tree, error) {

	matches := bson.M{}

	if opts != nil {
		// TODO: filter using opts
	}

	cur, err := m.trees.Find(context.TODO(), matches)
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
	_, err := m.trees.UpdateOne(context.TODO(), querySingle(tree.ID), update)
	return err
}

// DeleteTree deletes a tree from the db
func (m *Mongo) DeleteTree(id string) error {
	idobj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	_, err = m.trees.DeleteOne(context.TODO(), querySingle(idobj))
	return err
}

func querySingle(id primitive.ObjectID) bson.D {
	return bson.D{{Key: genericPrimaryKey, Value: id}}
}
