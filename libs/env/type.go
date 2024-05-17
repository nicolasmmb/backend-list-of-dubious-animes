package env

type env struct {
	// DB
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	// SERVER
	SERVER_HOST string
	SERVER_PORT string
	// OTEL
	SERVICE_NAME           string
	OTLP_ENDPOINT          string
	OTLP_SCHEMA_URL        string
	SERVICE_VERSION        string
	SERVICE_NAMESPACE      string
	DEPLOYMENT_ENVIRONMENT string
	// JWT
	JWT_SECRET             string
	JWT_EXPIRATION         int
	JWT_REFRESH_EXPIRATION int
}

func (e env) GetDBConnString() string {
	return "host=" + e.DB_HOST + " port=" + e.DB_PORT + " user=" + e.DB_USER + " dbname=" + e.DB_NAME + " password=" + e.DB_PASS + " sslmode=disable"
}

func (e env) GetServerAddr() string {
	return e.SERVER_HOST + ":" + e.SERVER_PORT
}
