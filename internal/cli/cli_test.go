package cli_test

import (
	"strings"
	"testing"

	"github.com/theguuholi/forge/internal/cli"
)

func TestRun(t *testing.T) {
	t.Run("includes version", func(t *testing.T) {
		got := cli.Run()
		want := "Forge version: " + cli.Version

		if !strings.Contains(got, want) {
			t.Errorf("Run() = %q; want it to contain %q", got, want)
		}
	})

	t.Run("includes available commands", func(t *testing.T) {
		got := cli.Run()
		want := "Commandos disponiveis: version, help"

		if !strings.Contains(got, want) {
			t.Errorf("Run() = %q; want it to contain %q", got, want)
		}
	})
}
