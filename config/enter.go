package config

// Config 映射每个配置
type Config struct {
	Mysql  *Mysql  `yaml:"mysql"`
	Logger *Logger `yaml:"logger"`
	Server *Server `yaml:"server"`
	Cloud  *Cloud  `yaml:"cloud"`
}
