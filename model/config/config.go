package config

type ConnMysql struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

type ConnGin struct {
	HttpPort string
	//AppMode  string
}
