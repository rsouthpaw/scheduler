package base

const (
	// Server Configurations
	SERVER_TYPE_LOCALHOST   Config = 1
	SERVER_TYPE_DEVELOPMENT Config = 2
	SERVER_TYPE_PRODUCTION  Config = 3

	// Server Details
	SERVER_VERSION = "1.0.0"
	API_KEY        = "vXBk4iN9JPNjjpNKWnBmlgjqjScoX3GQ"

	// Database Constants
	DB_ACTYV  = "scheduler"
	COL_USERS = "users"
	COL_TASKS = "tasks"
)

var (
	RUNNING_SERVER_TYPE Config
	BASE_URL            string
	PORT                string
	MONGO_BASE_URL      string
)

type Config int

func SetupServer(c Config) (string, error) {
	return setupServer(c)
}
