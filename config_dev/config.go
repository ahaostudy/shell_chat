package config

type Config struct {
	BaseUrl  string
	GptModel string
	ApiKey   string
	Stream   bool
	Colors   Colors
	Prefix   string
	Suffix   string
}

type Colors struct {
	Default string
	Yellow  string
	Green   string
	Blue    string
	Red     string
	Purple  string
}

var Cfg = Config{
	BaseUrl:  "https://api.openai.com",
	GptModel: "gpt-3.5-turbo",
	ApiKey:   "sk-xxx",
	Stream:   true,
	Colors: Colors{
		Default: "\033[0m",
		Yellow:  "\033[33m",
		Green:   "\033[32m",
		Blue:    "\033[38;5;117m",
		Red:     "\033[31m",
		Purple:  "\033[38;5;99m",
	},
	Prefix: "***",
	Suffix: "***",
}

// InitConfig dynamic initialization
func InitConfig() {
}
