package infra

import "github.com/AliceDiNunno/gobiquiti"

func LoadCloudKey() gobiquiti.Config {
	return gobiquiti.Config{
		Hostname: RequireEnvString("UBNT_HOSTNAME"),
		Username: RequireEnvString("UBNT_USERNAME"),
		Password: RequireEnvString("UBNT_PASSWORD"),
		Site:     RequireEnvString("UBNT_SITE"),
	}
}
