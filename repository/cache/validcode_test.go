package cache

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
)

func TestSendValidCode(t *testing.T) {
	config.InitConfig("../..")
	InitCache()
	code, err := GenerateValidCode()
	t.Log(code)
	if err != nil {
		t.Log(err)
	}
	err = CheckValidCode(code)
	if err != nil {
		t.Error(err)
	}
}
