package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Config   string `yaml:"config"` // 高级配置, 例如 charset
	DB       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	//MaxIdleConns int `yaml:"max-idle-Conns"`
	LogLevel     string `yaml:"log_level"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}
