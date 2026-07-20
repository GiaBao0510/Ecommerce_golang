package setting

// Cấu trúc cấu hình (config) chính của ứng dụng, được ánh xạ từ file YAML.
type Config struct {
	PostgreSQL PostgreSQLConfig `mapstructure:"postgres"`
	Logger     LoggerSetting    `mapstructure:"log"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Server     ServerConfig     `mapstructure:"server"`
	Cors       CORS_Config      `mapstructure:"cors"`
}

// Cấu trúc con cho phần cấu hình server (port, host, mode)
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

// Cấu trúc con cho phần cấu hình Database là PostgreSQL
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

// Cấu trúc con cho phần cấu hình logger
type LoggerSetting struct {
	Loglevel       string `mapstructure:"log_level"`
	LogAccessFile  string `mapstructure:"log_access_file"`
	LogErrorFile   string `mapstructure:"log_error_file"`
	LogAppFile     string `mapstructure:"log_app_file"`
	LogWarningFile string `mapstructure:"log_warning_file"`
	MaxSize        int    `mapstructure:"maxSize"`
	MaxBackups     int    `mapstructure:"maxBackups"`
	MaxAge         int    `mapstructure:"maxAge"`
	Compress       bool   `mapstructure:"compress"`
	LogFormat      string `mapstructure:"log_format"` // Xác định định dạng log: "json" hoặc "console"
}

// Cấu trúc con cho phần cấu hình caches là Redis
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

// Cấu trúc CORS
type CORS_Config struct {
	Allowed_origins   []string `mapstructure:"allowed_origins"`
	Allowed_methods   []string `mapstructure:"allowed_methods"`
	Allowed_headers   []string `mapstructure:"allowed_headers"`
	Allow_credentials bool     `mapstructure:"allow_credentials"`
	Max_age           int      `mapstructure:"max_age"`
}
