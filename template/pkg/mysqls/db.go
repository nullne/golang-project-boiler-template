package mysqls

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/linxGnu/mssqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

const dsnTpl = "%s:%s@(%s:%s)/%s?charset=utf8&collation=utf8_general_ci&parseTime=true"

type Config struct {
	Database string
	User     string
	Password string
	Port     string
	Masters  []string
	Slaves   []string

	logger     zerolog.Logger
	enabledLog bool
}

func (c Config) dsn() string {
	return fmt.Sprintf(dsnTpl, c.User, c.Password, "%s", c.Port, c.Database)
}

func (c Config) masterDSNs() []string {
	return c.dsns(c.Masters)
}

func (c Config) slaveDSNs() []string {
	return c.dsns(c.Slaves)
}

func (c Config) dsns(ips []string) []string {
	ans := make([]string, len(ips))
	for i, ip := range ips {
		ans[i] = fmt.Sprintf(c.dsn(), ip)
	}
	return ans
}

func (c Config) mysqlInstantiate(driverName, dsn string) (*sql.DB, error) {
	if driverName != "mysql" {
		return nil, errors.New("only mysql")
	}
	driver := &mysql.MySQLDriver{}
	db := sqldblogger.OpenDriver(
		dsn,
		driver,
		zerologadapter.New(c.logger),
	)
	return db, nil
}

type Option func(*Config)

func WithLogger(logger zerolog.Logger) Option {
	return func(c *Config) {
		c.logger = logger
		c.enabledLog = true
	}
}

func New(c Config, opts ...Option) (*mssqlx.DBs, error) {
	for _, opt := range opts {
		opt(&c)
	}
	var sqlOptions []mssqlx.Option
	if c.enabledLog {
		sqlOptions = append(sqlOptions, mssqlx.WithDBInstantiate(c.mysqlInstantiate))
	}
	db, errs := mssqlx.ConnectMasterSlaves("mysql", c.masterDSNs(), c.slaveDSNs(), sqlOptions...)
	for _, e := range errs {
		if e != nil {
			return nil, e
		}
	}
	return db, nil
}
