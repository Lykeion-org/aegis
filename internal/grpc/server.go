package grpc

import (
	"context"
	"net"

	auth "github.com/lykeion-org/aegis/internal/auth"
	pb "github.com/lykeion-org/aegis/internal/grpc/generated"
	"google.golang.org/grpc"
)

type AuthService struct {
	pb.UnimplementedAegisServiceServer
	AuthHandler auth.AuthHandler
	activeServer *grpc.Server
}

func NewAuthService(jwtSecret []byte) *AuthService{
	return &AuthService{
		AuthHandler: auth.NewAuthHandler(jwtSecret),
	}
}

func (s *AuthService)StartServer(target string) error {
	grpcServer := grpc.NewServer()
	pb.RegisterAegisServiceServer(grpcServer,s)

	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	go func(){
		err := grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
	
	s.activeServer = grpcServer
	return nil
}

func (s *AuthService) StopServer() error {
	s.activeServer.GracefulStop()
	return nil
}

func (s *AuthService) GenerateTokens(ctx context.Context, req *pb.GenerateTokensRequest) (*pb.GenerateTokensResponse, error){
	tokens, err := s.AuthHandler.CreateToken(ctx, req.UserUid, req.UserRole)
	if err != nil {
		return nil, err
	}

	return &pb.GenerateTokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error){
	token, err := s.AuthHandler.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshTokenResponse{
		AccessToken: token,
	}, nil

}
func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error){
	claims, err := s.AuthHandler.ValidateAccessToken(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateTokenResponse{
		UserUid: claims.UserUid,
		UserRole: claims.Role,
	}, nil
}

