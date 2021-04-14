package cfg

import "os"

func GetGrpcPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		panic("grpc port not specified")
	}

	return ":" + p
}

func GetConnString() string {
	cs := os.Getenv("CONNECTION_STRING")
	if cs == "" {
		panic("connection string not specified")
	}

	return cs
}
