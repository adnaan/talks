package main

import (
	"testing"

	"github.com/adnaan/talks/dependency_injection_july2017/user/mock"
)

func TestService(t *testing.T) {

	mockService := mock.NewService(mockConfig)
}
