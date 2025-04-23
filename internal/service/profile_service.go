package service

import (
	"context"

	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/repository"
	pb "github.com/Skapar/simple-rest/proto"
)

type ProfileServiceImpl struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &ProfileServiceImpl{repo: repo}
}

func (s *ProfileServiceImpl) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	profile, err := s.repo.GetProfileByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileResponse{
		Profile: &pb.Profile{
			Id:        profile.ID,
			UserId:    profile.UserID,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Bio:       profile.Bio,
			AvatarUrl: profile.AvatarURL,
			CreatedAt: profile.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: profile.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *ProfileServiceImpl) CreateProfile(ctx context.Context, in *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	profile := &entities.Profile{
		UserID:    in.UserId,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Bio:       in.Bio,
		AvatarURL: in.AvatarUrl,
	}

	created, err := s.repo.CreateProfile(profile)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{
		Profile: &pb.Profile{
			Id:        created.ID,
			UserId:    created.UserID,
			FirstName: created.FirstName,
			LastName:  created.LastName,
			Bio:       created.Bio,
			AvatarUrl: created.AvatarURL,
			CreatedAt: created.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: created.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *ProfileServiceImpl) UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	profile := &entities.Profile{
		ID:        in.Id,
		UserID:    in.UserId,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Bio:       in.Bio,
		AvatarURL: in.AvatarUrl,
	}

	updated, err := s.repo.UpdateProfile(profile)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		Profile: &pb.Profile{
			Id:        updated.ID,
			UserId:    updated.UserID,
			FirstName: updated.FirstName,
			LastName:  updated.LastName,
			Bio:       updated.Bio,
			AvatarUrl: updated.AvatarURL,
			CreatedAt: updated.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: updated.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *ProfileServiceImpl) DeleteProfile(ctx context.Context, in *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	err := s.repo.DeleteProfile(in.Id)
	if err != nil {
		return &pb.DeleteProfileResponse{Success: false}, err
	}
	return &pb.DeleteProfileResponse{Success: true}, nil
}
