package database

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

// Session This will establish one or more connections with the cluster of
// servers defined by the mgo.Dial parameter(comma separated)
var Session *mgo.Session

func init() {
	var err error
	dbHosts := os.Getenv("MONGODB_DB_HOSTS")
	if dbHosts != "" {
		Session, err = mgo.Dial(dbHosts)
	} else {
		dbHost := os.Getenv("MONGODB_DB_HOST")
		Session, err = mgo.Dial(dbHost)
	}
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	Session.SetSocketTimeout(20 * time.Second)
	Session.SetSyncTimeout(20 * time.Second)
}

// DB Return MongoDB `festinare_api` Database Connection
func DB() *mgo.Database {
	return Session.DB("festinare_api")
}

// GetUserCount Return total users count
func GetUserCount() int64 {
	// C returns a value representing the named collection.
	// Creating this value is a very lightweight operation, and involves no network communication.
	collection := DB().C("users")
	numberOfUsers, err := collection.Count()
	if err != nil {
		// TODO: Log Error
		return 0
	}
	return int64(numberOfUsers)
}

// GetClientCount Return total clients count
func GetClientCount() int64 {
	// C returns a value representing the named collection.
	// Creating this value is a very lightweight operation, and involves no network communication.
	collection := DB().C("clients")
	numberOfClients, err := collection.Count()
	if err != nil {
		// TODO: Log Error
		return 0
	}
	return int64(numberOfClients)
}
