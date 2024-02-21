package impl

import "GO_RESTful_API/pkg/logger"

type TestUUIDGenerator struct{}

func (t *TestUUIDGenerator) GenerateUUID() string {
	logger.Log("trace", "New test UUID was created.")
	return "test-uuid"
}
