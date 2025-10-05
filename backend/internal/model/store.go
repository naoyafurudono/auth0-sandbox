package model

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

var (
	ErrNotFound = errors.New("not found")
)

type Store struct {
	users    map[string]*User
	profiles map[string]*UserProfile
	data     map[string][]*UserData
	mu       sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		users:    make(map[string]*User),
		profiles: make(map[string]*UserProfile),
		data:     make(map[string][]*UserData),
	}
}

func (s *Store) GetOrCreateUser(auth0ID, email, name string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Auth0Id == auth0ID {
			return user, nil
		}
	}

	var namePtr *string
	if name != "" {
		namePtr = &name
	}

	user := &User{
		Id:        uuid.New().String(),
		Auth0Id:   auth0ID,
		Email:     openapi_types.Email(email),
		Name:      namePtr,
		CreatedAt: time.Now(),
	}
	s.users[user.Id] = user
	return user, nil
}

func (s *Store) GetUserByAuth0ID(auth0ID string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.Auth0Id == auth0ID {
			return user, nil
		}
	}
	return nil, ErrNotFound
}

func (s *Store) GetUserProfile(userID string) (*UserProfile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	profile, ok := s.profiles[userID]
	if !ok {
		return nil, ErrNotFound
	}
	return profile, nil
}

func (s *Store) CreateOrUpdateUserProfile(userID string, update *UserProfileUpdate) (*UserProfile, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	profile, exists := s.profiles[userID]
	if !exists {
		profile = &UserProfile{
			Id:     uuid.New().String(),
			UserId: userID,
		}
	}

	profile.DisplayName = update.DisplayName
	profile.Bio = update.Bio
	profile.AvatarUrl = update.AvatarUrl
	profile.UpdatedAt = time.Now()

	s.profiles[userID] = profile
	return profile, nil
}

func (s *Store) GetUserData(userID string) ([]*UserData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.data[userID]
	if !ok {
		return []*UserData{}, nil
	}
	return data, nil
}

func (s *Store) CreateUserData(userID string, create *UserDataCreate) (*UserData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := &UserData{
		Id:        uuid.New().String(),
		UserId:    userID,
		Content:   create.Content,
		CreatedAt: time.Now(),
	}

	s.data[userID] = append(s.data[userID], data)
	return data, nil
}
