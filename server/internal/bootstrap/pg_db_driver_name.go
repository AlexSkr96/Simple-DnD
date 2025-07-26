package bootstrap

import (
	"database/sql"
	"github.com/AlexSkr96/Simple-DnD/internal/configs"
	"github.com/AlexSkr96/Simple-DnD/pkg/database"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
)

type PGDBDriverName string

func NewPGDBDriverName(
	conf *configs.GORMConfig,
	logger logging.Logger,
) PGDBDriverName {
	const driverName = "pgxDBWithHooks"

	warnETThresholdsMS := map[string]int64{}
	queryETLogger := database.NewQueryETLogger(conf.DBTrackerName, logger, warnETThresholdsMS)

	hooks := sqlhooks.Compose(queryETLogger)

	sql.Register(driverName, sqlhooks.Wrap(&stdlib.Driver{}, hooks))
	sqlx.BindDriver(driverName, sqlx.DOLLAR)

	return driverName
}
