package models

import (
"strings"
"strconv"
)


type IScore interface {
	Status() (string, error)
	GetPoint(key string) error
	Reset() error
	GetUserName(key string) string
	SetUserName(key, username string) error
	ClearUserName(key string) error
}

func NewScore() IScore {
	return &score{
		board: map[string]int{},
		usernames: map[string]string{},
	}
}

type score struct {
	board map[string]int
	usernames map[string]string
}

func (s *score) Status() (string, error) {
	scores := []string{}
	for k, v := range s.board {
		scores = append(scores, s.usernames[k]+ ": "  + strconv.Itoa(v))
	}

	return strings.Join(scores, ", "), nil
}

func (s *score) GetPoint(key string) error {
	s.board[key] += 2
	return nil
}

func (s *score) Reset() error {
	s.board = map[string]int{}
	return nil
}

func (s *score) GetUserName(key string) string {
	return s.usernames[key]
}

func (s *score) SetUserName(key, username string) error {
	s.usernames[key] = username
	return nil
}

func (s *score) ClearUserName(key string) error {
	delete(s.usernames, key)
	return nil
}
