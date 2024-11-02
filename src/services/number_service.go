package services

import (
	"errors"
	"math/rand/v2"
)

type NumberService struct{}

type INumberService interface {
	GenerateRandomNumber() int32
	Calculate(a int32, b int32) (int32, error)
}

func NewNumberService() INumberService {
	return &NumberService{}
}

func (s *NumberService) GenerateRandomNumber() int32 {
	randomNumber := rand.Int32N(100000000)
	return randomNumber
}

func (s *NumberService) Calculate(a int32, b int32) (int32, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("a and b cannot be zero")
	}
	result := a + b
	return result, nil
}
