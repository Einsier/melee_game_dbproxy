package configs

const (
	// mongodb
	MongoPoolSize   uint64 = 100
	DBName          string = "melee_game"
	MongoURI        string = "mongodb://localhost:27017"
	MongoURIForTest string = "mongodb://localhost:27017"

	// tcp port
	TcpPort string = ":8890"
)
