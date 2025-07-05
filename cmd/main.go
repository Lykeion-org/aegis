package main

import(
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/config"
	utils "github.com/Lykeion-org/go-shared/pkg/helpers"
	grpc "github.com/Lykeion-org/aegis/internal/grpc"
)


func main(){
	cfg, err := config.LoadConfigFile[Config]("../config.yaml")
	if err != nil {
		panic(err)
	}

	utils.InitializeStandardLogger("debug")

	slog.Info("Initializing application")

	server := grpc.NewAuthService([]byte(cfg.JwtSecret))
	slog.Info("Starting grpc server", "port", cfg.AuthenticationServerPort)
	err = server.StartServer(cfg.AuthenticationServerPort)
	if err != nil {
		slog.Error("Failed to start grpc server", "error", err)
	}

	slog.Info("Application started")

	select{}

}