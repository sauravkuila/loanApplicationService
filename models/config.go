package models

type DbConfig struct {
	Conn struct {
		Host                   string `yaml:"host"`
		Port                   string `yaml:"port"`
		Username               string `yaml:"username"`
		Password               string `yaml:"password"`
		AuthenticationDatabase string `yaml:"authenticationDatabase"`
	} `yaml:"conn"`
	Tmp string `yaml:"tmp"`
}

type DbConf struct {
	Driver          string `yaml:"driverName"`
	Url             string `yaml:"url"`
	MaxOpenConn     int    `yaml:"maxOpenConnections"`
	MaxIdleConn     int    `yaml:"maxIdleConnections"`
	ConnMaxTime     int    `yaml:"connectionMaxLifetimeInSeconds"`
	ConnMaxIdleTime int    `yaml:"connectionMaxIdleTimeInSeconds"`
}
