db.createUser(
  {
    user: "mock-user",
    pwd: "mock-pwd",
    roles: [
      {
	      role: "readWrite",
	      db: "mock-db"
      }
    ]
  }
)
