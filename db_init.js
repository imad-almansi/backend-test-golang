db.createUser(
  {
      user: _getEnv("MONGO_READ_USER"),
      pwd: _getEnv("MONGO_READ_PASSWORD"),
      roles: [
          {
              role: "read",
              db: _getEnv("DB_NAME")
          }
      ]
  }
);

db = db.getSiblingDB(_getEnv("DB_NAME"));

db.createCollection(_getEnv("DB_COLLECTION"));

var data = cat("db.json")
data = JSON.parse(data)

db.getCollection(_getEnv("DB_COLLECTION")).insertMany(data);