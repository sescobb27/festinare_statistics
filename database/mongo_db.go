package database

import (
	"os"
	"time"

	"github.com/sescobb27/festinare_statistics/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func UserCollection() *mgo.Collection {
	return DB().C("users")
}

func ClientCollection() *mgo.Collection {
	return DB().C("clients")
}

// GetUserCount Return total users count
func GetUserCount() int64 {
	// C returns a value representing the named collection.
	// Creating this value is a very lightweight operation, and involves no network communication.
	numberOfUsers, err := UserCollection().Count()
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
	numberOfClients, err := ClientCollection().Count()
	if err != nil {
		// TODO: Log Error
		return 0
	}
	return int64(numberOfClients)
}

// GetAvailableDiscounts Returns all discounts based on categories preferences
// if number > 0 is set as a limit if not then it doesn't do anything
func GetAvailableDiscounts(number int, categories []*models.Category) []*models.Discount {
	names := make([]string, len(categories))
	for index, category := range categories {
		names[index] = category.Name
	}
	query := ClientCollection().
		Find(bson.M{"categories.name": bson.M{"$in": names}})

	if number > 0 {
		query = query.Limit(number)
	}

	iterator := query.Iter()
	discounts := []*models.Discount{}
	var client models.Client
	for iterator.Next(&client) {
		discounts = append(discounts, client.Discounts...)
	}

	if err := iterator.Close(); err != nil {
		panic(err)
	}

	return discounts
}

// CategoryPreference
type CategoryPreference struct {
	Category *models.Category
	Count    int
}

// UserCategories count how many categories have been selected (preferences)
func UserCategoryPreference() []*CategoryPreference {
	// db.users.aggregate([
	//   {"$unwind": "$categories"},
	//   {
	//     "$group": {
	//       "_id": {
	//         "category": "$categories.name"
	//       },
	//       "count": { "$sum": 1 }
	//     }
	//   },
	//   {
	//     "$sort": { "count": -1 }
	//   }
	// ])
	iterator := UserCollection().Pipe([]bson.M{
		{"$unwind": "$categories"},
		{
			"$group": bson.M{
				"_id":   bson.M{"category": "$categories.name"},
				"count": bson.M{"$sum": 1},
			},
		},
		{"$sort": bson.M{"count": -1}},
	}).Iter()

	// result -> map[_id:map[category:Disco] count:353]
	var result map[string]interface{}
	preferences := []*CategoryPreference{}

	for iterator.Next(&result) {
		query, _ := result["_id"].(map[string]interface{})
		preference := &CategoryPreference{
			Category: &models.Category{Name: query["category"].(string)},
			Count:    result["count"].(int),
		}
		preferences = append(preferences, preference)
	}

	if err := iterator.Close(); err != nil {
		panic(err)
	}

	return preferences
}
