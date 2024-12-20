package teststore

import (
	"web/internal/app/model"
	"web/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
	// topicRepository *TopicRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
