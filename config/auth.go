/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// JWTConfig 配置信息
type JWTConfig struct {
	SigningKey    string `mapstructure:"signingKey" json:"signingKey"`
	TokenKey      string `mapstructure:"tokenKey" json:"tokenKey"`
	EffectiveTime int    `mapstructure:"effectiveTime" json:"effectiveTime"`
}

// LDAPConfig ldap 设置
type LDAPConfig struct {
	LdapHost       string `mapstructure:"ldapHost" json:"ldapHost"`
	LdapPort       int    `mapstructure:"ldapPort" json:"ldapPort"`
	BaseDN         string `mapstructure:"baseDN" json:"baseDN"`
	SearchProperty string `mapstructure:"searchProperty" json:"searchProperty"`
	BindDN         string `mapstructure:"bindDN" json:"bindDN"`
	BindPassword   string `mapstructure:"bindPassword" json:"bindPassword"`
}

// AuthConfig Auth配置
type AuthConfig struct {
	JWTInfo  JWTConfig  `mapstructure:"jwt" json:"jwt"`
	LADPInfo LDAPConfig `mapstructure:"ldap" json:"ldap"`
}
