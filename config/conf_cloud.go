package config

type Cloud struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	ZoneName  string `yaml:"zoneName"`
	Host      string `yaml:"host"`
}
