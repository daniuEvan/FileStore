/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// OrmDatabasePoolConfig 配置
type OrmDatabasePoolConfig struct {
	Status          string `mapstructure:"status" json:"status"` // enable 开启数据库连接池 disable 不启用数据库连接池
	MaxIdleConns    int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"`
}

// MysqlConfig mysql 配置
type MysqlConfig struct {
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

// PostgresConfig pg 配置
type PostgresConfig struct {
	Schema   string `mapstructure:"schema" json:"schema"`
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host                string `mapstructure:"host" json:"host"`
	Port                int    `mapstructure:"port" json:"port"`
	DB                  int    `mapstructure:"db" json:"db"`
	Username            string `mapstructure:"username" json:"username"`
	Password            string `mapstructure:"password" json:"password"`
	ConnectTimeout      int    `mapstructure:"connectTimeout" json:"connectTimeout"`
	PoolMaxIdleConns    int    `mapstructure:"poolMaxIdleConns" json:"poolMaxIdleConns"`
	PoolMaxOpenConns    int    `mapstructure:"poolMaxOpenConns" json:"poolMaxOpenConns"`
	PoolConnMaxLifetime int    `mapstructure:"poolConnMaxLifetime" json:"poolConnMaxLifetime"`
}

type DatabaseConfig struct {
	DBType              string                `mapstructure:"dbType" json:"dbType"`
	TablePrefix         string                `mapstructure:"tablePrefix" json:"tablePrefix"`
	MysqlInfo           MysqlConfig           `mapstructure:"mysql" json:"mysql"`
	PostgresInfo        PostgresConfig        `mapstructure:"postgres" json:"postgres"`
	RedisInfo           RedisConfig           `mapstructure:"redis" json:"redis"`
	OrmDatabasePoolInfo OrmDatabasePoolConfig `mapstructure:"ormDatabasePool" json:"ormDatabasePool"`
}
