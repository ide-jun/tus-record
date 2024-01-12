package database

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

func Open(host string, port uint16, dbname, username, password string) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
	cfg.DBName = dbname
	cfg.User = username
	cfg.Passwd = password
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
