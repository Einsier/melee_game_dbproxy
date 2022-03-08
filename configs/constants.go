package configs

var (
	// mongodb
	MongoPoolSize   uint64 = 100
	DBName          string = "melee_game"
	MongoURI        string
	MongoURIForTest string = "mongodb://localhost:27017"

	// tcp port
	TcpPort string = ":1234"
)
