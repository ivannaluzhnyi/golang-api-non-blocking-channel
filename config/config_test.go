package config

import "testing"

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	if config.Port != ":4000" {
		t.Errorf("Port not good")
	}
}
