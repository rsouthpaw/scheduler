package base

import (
	"crypto/rand"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	file_name_timestamp_format = "212006-150405"
)

func setupServer(c Config) (string, error) {
	return setupVariables(c), nil
}
func setupVariables(config Config) string {

	startLog := fmt.Sprintln("\n\nStarting...")
	startLog += fmt.Sprintln("Version", SERVER_VERSION)
	RUNNING_SERVER_TYPE = config
	if config == SERVER_TYPE_LOCALHOST {
		PORT = ":9010"
		BASE_URL = "http://localhost" + PORT
		startLog += fmt.Sprintln("Localhost")
		startLog += fmt.Sprintln("Port", PORT)
		MONGO_BASE_URL = "localhost:27017"

	} else if config == SERVER_TYPE_DEVELOPMENT {
		PORT = ":9010"
		BASE_URL = "https://actyv.humanite.in" + PORT
		startLog += fmt.Sprintln("Development")
		startLog += fmt.Sprintln("Port", PORT)
		MONGO_BASE_URL = "localhost:27017"

	} else if config == SERVER_TYPE_PRODUCTION {
		PORT = ":9010"
		BASE_URL = "https://actyv.humanite.in" + PORT
		startLog += fmt.Sprintln("Production")
		startLog += fmt.Sprintln("Port", PORT)
		MONGO_BASE_URL = "localhost:27017"
	} else {
		startLog += fmt.Sprintln("Exiting...", "invalid config type")
	}

	return startLog
}
func randStr() string {
	var dictionary = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 10)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
func generateNewFileName() string {
	return time.Now().Format(file_name_timestamp_format)
}
