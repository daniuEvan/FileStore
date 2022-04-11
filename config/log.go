/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// LogConfig 日志配置
type LogConfig struct {
	LogLevel     string `mapstructure:"logLevel" json:"logLevel"`
	LogPath      string `mapstructure:"logPath" json:"logPath"`
	LogInConsole bool   `mapstructure:"logInConsole" json:"logInConsole"`
	MaxSize      int    `mapstructure:"maxSize" json:"maxSize"`
	MaxBackups   int    `mapstructure:"maxBackups" json:"maxBackups"`
	MaxAge       int    `mapstructure:"maxAge" json:"maxAge"`
	Compress     bool   `mapstructure:"compress" json:"compress"`
}
