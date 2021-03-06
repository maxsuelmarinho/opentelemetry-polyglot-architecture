package repository

import (
	"fmt"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	dbDriverNameDefault = "postgres"
)

// Database defines wrapper operations
type Database interface {
	Initialize()
	GetConnection() *sqlx.DB
}

type database struct {
	db     *sqlx.DB
	logger *log.Logger
}

// NewDatabase creates a new database
func NewDatabase(l *log.Logger) Database {
	return &database{
		logger: l,
	}
}

// Initialize initializes database
func (d *database) Initialize() {
	d.setDefaultPoolConfiguration()

	if viper.GetBool("DB_MIGRATION_ENABLED") {
		d.runMigration()
	}

	driverName := viper.GetString("DB_DRIVER_NAME")

	connectionStringParameters, err := d.getConnectionStringParameters(driverName, viper.GetString("DB_USER_DML"), viper.GetString("DB_PASS_DML"))
	if err != nil {
		d.logger.WithError(err).Fatal("Could not create connection string")
	}

	dbDML, err := sqlx.Connect(driverName, connectionStringParameters)
	if err != nil {
		d.logger.WithError(err).Fatal("Could not connect using DML user")
	}

	dbDML.SetMaxOpenConns(viper.GetInt("DB_POOL_MAX_OPEN_CONNS"))
	dbDML.SetMaxIdleConns(viper.GetInt("DB_POOL_MAX_IDLE_CONNS"))
	dbDML.SetConnMaxLifetime(viper.GetDuration("DB_POOL_CONN_MAX_LIFETIME"))

	d.db = dbDML
}

func (d *database) GetConnection() *sqlx.DB {
	return d.db
}

func (d *database) getConnectionStringParameters(driverName, user, password string) (string, error) {
	switch driverName {
	case "postgres":
		return d.getConnectionStringParametersToPostgres(user, password), nil
	default:
		return "", fmt.Errorf("driver name %s is not supported", driverName)
	}
}

func (d *database) getConnectionStringParametersToPostgres(user, password string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d application_name=%s",
		user,
		password,
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SERVER"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_APPLICATION_NAME"))
}

func (d *database) setDefaultPoolConfiguration() {
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_DRIVER_NAME", dbDriverNameDefault)
	viper.SetDefault("DB_MIGRATION_ENABLED", true)
	viper.SetDefault("DB_POOL_MAX_OPEN_CONNS", 0)
	viper.SetDefault("DB_POOL_MAX_IDLE_CONNS", 2)
	viper.SetDefault("DB_POOL_CONN_MAX_LIFETIME", 5*time.Minute)

	applicationName := fmt.Sprintf("%s-%s", viper.GetString("APP_NAME"), uuid.Must(uuid.NewRandom()).String())
	viper.SetDefault("DB_APPLICATION_NAME", applicationName)
}

func (d *database) runMigration() {
	driverName := viper.GetString("DB_DRIVER_NAME")
	connectionStringParameters, err := d.getConnectionStringParameters(driverName, viper.GetString("DB_USER_DDL"), viper.GetString("DB_PASS_DDL"))
	if err != nil {
		d.logger.WithError(err).Fatal("could not create connection string")
	}

	dbDDL, err := sqlx.Connect(driverName, connectionStringParameters)
	if err != nil {
		d.logger.WithError(err).Fatal("could not connect using DDL user")
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox(viper.GetString("DB_MIGRATION_PATH")),
	}

	migrate.SetTable(viper.GetString("DB_MIGRATION_TABLE"))

	if "create-drop" == viper.GetString("DB_DDL_AUTO") {
		_, err = migrate.Exec(dbDDL.DB, driverName, migrations, migrate.Down)
		if err != nil {
			d.logger.WithError(err).Fatal("could not start sql migration down")
		}
	}

	_, err = migrate.Exec(dbDDL.DB, driverName, migrations, migrate.Up)
	if err != nil {
		d.logger.WithError(err).Fatal("could not start sql migration up")
	}
	dbDDL.Close()
}
