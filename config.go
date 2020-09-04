package gormx

var datasource = new(Datasource)

type Datasource struct {
	Databases []*DatabaseConfig `yaml:"databases,flow"`
}

//DatabaseConfig db datasource
type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	DSN         string `yaml:"dsn"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxOpen     int    `yaml:"maxOpen"`
	MaxLifetime int    `yaml:"maxLifetime"`
	ShowSQL     bool   `yaml:"showSQL"`
}

//Databases Multi Database datasource
func Databases() []*DatabaseConfig {
	return datasource.Databases
}

//Database first database datasource
func Database() *DatabaseConfig {
	if len(datasource.Databases) > 0 {
		return datasource.Databases[0]
	}
	return nil
}
