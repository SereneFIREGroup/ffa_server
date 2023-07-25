package db

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	logUtils "github.com/serenefiregroup/ffa_server/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func buildDBM() (*gorm.DB, error) {
	connstring := config.String("db_spec", "")
	dbMaxIdle := config.Int("db_max_idle", 10)
	dbMaxOpen := config.Int("db_max_open", 1000)
	connMaxLifetime := config.Int("db_conn_max_life_time", 1800)
	dbShowLog := config.Bool("db_show_log", false)

	dbConfig := new(gorm.Config)
	if dbShowLog {
		dbConfig.Logger = newDbLogger()
	}

	db, err := gorm.Open(mysql.Open(connstring), dbConfig)
	if err != nil {
		return nil, errors.Sql(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Sql(err)
	}
	sqlDB.SetMaxIdleConns(dbMaxIdle)
	sqlDB.SetMaxOpenConns(dbMaxOpen)
	if connMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}
	return db, nil
}

func newDbLogger() logger.Interface {
	logFile, _ := logUtils.GetLogFile()
	writer := io.MultiWriter(os.Stdout, logFile)
	return logger.New(
		log.New(writer, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
}

func InitDB() (err error) {
	if DB, err = buildDBM(); err != nil {
		return
	}
	return
}

func Transact(txFunc func(tx *gorm.DB) error) (err error) {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return err
	}
	err = txFunc(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
