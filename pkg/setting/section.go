package setting

// Cấu trúc cấu hình (config) chính của ứng dụng, được ánh xạ từ file YAML.
type Config struct {
	PostgreSQL     PostgreSQLConfig     `mapstructure:"postgres"`
	Logger         LoggerSetting        `mapstructure:"log"`
	Redis          RedisConfig          `mapstructure:"redis"`
	Server         ServerConfig         `mapstructure:"server"`
	Cors           CORS_Config          `mapstructure:"cors"`
	Authentication AuthenticationConfig `mapstructure:"authentication"`
	RateLimit      RateLimitConfig      `mapstructure:"rate_limit"`
	CronJob        CronJob_config       `mapstructure:"cronjob"`
}

// Cấu hình cho Authentication, bao gồm các thông tin liên quan khác.
type AuthenticationConfig struct {
	MailJet    MailJetConfig    `mapstructure:"mailjet"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Cloudflare CloudflareConfig `mapstructure:"cloudflare"`
	OAuth2     OAuth2Config     `mapstructure:"oauth2"`
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

// Cấu truccs kết nối mailinject
type MailJetConfig struct {
	API_key    string `mapstructure:"api_key"`
	Secret_key string `mapstructure:"secret_key"`
	From_mail  string `mapstructure:"from_mail"`
	From_name  string `mapstructure:"from_name"`
	App_url    string `mapstructure:"app_url"`
}

// Cấu trúc CORS
type CORS_Config struct {
	Allowed_origins   []string `mapstructure:"allowed_origins"`
	Allowed_methods   []string `mapstructure:"allowed_methods"`
	Allowed_headers   []string `mapstructure:"allowed_headers"`
	Allow_credentials bool     `mapstructure:"allow_credentials"`
	Max_age           int      `mapstructure:"max_age"`
}

// cấu trúc cho các cronjob
type CronJob_config struct {
	Backup_cron           string `mapstructure:"backup_cron"`
	Backup_retention_days int    `mapstructure:"backup_retention_days"`
	Backup_dir            string `mapstructure:"backup_dir"`
}

type JWTConfig struct {
	Secret                       string `mapstructure:"secret"`
	Issuer                       string `mapstructure:"issuer"`
	Audience                     string `mapstructure:"audience"`
	AccessTokenExpirationMinutes int    `mapstructure:"AccessTokenExpirationMinutes"`
	RefreshTokenExpirationDays   int    `mapstructure:"RefreshTokenExpirationDays"`
	EncrypKey                    string `mapstructure:"encrypKey"`
}

type CloudflareConfig struct {
	TokenName       string `mapstructure:"r2_token_name"`
	BucketName      string `mapstructure:"r2_bucket_name"`
	TokenValue      string `mapstructure:"r2_token_value"`
	AccountID       string `mapstructure:"r2_account_id"`
	AccessKeyID     string `mapstructure:"r2_access_key_id"`
	SecretAccessKey string `mapstructure:"r2_secret_access_key"`
	Endpoint        string `mapstructure:"r2_endpoint"`
}

/*================ RATE LIMITER ============= */
type RateLimitConfig struct {
	PerClient PerClientConfig `mapstructure:"per_client"`
}

type PerClientConfig struct {
	Enabled              bool   `mapstructure:"enabled"`
	Key_by               string `mapstructure:"key_by"`
	Fallback_to_ip       bool   `mapstructure:"fallback_to_ip"`
	Unauthenticated_tier string `mapstructure:"unauthenticated_tier"`
	Request_sec_public   int    `mapstructure:"request_sec_public"`
	Burst_public         int    `mapstructure:"burst_public"`
	Request_sec_private  int    `mapstructure:"request_sec_private"`
	Burst_private        int    `mapstructure:"burst_private"`
}

/* =========== OAUTH2 ===============*/
type OAuth2Config struct {
	Google GoogleOAuth2Config `mapstructure:"google"`
}

type GoogleOAuth2Config struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURL  string `mapstructure:"redirect_url"`
}
