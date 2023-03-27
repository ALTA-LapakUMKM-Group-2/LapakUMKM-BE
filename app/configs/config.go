package configs

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
}

func InitConfig() *AppConfig {
	return InitEnv()
}
