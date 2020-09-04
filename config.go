package gormx

var datasource = new(Datasource)

type Datasource struct {
	Databases []*DatabaseConfig `yaml:"databases,flow"`
}

//DatabaseConfig db datasource
type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	DSN         string `yaml:"dsn"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxOpen     int    `yaml:"max_open"`
	MaxLifetime int    `yaml:"max_lifetime"`
	ShowSQL     bool   `yaml:"show_sql"`
}
