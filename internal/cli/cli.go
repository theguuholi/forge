package cli

const Version = "0.0.1"

func Run() string {
	text := formatVersion() + "\n" + helperText()
	return text
}

func formatVersion() string {
	return "Forge version: " + Version
}
