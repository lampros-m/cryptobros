package config

type Config struct {
	DebugMode bool // Indicates if the debug mode is enabled
}

func NewConfig(debugMode bool) *Config {
	return &Config{
		DebugMode: debugMode,
	}
}
