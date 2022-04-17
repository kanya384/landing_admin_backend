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
	ctx := context.Background()
	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login_for_get_test", Pass: "password", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}
	err := s.repository.Users.Create(ctx, user)
	if err != nil {
		s.FailNow("error creating user for get users list tests", err)
		return
	}

	cases := map[string]struct {
		empty bool
		err   interface{}
	}{
		"get users list": {
			empty: false,
			err:   nil,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			users, err := s.repository.Users.Get(ctx)
			s.NotEmpty(users)
			if !cs.empty {
				s.Equal(err, cs.err)
			}
		})
	}
}

func (s *UsersTestSuite) TestGetUserByCredetinals() {
	ctx := context.Background()
	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login_for_get_by_cred_test", Pass: "password", Role: "admin"}
	err := s.repository.Users.Create(ctx, user)
	if err != nil {
		s.FailNow("error creating user for get user by credetinasl test", err)
		return
	}

	cases := map[string]struct {
		input map[string]string
		want  interface{}
		err   interface{}
	}{
		"get user success": {
			input: map[string]string{"login": "login_for_get_by_cred_test", "pass": "password"},
			want:  domain.UserInfo{ID: user.ID, Name: user.Name, Login: user.Login, Role: user.Role},
			err:   nil,
		},
		"get user error": {
			input: map[string]string{"login": "login_for_get_by_cred_test", "pass": "password2"},
			want:  domain.UserInfo{},
			err:   domain.ErrCheckUserAndPass,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			user, err := s.repository.Users.GetUserByCredetinals(ctx, cs.input["login"], cs.input["pass"])
			s.Equal(cs.want, user)
			if cs.err != nil {
				s.ErrorContains(err, cs.err.(string))
			} else {
				s.Equal(err, nil)
			}
		})
	}
}

func (s *UsersTestSuite) TestCreate() {
	ctx := context.Background()
	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login", Pass: "password", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}

	//for duplicate test
	userCreated := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login2", Pass: "password", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}
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

func (s *UsersTestSuite) TestUpdate() {
	ctx := context.Background()

	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login_update", Pass: "password", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}
	err := s.repository.Users.Create(ctx, user)
	if err != nil {
		s.FailNow("error creating user for update tests", err)
		return
	}
	cases := map[string]struct {
		input domain.User
		want  string
		err   interface{}
	}{
		"update user success": {
			input: domain.User{ID: user.ID, Name: "New name", Login: "new login", Pass: "pass new", Role: "admin", CreatedAt: user.CreatedAt, UpdateAt: user.UpdateAt, ModifiedBy: user.ModifiedBy},
			err:   nil,
		},
		"update err not founded": {
			input: domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login_upd2", Pass: "Passs", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()},
			err:   domain.ErrNoFieldWithID,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			err := s.repository.Users.Update(ctx, cs.input)
			if cs.err != nil {
				s.ErrorContains(err, cs.err.(string))
			} else {
				s.Equal(nil, err)
			}
		})
	}
}

func (s *UsersTestSuite) TestDelete() {
	ctx := context.Background()

	user := domain.User{ID: primitive.NewObjectID(), Name: "Name", Login: "login_update", Pass: "password", Role: "admin", CreatedAt: time.Now(), UpdateAt: time.Now(), ModifiedBy: primitive.NewObjectID()}
	err := s.repository.Users.Create(ctx, user)
	if err != nil {
		s.FailNow("error creating user for delete tests", err)
		return
	}
	cases := map[string]struct {
		input primitive.ObjectID
		want  string
		err   interface{}
	}{
		"delete user success": {
			input: user.ID,
			err:   nil,
		},
		"delete not founded": {
			input: primitive.NewObjectID(),
			err:   domain.ErrNoFieldWithID,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			err := s.repository.Users.Delete(ctx, cs.input)
			if cs.err != nil {
				s.ErrorContains(err, cs.err.(string))
			} else {
				s.Equal(nil, err)
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
