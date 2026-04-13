package setting

type Config struct {
	PostgreSQL PostgreSQLConfig `mapstructure:"postgres"`
	Logger     LoggerSetting    `mapstructure:"log"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Server     ServerConfig     `mapstructure:"server"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type PostgreSQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Loglevel   string `mapstructure:"log_level"`
	LogFile    string `mapstructure:"log_file"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisConfig struct {
	Address         string `mapstructure:"address"`
	Port            string `mapstructure:"port"`
	Password        string `mapstructure:"password"`
	DB              int    `mapstructure:"db"`
	IdleTimeout     string `mapstructure:"idleTimeout"`
	MaxConnLifetime string `mapstructure:"maxConnLifetime"`
	WaitTimeout     string `mapstructure:"waitTimeout"`
	ReadTimeout     string `mapstructure:"readTimeout"`
	WriteTimeout    string `mapstructure:"writeTimeout"`
	PoolSize        int    `mapstructure:"poolSize"`
}
