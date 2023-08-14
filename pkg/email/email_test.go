package email

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
)

func TestSendValidCode(t *testing.T) {
	config.InitConfig("../..")
	t.Log(config.Conf.System.Host)
	t.Log(config.Conf.SES.SecretKey)
	err := SendValidCode("ruidongli2002@gmail.com", "1234")
	if err != nil {
		t.Error(err)
	}
}
