package utils

import (
	"os"
	"path/filepath"
)

var (
	KvantHome   string
	KvantConfig string
)

func GetKvantHome() string {
	if KvantHome != "" {
		return KvantHome
	}

	home := os.Getenv("KVANTHOME")

	if home != "" {
		return home
	}

	return os.ExpandEnv(filepath.Join("$HOME", ".kvant"))
}

func GetKvantConfigPath() string {
	if KvantConfig != "" {
		return KvantConfig
	}

	return GetKvantHome() + "/config/config.toml"
}
