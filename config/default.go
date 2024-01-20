package config

var defaultConfig = Config{
	Logger: "console",
	Server: Server{
		Host: "localhost",
		Port: 8000,
	},
}
