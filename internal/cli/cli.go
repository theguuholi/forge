package cli

import "fmt"

const Version = "0.0.1"

func Run(args []string) (string, error) {
	if len(args) == 0 {
		text := formatVersion() + "\n" + helperText()
		return text, nil
	}

	switch args[0] {
	case "version":
		return formatVersion(), nil
	case "help":
		return helperText(), nil
	case "about":
		return about(), nil
	default:
		return "", fmt.Errorf("unknown command: %q", args[0])
	}
}

func formatVersion() string {
	return "Forge version: " + Version
}

func about() string {
	return "Forge: plataforma para publicar e administrar sites estaticos."
}
