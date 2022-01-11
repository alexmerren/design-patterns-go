package main

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(pinNumber int) error {
	if s.code != pinNumber {
		return errors.New("Pin number does not match")
	}
	fmt.Println("Pin Number Correct")
	return nil
}
