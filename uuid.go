package drbot

import (
	"fmt"
	"math/rand/v2"
)

const (
	uuidFormat = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
)

func GetNewUUID() string {
	uuid := []byte(uuidFormat)

	for i := range uuid {
		switch uuid[i] {
		case 'x':
			uuid[i] = byte(rand.IntN(16))
		case 'y':
			uuid[i] = byte(rand.IntN(4) | 8)
		}
	}

	return fmt.Sprintf("%s", uuid)
}
