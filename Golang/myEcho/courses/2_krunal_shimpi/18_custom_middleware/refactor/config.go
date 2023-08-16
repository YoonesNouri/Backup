package refactor

//* ConfigDatabase Cofiguration properties based on env variables.
type ConfigDatabase struct {
	AppName  string `env="APP_NAME" env_default="TRONICS"`
	AppEnv   string `env:"APP_ENV" env_default="dev"`
	Port     string `env:"MY_APP_PORT" env_default=8080"`
	Host     string `env:"HOST" env_default="localhost"`
	LogLevel string `env:"LOG_LEVEL" env_default="ERROR"`
}

var cfg ConfigDatabase
