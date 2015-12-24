package main

import (
	"testing"
)

func TestNewClientOAuth(t *testing.T) {
	setEnv("")
    NewClientOAuth("")
}

func TestNewClientOAuth_Enterprise(t *testing.T) {
	endpoint := "https://github.enterprise.com/api/v3/"
    NewClientOAuth("")
}
