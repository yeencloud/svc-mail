package config

type TemplaterConfig struct {
	Engine string `config:"TEMPLATER_ENGINE" default:"HERMES"`
}
