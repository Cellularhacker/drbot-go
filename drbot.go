package drbot

import (
	"github.com/Cellularhacker/logger-go"
)

var initialized = false

var (
	apiEndpoint      = ""
	chatAPIEndpoint  = ""
	adminAPIEndpoint = ""
)

// SetVariables - Mark: You should have to call this function while initializing step or before run main.
func SetVariables(drbotAPIEndpoint, drbotChatAPIEndpoint, drbotAdminAPIEndpoint string) {
	//	MARK: Check if it initialized
	if drbotAPIEndpoint == "" {
		logger.L.Fatalf("'drbotAPIEndpoint' is missing!")
	}
	apiEndpoint = drbotAPIEndpoint
	if drbotChatAPIEndpoint == "" {
		logger.L.Fatalf("'drbotChatAPIEndpoint' is missing!")
	}
	chatAPIEndpoint = drbotChatAPIEndpoint
	if drbotAdminAPIEndpoint == "" {
		logger.L.Fatalf("'drbotAdminAPIEndpoint' is missing!")
	}
	adminAPIEndpoint = drbotAdminAPIEndpoint

	initialized = true
}

func IsInitialized() bool {
	return initialized
}
