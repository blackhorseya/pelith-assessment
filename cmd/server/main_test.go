//go:build external

package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	viper.Set("config", "../../deployments/dev/config.yaml")
	server, clean, err := NewCmd(viper.GetViper())
	if err != nil {
		t.Error(err)
	}
	defer clean()

	ctx := context.Background()
	err = server.Start(ctx)
	if err != nil {
		t.Error(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan

	err = server.Shutdown(ctx)
	if err != nil {
		t.Error(err)
	}
}
