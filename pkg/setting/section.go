package setting

type Config struct {
	PostgreSQL PostgreSQLConfig `mapstructure:"postgres"`
	Logger     LoggerSetting    `mapstructure:"log"`
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

type Redis struct {
	Host string `mapstructure:"host"`
	Port int   `mapstructure:"port"`
	Password string `mapstructure:"password"`
	IdleTimeout string `mapstructure:"idleTimeout"`
	MaxConnLifetime string `mapstructure:"maxConnLifetime"`
	WaitTimeout string `mapstructure:"waitTimeout"`
}