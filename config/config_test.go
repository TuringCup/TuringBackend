package config

import "testing"

func TestConfigSystem(t *testing.T) {
	InitConfig(".")
	if Conf.System.Port != "5001" {
		t.Errorf("Expected 5001 but found %s", Conf.System.Port)
	}
	if Conf.System.Host != "localhost" {
		t.Errorf("Expected localhost but found %s", Conf.System.Host)
	}
}
