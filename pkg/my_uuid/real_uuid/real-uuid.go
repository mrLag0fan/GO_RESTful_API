package real_uuid

import "github.com/google/uuid"

type RealUUIDGenerator struct{}

func (r *RealUUIDGenerator) GenerateUUID() string {
	return uuid.New().String()
}
