package config

import "testing"

func TestConfigSystem(t *testing.T) {
	InitConfig()
	if Conf.System.Port != "5001" {
		t.Errorf("Expected 5001 but found %s", Conf.System.Port)
	}
	if Conf.System.Host != "localhost" {
		t.Errorf("Expected localhost but found %s", Conf.System.Host)
	}
}

func TestConfigDb(t *testing.T) {
	InitConfig()
	if Conf.DB.Host != "db" {
		t.Errorf("Expected db but found %s", Conf.DB.Host)
	}
	if Conf.DB.Port != "3306" {
		t.Errorf("Expected 3306 but found %s", Conf.DB.Port)
	}
	if Conf.DB.Charset != "utf8mb4" {
		t.Errorf("Expected utf8mb4 but found %s", Conf.DB.Charset)
	}
	if Conf.DB.UserName != "root" {
		t.Errorf("Expected root but found %s", Conf.DB.UserName)
	}
	if Conf.DB.Password != "TuringCupBackend123!@#" {
		t.Errorf("Expected TuringCupBackend123!@# but found %s", Conf.DB.Password)
	}
}
