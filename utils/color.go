package utils

import (
	"shell_chat/config"
)

func ColorString(text, color string) string {
	return color + text + config.Cfg.Colors.Default
}
