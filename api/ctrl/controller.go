package ctrl

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/adrianosela/tagatree/api/auth"
	"github.com/adrianosela/tagatree/api/config"
	"github.com/adrianosela/tagatree/api/store"
)

// Controller manages the server
type Controller struct {
	trees store.DB
	auth  *auth.Authenticator
}

// NewController is the Controller constructor
func NewController(conf config.Conf) (*Controller, error) {
	db, err := store.NewMongo(conf.MongoDBConnStr, conf.MongoDBName)
	if err != nil {
		return nil, err
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &Controller{
		trees: db,
		auth:  auth.NewAuthenticator(db, priv, "tagatree.adrianosela.com", "api"),
	}, nil
}
