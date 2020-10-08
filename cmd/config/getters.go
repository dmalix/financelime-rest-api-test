package config

type getter struct{}

var Get getter

func (g getter) Config() Config {
	return config
}
