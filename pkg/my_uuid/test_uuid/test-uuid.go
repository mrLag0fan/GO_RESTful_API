package test_uuid

type TestUUIDGenerator struct{}

func (t *TestUUIDGenerator) GenerateUUID() string {
	return "test-uuid"
}
