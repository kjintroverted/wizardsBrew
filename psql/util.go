package psql

import "github.com/google/uuid"

func GetUID() string {
	uuid := uuid.New()
	return uuid.String()
}
