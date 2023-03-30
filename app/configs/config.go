package configs

import "lapakUmkm/utils/helpers"

type AppConfig struct {
	DBUSERNAME          string
	DBPASS              string
	DBHOST              string
	DBPORT              string
	DBNAME              string
	JWTKEY              string
	CLIENTIDGOOGLE      string
	CLIENTSECRETGOOGLE  string
	SERVER_KEY_MIDTRANS string
	GMAILPASSWORD       string
}

func InitConfig() *AppConfig {
	return InitEnv()
}

func LoadConfig(cfg *AppConfig) {
	helpers.OauthConfig.ClientID = cfg.CLIENTIDGOOGLE
	helpers.OauthConfig.ClientSecret = cfg.CLIENTSECRETGOOGLE
	helpers.ServerKey = cfg.SERVER_KEY_MIDTRANS
	helpers.GMAILPASS = cfg.GMAILPASSWORD
}
