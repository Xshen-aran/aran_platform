package config

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    string `json:"maxsize"`
	MaxAge     string `json:"max_age"`
	MaxBackups string `json:"max_backups"`
}

var Conf = new(LogConfig)
