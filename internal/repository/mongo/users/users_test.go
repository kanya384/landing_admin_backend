package users_test

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	mng "landing_admin_backend/internal/repository/mongo"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestRepositoryDeviceTypesTestSuite(t *testing.T) {
	suite.Run(t, &UsersTestSuite{})
}

type UsersTestSuite struct {
	suite.Suite
	client     *mongo.Client
	database   *mongo.Database
	repository *repository.Repository
}

func (s *UsersTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		s.FailNow("Failed to create Postgres client: %s", err)
	}
	s.client = client
	s.database = client.Database("Test")
	s.repository = mng.Setup(s.database)
}

func (s *UsersTestSuite) TestGet() {
	//ctx := context.Background()
}

func (s *UsersTestSuite) TestCreate() {
	ctx := context.Background()
	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login", Pass: "password", Role: 0, CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}

	//for diplicate test
	userCreated := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login2", Pass: "password", Role: 0, CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}
	s.repository.Users.Create(ctx, userCreated)
	//

	cases := map[string]struct {
		input domain.User
		want  string
		err   interface{}
	}{
		"create user success": {
			input: user,
			err:   nil,
		},
		"duplicate login": {
			input: userCreated,
			err:   "E11000 duplicate key error collection",
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			err := s.repository.Users.Create(ctx, cs.input)
			if cs.err != nil {
				s.ErrorContains(err, cs.err.(string))
			} else {
				s.Equal(err, nil)
			}
		})
	}
}

func (s *UsersTestSuite) TearDownSuite() {
	if s.database != nil {
		ctx := context.Background()

		if err := s.database.Drop(context.Background()); err != nil {
			s.Fail("Failed to drop MongoDB test database '%s': %s", err)
		}

		if err := s.client.Disconnect(ctx); err != nil {
			s.Fail("Failed to close connection to MongoDB: %s", err)
		}
	}
}
