package cache

import (
	"testing"

	"github.com/TuringCup/TuringBackend/config"
)

func TestInitCache(t *testing.T) {
	config.InitConfig("../..")
	InitCache()
}
