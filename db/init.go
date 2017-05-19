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

	// Creating cluster for cassandra
	cluster := gocql.NewCluster(hosts...)
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = time.Second * 15
	// Auth if username is set
	if libs.Cfg.CassandraUsername != "" {
		libs.Log.Debug("Connecting with username", libs.Cfg.CassandraUsername)
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: libs.Cfg.CassandraUsername,
			Password: libs.Cfg.CassandraPassword,
		}
	}
	if libs.Cfg.CassandraCaPath != "" {
		libs.Log.Debug("Using Ca")
		cluster.SslOpts = &gocql.SslOptions{
			CaPath:                 libs.Cfg.CassandraCaPath,
			EnableHostVerification: libs.Cfg.CassandraSslHostValidation,
		}
	}

	// Initializing session
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	// Storring cassandra session
	Sess = session

	return nil
}
