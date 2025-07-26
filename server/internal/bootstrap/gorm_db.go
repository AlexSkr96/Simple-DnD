package bootstrap

import (
	"github.com/AlexSkr96/Simple-DnD/internal/configs"
	gormint "github.com/AlexSkr96/Simple-DnD/pkg/gorm"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewGORMDB(
	logger logging.Logger,
	conf *configs.GORMConfig,
	driverName PGDBDriverName,
) (db *gorm.DB, cleanup func(), err error) { // nolint: dupl
	sqlDB, err := sqlx.Open(string(driverName), conf.DBConn)

	cleanup = func() {
		if sqlDB == nil {
			return
		}

		err := sqlDB.Close()
		if err != nil {
			logger.Error(errors.WithStack(err))
		}

		logger.Info("closing internal postgres connections")
	}

	if err != nil {
		return nil, cleanup, errors.WithStack(err)
	}

	// Validate and set connection pool parameters with defaults if unset
	maxOpenConns := conf.DBMaxOpenConns
	if maxOpenConns <= 0 {
		maxOpenConns = 10 // Default value
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)
	maxIdleConns := conf.DBMaxIdleConns
	if maxIdleConns <= 0 {
		maxIdleConns = 5 // Default value
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	connMaxIdleTime := conf.DBConnMaxIdleTimeMin
	if connMaxIdleTime <= 0 {
		connMaxIdleTime = 5 // Default value in minutes
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Minute)
	connMaxLifetime := conf.DBConnMaxLifetimeMin
	if connMaxLifetime <= 0 {
		connMaxLifetime = 30 // Default value in minutes
	}
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		return nil, cleanup, errors.WithStack(err)
	}

	gormConf := &gorm.Config{
		Logger: gormint.NewLogger(logger),
	}
	dialector := postgres.New(postgres.Config{Conn: sqlDB})

	db, err = gorm.Open(dialector, gormConf)
	if err != nil {
		return nil, cleanup, errors.WithStack(err)
	}

	return db, cleanup, nil
}
