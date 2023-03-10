package config

import "os"

type Config struct {
	Mode       string
	OutputFile string
}

// If the provided environment variable exists and is not empty its value will
// be returned. Otherwise the provided default value will be returned.
func GetEnvWithDefault(env string, defaultValue string) string {
	value, exists := os.LookupEnv(env)

	if exists && value != "" {
		return value
	}

	return defaultValue
}

func ParseConfig() *Config {
	config := Config{
		Mode:       GetEnvWithDefault("GIFTK_MODE", "generate"),
		OutputFile: GetEnvWithDefault("GIFTK_OUTPUT_FILE", "./output.gif"),
	}
	return &config
}
