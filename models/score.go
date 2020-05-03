package models

import (
"strings"
"strconv"
)


type IScore interface {
	Status() (string, error)
	GetPoint(key string) error
	Reset() error
}

func NewScore() IScore {
	return &score{
		board: map[string]int{},
	}
}

type score struct {
	board map[string]int
}

func (s *score) Status() (string, error) {
	scores := []string{}
	for k, v := range s.board {
		scores = append(scores, "player-" + k + ": "  + strconv.Itoa(v))
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
