package main

import(
	"fmt"
	aegis_grpc "github.com/lykeion-org/aegis/internal/grpc"
	rest_api "github.com/lykeion-org/aegis/internal/api"
)

func main(){

	fmt.Println("Initializing aegis service")

	jwtSecret := []byte("your_jwt_secret_here")
	servicePort := ":30002"
	restPort := ":30800"
	server := aegis_grpc.NewAuthService(jwtSecret)

	fmt.Printf("Starting grpc server on port %s\n", servicePort)
	err := server.StartServer(servicePort)
	if err != nil {
		fmt.Printf("Failed to start service: %s\n", err)
	}

	fmt.Println("Listening for incoming requests")

	api := rest_api.NewApi(server.AuthHandler)
	api.InitializeApi(restPort)

	select {}

}