package email

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
)

func TestSendValidCode(t *testing.T) {
	config.InitConfig("../..")
	t.Log(config.Conf.SES.Id)
	t.Log(config.Conf.SES.Key)
	err := SendValidCode("ruidongli2002@gmail.com", "1234")
	if err != nil {
		t.Error(err)
	}
}
