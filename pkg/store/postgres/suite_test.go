package postgres_test

import (
	"GO_RESTful_API/pkg/store/postgres"
	"context"
	"testing"

	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"

	"github.com/stretchr/testify/suite"
)

var (
	conf config.PostgresConfig
)

type StoreSuite struct {
	suite.Suite

	DB *postgres.PostgresStore

	clearTables []string
}

func TestSuite(t *testing.T) {
	if err := env.Parse(&conf); err != nil {
		// logger.Errorf("store_suite_test.go", "TestSuite", "env.Parse", err)
	}

	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {
	ctx := context.Background()

	db, err := postgres.NewStore(context.Background(), conf)
	s.Nil(err)
	s.DB = db

	s.clearTables = []string{}

	s.cleanDB()
}

func (s *StoreSuite) BeforeTest() {
	s.cleanDB()
}

func (s *StoreSuite) TearDownTest() {
	s.cleanDB()
}

func (s *StoreSuite) TearDownSuite() {
	s.cleanDB()
}

func (s *StoreSuite) cleanDB() {
	for _, v := range s.clearTables {
		_, err := s.DB.Conn().Exec(context.Background(), "TRUNCATE "+v+";")
		if err != nil {
			logger.Errorf("store_suite_test.go", "cleanDB", "Exec", err)
		}
	}
}
