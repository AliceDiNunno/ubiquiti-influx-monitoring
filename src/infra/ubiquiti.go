package infra

type UbiquitiServer struct {
	Hostname string
	Username string
	Password string
	Site     string
}

func LoadCloudKey() UbiquitiServer {
	return UbiquitiServer{
		Hostname: RequireEnvString("UBNT_HOSTNAME"),
		Username: RequireEnvString("UBNT_USERNAME"),
		Password: RequireEnvString("UBNT_PASSWORD"),
		Site:     RequireEnvString("UBNT_SITE"),
	}
}
