package server

import (
	"context"

	"github.com/Skapar/simple-rest/internal/service"
	profile "github.com/Skapar/simple-rest/proto"
)

type Server struct {
	profile.UnsafeProfileServiceServer

	ProfileService service.ProfileService
}

func New(
	profileService service.ProfileService,
) *Server {
	return &Server{
		ProfileService: profileService,
	}
}

func (s *Server) GetProfile(ctx context.Context, in *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {
	return s.ProfileService.GetProfile(ctx, in)
}

func (s *Server) CreateProfile(ctx context.Context, in *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	return s.ProfileService.CreateProfile(ctx, in)
}

func (s *Server) UpdateProfile(ctx context.Context, in *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	return s.ProfileService.UpdateProfile(ctx, in)
}

func (s *Server) DeleteProfile(ctx context.Context, in *profile.DeleteProfileRequest) (*profile.DeleteProfileResponse, error) {
	return s.ProfileService.DeleteProfile(ctx, in)
}
