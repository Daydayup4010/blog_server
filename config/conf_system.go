package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *System) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
