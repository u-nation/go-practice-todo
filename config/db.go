package config

import "fmt"

type DBConfig struct {
	Writer       *DBConnection `ignored:"true"`
	Reader       *DBConnection `ignored:"true"`
	MigrationDir string        `required:"true" split_words:"true"`
	MaxIdleConns int           `required:"true" split_words:"true"`
	MaxOpenConns int           `required:"true" split_words:"true"`
}

type DBConnection struct {
	DBName   string `required:"true" envconfig:"DB_NAME"`
	Host     string `required:"true" envconfig:"HOST"`
	Port     int    `required:"true" envconfig:"PORT"`
	User     string `required:"true" envconfig:"USER"`
	Password string `required:"true" envconfig:"PASSWORD"`
}

func (conns *DBConnection) GetConnectionUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conns.User,
		conns.Password,
		conns.Host,
		conns.Port,
		conns.DBName,
	)
}
