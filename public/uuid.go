package public

import "github.com/google/uuid"

func GetUUid() string {
	newUUID, _ := uuid.NewRandom()
	return newUUID.String()
}
