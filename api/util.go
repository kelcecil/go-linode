package api

import (
	"os"
)

func GetAPIKeyFromEnvironment() string {
	return os.Getenv("LINODE_KEY")
}
