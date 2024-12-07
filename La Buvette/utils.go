package rpg

import "fmt"

const (
	Reset        = "\x1b[0m"
	Black        = "\x1b[30m"
	Red          = "\x1b[31m"
	Green        = "\x1b[32m"
	Yellow       = "\x1b[33m"
	Blue         = "\x1b[34m"
	Magenta      = "\x1b[35m"
	Cyan         = "\x1b[36m"
	Orange       = "\x1b[33m"
	White        = "\x1b[37m"
	Gray         = "\x1b[90m"
	LightRed     = "\x1b[91m"
	LightGreen   = "\x1b[92m"
	LightYellow  = "\x1b[93m"
	LightBlue    = "\x1b[94m"
	LightMagenta = "\x1b[95m"
	LightCyan    = "\x1b[96m"
	DarkOrange   = "\x1b[38;5;208m"

	Blackground = "\x1b[49m"
	Bold        = "\x1b[1m"
)

func ClearScreen() {
	fmt.Print("\033c")
}
