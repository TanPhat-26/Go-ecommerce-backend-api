package setting

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App    APPSetting
	DB     DBSetting
	Redis  RedisSetting
	JWT    JWTSetting
	Logger LoggerSetting
}

type APPSetting struct {
	Env  string
	Name string
	Port string
}

type DBSetting struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLmode  string
	Timezone string
}

type RedisSetting struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTSetting struct {
	AccessSecret     string
	RefreshSecret    string
	AccessTTLMinutes int
	RefreshTTLHours  int
}

type LoggerSetting struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

func LoadConfig(configPath ...string) (*Config, error) {
	v := viper.New()
	if len(configPath) > 0 && configPath[0] != "" {
		v.SetConfigFile(configPath[0])
	} else {
		v.SetConfigFile(".env")
	}
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Read config file failed: %w", err)
	}

	config := &Config{
		App: APPSetting{
			Env:  v.GetString("APP_ENV"),
			Name: v.GetString("APP_NAME"),
			Port: v.GetString("APP_PORT"),
		},
		DB: DBSetting{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			Name:     v.GetString("DB_NAME"),
			SSLmode:  v.GetString("DB_SSLMODE"),
			Timezone: v.GetString("DB_TIMEZONE"),
		},
		Redis: RedisSetting{
			Host:     v.GetString("REDIS_HOST"),
			Port:     v.GetString("REDIS_PORT"),
			Password: v.GetString("REDIS_PASSWORD"),
			DB:       v.GetInt("REDIS_DB"),
		},
		JWT: JWTSetting{
			AccessSecret:     v.GetString("JWT_ACCESS_SECRET"),
			RefreshSecret:    v.GetString("JWT_REFRESH_SECRET"),
			AccessTTLMinutes: v.GetInt("JWT_ACCESS_TTL_MINUTES"),
			RefreshTTLHours:  v.GetInt("JWT_REFRESH_TTL_HOURS"),
		},
		Logger: LoggerSetting{
			Level:    v.GetString("LOG_LEVEL"),
			Format:   v.GetString("LOG_FORMAT"),
			Output:   v.GetString("LOG_OUTPUT"),
			FilePath: v.GetString("LOG_FILE_PATH"),
		},
	}
	return config, nil
}

func (d DBSetting) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
		d.Host,
		d.User,
		d.Password,
		d.Name,
		d.Port,
		d.SSLmode,
		d.Timezone,
	)
}

func (r RedisSetting) Addr() string {
	return strings.Join([]string{r.Host, r.Port}, ":")
}
