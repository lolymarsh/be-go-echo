package config

type DatabaseConfigs struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Name                  string
	ConnMaxIdleTime       int
	ConnectionMaxLifeTime int
	MaxIdleConns          int
	MaxOpenConns          int
}
