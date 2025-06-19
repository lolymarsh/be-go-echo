package configs

type ServerConfigs struct {
	PortAPI      string
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
	TimeZone     string
	TimeFormat   string
	Format       string
}
