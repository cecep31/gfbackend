package utility

import (
	"os"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/joho/godotenv"
)

// LoadEnvFiles loads environment files in order of precedence
// This follows the convention for managing multiple environments
func LoadEnvFiles() error {
	env := GetEnv("APP_ENV", "development")

	// Load in order of precedence (later files override earlier ones)
	files := []string{
		".env",                    // Base .env file
		".env." + env,             // Environment-specific file (.env.development, .env.production, etc.)
	}

	// Add local files if not in test environment
	if env != "test" {
		files = append(files, ".env.local")
	}
	files = append(files, ".env."+env+".local")

	// Load files (ignore errors for missing files)
	for _, file := range files {
		_ = godotenv.Load(file)
	}

	return nil
}

// GetEnv gets an environment variable with a fallback default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvInt gets an environment variable as integer with a fallback default value
func GetEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// GetEnvBool gets an environment variable as boolean with a fallback default value
func GetEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		value = strings.ToLower(value)
		return value == "true" || value == "1" || value == "yes" || value == "on"
	}
	return defaultValue
}

// GetEnvSlice gets an environment variable as slice (comma-separated) with a fallback default value
func GetEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}

// ConfigHelper provides helper methods for accessing configuration with environment variable support
type ConfigHelper struct{}

// NewConfigHelper creates a new configuration helper
func NewConfigHelper() *ConfigHelper {
	return &ConfigHelper{}
}

// GetString gets a configuration value as string with environment variable fallback
func (c *ConfigHelper) GetString(key, defaultValue string) string {
	// First try to get from GoFrame config
	if value := g.Cfg().MustGet(nil, key).String(); value != "" {
		return value
	}
	// Fallback to environment variable
	return GetEnv(strings.ToUpper(strings.ReplaceAll(key, ".", "_")), defaultValue)
}

// GetInt gets a configuration value as integer with environment variable fallback
func (c *ConfigHelper) GetInt(key string, defaultValue int) int {
	// First try to get from GoFrame config
	if value := g.Cfg().MustGet(nil, key).Int(); value != 0 {
		return value
	}
	// Fallback to environment variable
	return GetEnvInt(strings.ToUpper(strings.ReplaceAll(key, ".", "_")), defaultValue)
}

// GetBool gets a configuration value as boolean with environment variable fallback
func (c *ConfigHelper) GetBool(key string, defaultValue bool) bool {
	// First try to get from GoFrame config
	if value := g.Cfg().MustGet(nil, key); !value.IsEmpty() {
		return value.Bool()
	}
	// Fallback to environment variable
	return GetEnvBool(strings.ToUpper(strings.ReplaceAll(key, ".", "_")), defaultValue)
}

// Example usage functions

// GetDatabaseConfig returns database configuration from environment variables
func GetDatabaseConfig() map[string]string {
	return map[string]string{
		"host":     GetEnv("DB_HOST", "localhost"),
		"port":     GetEnv("DB_PORT", "3306"),
		"user":     GetEnv("DB_USER", "root"),
		"password": GetEnv("DB_PASSWORD", ""),
		"name":     GetEnv("DB_NAME", "gfbackend"),
		"type":     GetEnv("DB_TYPE", "mysql"),
	}
}

// GetJWTConfig returns JWT configuration from environment variables
func GetJWTConfig() map[string]string {
	return map[string]string{
		"secret": GetEnv("JWT_SECRET", "default-secret-change-in-production"),
		"expire": GetEnv("JWT_EXPIRE", "24h"),
		"issuer": GetEnv("JWT_ISSUER", "gfbackend"),
	}
}

// GetServerConfig returns server configuration from environment variables
func GetServerConfig() map[string]interface{} {
	return map[string]interface{}{
		"port":  GetEnvInt("SERVER_PORT", 8000),
		"debug": GetEnvBool("APP_DEBUG", false),
		"env":   GetEnv("APP_ENV", "development"),
	}
}