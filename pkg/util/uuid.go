package util

import "github.com/google/uuid"

// GetUUID
// Get uuid
// 获得UUID码
func GetUUID() string {
	uid := uuid.New()
	return uid.String()
}
