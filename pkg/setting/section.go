package setting

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Logger   LogConfig      `mapstructure:"logger"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type JWTConfig struct {
	Secret         string `mapstructure:"secret"`
	ExpiresMinutes int    `mapstructure:"expires_minutes"`
}

type LogConfig struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
