package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	// Port ポート
	// 0 が指定された場合は動的にポートを割り当てる。
	Port uint16 `env:"PORT" envDefault:"8080"`

	// DBHost データベースホスト
	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	// DBPort データベースポート
	DBPort uint16 `env:"DB_PORT" envDefault:"3306"`
	// DBName データベース名
	DBName string `env:"DB_NAME,required"`
	// DBUsername データベースユーザー名
	DBUsername string `env:"DB_USERNAME,required"`
	// DBPassword データベースパスワード
	DBPassword string `env:"DB_PASSWORD"`
}

func Get() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}
