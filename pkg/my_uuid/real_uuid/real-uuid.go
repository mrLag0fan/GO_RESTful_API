package real_uuid

import (
	"GO_RESTful_API/pkg/logger"
	"github.com/google/uuid"
)

type RealUUIDGenerator struct{}

func (r *RealUUIDGenerator) GenerateUUID() string {
	logger.Log("trace", "New real UUID was created.")
	return uuid.New().String()
}
