package configs

type GORMConfig struct {
	DBConn               string `envconfig:"DB_CONN"                    required:"true"`
	DBTrackerName        string `default:"dnd-db"                       envconfig:"DB_TRACKER_NAME"        required:"true"`
	DBMaxOpenConns       int    `envconfig:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns       int    `envconfig:"DB_MAX_IDLE_CONNS"`
	DBConnMaxIdleTimeMin int    `envconfig:"DB_CONN_MAX_IDLE_TIME_MIN"`
	DBConnMaxLifetimeMin int    `envconfig:"DB_CONN_MAX_LIFETIME_MIN"`
	DBMetricsRefreshMin  int    `default:"5"                            envconfig:"DB_METRICS_REFRESH_MIN"`
}
