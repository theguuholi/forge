package cli_test

import (
	"testing"

	"github.com/theguuholi/forge/internal/cli"
)

func TestRun(t *testing.T) {
	t.Run("version returns only the version", func(t *testing.T) {
		got, err := cli.Run([]string{"version"})

		if err != nil {
			t.Fatalf("Run(version) returned unexpected error: %v", err)
		}

		want := "Forge version: " + cli.Version

		if got != want {
			t.Errorf("Run(version) = %q; want %q", got, want)
		}
	})

	t.Run("help returns available commands", func(t *testing.T) {
		got, err := cli.Run([]string{"help"})

		if err != nil {
			t.Fatalf("Run(help) returned unexpected error: %v", err)
		}

		want := "Commandos disponiveis: version, help"

		if got != want {
			t.Errorf("Run(help) = %q; want %q", got, want)
		}
	})

	t.Run("unknown command returns error", func(t *testing.T) {
		got, err := cli.Run([]string{"unknown"})

		if err == nil {
			t.Fatal("expected an error for an unknown command")
		}

		if got != "" {
			t.Errorf("Run(unknown) output = %q; want empty", got)
		}

		want := `unknown command: "unknown"`
		if err.Error() != want {
			t.Errorf("Run(unknown) error = %q; want %q", err.Error(), want)
		}
	})

	t.Run("no arguments returns help", func(t *testing.T) {
		got, err := cli.Run(nil)

		if err != nil {
			t.Fatalf("Run(nil) returned unexpected error: %v", err)
		}

		want := "Forge version: 0.0.1\nCommandos disponiveis: version, help"

		if got != want {
			t.Errorf("Run(nil) = %q; want %q", got, want)
		}
	})

	t.Run("return about", func(t *testing.T) {
		got, err := cli.Run([]string{"about"})

		if err != nil {
			t.Fatalf("Run([]string{\"about\"}) returned unexpected error: %v", err)
		}

		want := "Forge: plataforma para publicar e administrar sites estaticos."

		if got != want {
			t.Errorf("Run([]string{\"about\"}) = %q; want %q", got, want)
		}
	})
}
