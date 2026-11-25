package cli

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
	default:
		return "Unknown command", nil
	}
}

func formatVersion() string {
	return "Forge version: " + Version
}
