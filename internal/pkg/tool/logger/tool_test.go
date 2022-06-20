package logger

import (
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTool_Debug(t *testing.T) {
	assert.Equal(t, "debug", build.RunMode())
}

func TestTool_viper(t *testing.T) {
	path := config.RootPath() + "/config.yaml"
	vp := viper.New()
	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		t.Fatalf(err.Error())
	}
}
