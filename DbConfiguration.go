package main

// DbConfiguration presents database configuration
type DbConfiguration struct {
	Host     string `json:"Host" envconfig:"DB_HOST"`
	User     string `json:"User" envconfig:"DB_USER"`
	Password string `json:"Password" envconfig:"DB_PASSWORD"`
	Schema   string `json:"Schema" envconfig:"DB_SCHEMA"`
}
