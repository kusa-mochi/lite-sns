package server_configs

type AppConfig struct {
	ApiPrefix string `json:"api_prefix"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
}

type DbConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type FrontendConfig struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type SmtpConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ServerConfigs struct {
	App      AppConfig      `json:"app"`
	Db       DbConfig       `json:"db"`
	Frontend FrontendConfig `json:"frontend"`
	Smtp     SmtpConfig     `json:"smtp"`
}
