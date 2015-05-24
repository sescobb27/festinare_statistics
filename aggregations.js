// count how many categories have been selected
db.users.aggregate([
  {"$unwind": "$categories"},
  {
    "$group": {
      "_id": {
        "category": "$categories.name"
      },
      "count": { "$sum": 1 }
    }
  },
  {
    "$sort": { "count": -1 }
  }
])

//
db.users.aggregate([
  {"$unwind": "$discounts"},
  {"$unwind": "$discounts.categories"},
  {
    "$group": {
      "_id": {
        "category": "$discounts.categories.name"
      },
      "count": { "$sum": 1 }
    }
  },
  {
    "$sort": { "count": -1 }
  }
])
