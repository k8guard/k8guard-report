package db

import (
	"time"

	"github.com/gocql/gocql"
	libs "github.com/k8guard/k8guardlibs"
)

// Here we'll store connection
var Sess *gocql.Session

// This is wrapper for gocql
func Connect(hosts []string) error {

	libs.Log.Info("Connecting to db")

	// Creating Cluster for cassandra
	cluster := gocql.NewCluster(hosts...)
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Second * 5

	// Initializing session
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	// Storring cassandra session
	Sess = session

	return nil
}
