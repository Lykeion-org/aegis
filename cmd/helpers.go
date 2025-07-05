package main

type Config struct {
	AuthenticationServerPort 	string 		`yaml:"authentication_service_port"`
	JwtSecret 					string 		`yaml:"jwt_secret"`
}