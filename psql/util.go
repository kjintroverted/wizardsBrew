package psql

import "github.com/google/uuid"

// GetUID generates a UID
func GetUID() string {
	uuid := uuid.New()
	return uuid.String()
}
