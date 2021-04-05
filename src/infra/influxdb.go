package infra

type InfluxDB struct {
	Hostname string
	Port     int
	Username string
	Password string
}

func LoadInflux() InfluxDB {
	return InfluxDB{
		Hostname: RequireEnvString("INFLUX_HOSTNAME"),
		Port:     RequireEnvInt("INFLUX_PORT"),
		Username: RequireEnvString("INFLUX_USER"),
		Password: RequireEnvString("INFLUX_PASSWORD"),
	}
}
