package config

// Config for the application
type Config struct {

	// server configuration
	Server struct {
		Port string `json:"port"`
	} `json:"server"`

	// configuration for redis
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`

	// encodec service config
	EncoDecConfig struct {
		
	}
	

	// if false application is launched in debug mode else in production mode
	Production bool `json:"production"`
}
