package staterepo

import "alfachatback/internal/chat/repository/models"

type stateStore struct {
	states map[int]models.State
}

func NewStateDB() *stateStore {
	return &stateStore{
		states: make(map[int]models.State),
	}
}
