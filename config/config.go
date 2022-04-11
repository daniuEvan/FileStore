/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// ServerConfig 全局配置
type ServerConfig struct {
	Host         string         `mapstructure:"host" json:"host"`
	Port         int            `mapstructure:"port" json:"port"`
	Language     string         `mapstructure:"language" json:"language"`
	SmInfo       SmConfig       `mapstructure:"sm" json:"sm"`
	LogInfo      LogConfig      `mapstructure:"log" json:"log"`
	AuthInfo     AuthConfig     `mapstructure:"auth" json:"auth"`
	DatabaseInfo DatabaseConfig `mapstructure:"database" json:"database"`
}
