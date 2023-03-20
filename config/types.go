package config

type Config struct {
	Database   MySQL        `json:"mysql"`
	EndPoint   DansEndPoint `json:"dans_endpoint"`
	HTTPServer HTTP         `json:"http_server"`
	JWTKey     string       `json:"jwt_key"`
}

type MySQL struct {
	URL       string `json:"url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	TableName string `json:"table_name"`
	TimeoutMS int    `json:"timeout_ms"`
}

type DansEndPoint struct {
	URL       string `json:"url"`
	TimeoutMS int    `json:"timeout_ms"`
}

type HTTP struct {
	TimeoutMS int `json:"timeout_ms"`
}
