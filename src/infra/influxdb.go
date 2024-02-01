package infra

type InfluxDB struct {
	Hostname string
	Port     int
	IsSecure bool

	Token string

	Organization string
	Bucket       string
}

func LoadInflux() InfluxDB {
	db := InfluxDB{
		Hostname: RequireEnvString("INFLUX_HOSTNAME"),
		Port:     RequireEnvInt("INFLUX_PORT"),
		IsSecure: GetEnvStringWithDefault("INFLUX_SECURE", "false") == "true",

		Token: RequireEnvString("INFLUX_TOKEN"),

		Organization: RequireEnvString("INFLUX_ORGANIZATION"),
		Bucket:       RequireEnvString("INFLUX_BUCKET"),
	}

	return db
}
