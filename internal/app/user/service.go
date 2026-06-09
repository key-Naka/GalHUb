package user

import (
	"context"
	usercmd "galhub/internal/app/user/command"
	userquery "galhub/internal/app/user/query"
	domain "galhub/internal/domain/user"
	"galhub/internal/infrastructure/config"
	jwtpkg "galhub/internal/pkg/jwt"
	"galhub/internal/pkg/password"
	"time"
)

type Service struct {
	repo      domain.Repository
	jwtConfig config.JWTConfig
}

func NewService(
	repo domain.Repository,
	jwtConfig config.JWTConfig,
) *Service {
	return &Service{
		repo:      repo,
		jwtConfig: jwtConfig,
	}
}

func (s *Service) Register(
	ctx context.Context,
	cmd *usercmd.RegisterCommand,
) error {
	if cmd == nil {
		return ErrInvalidCommand
	}

	existUser, err := s.repo.GetByUsername(
		ctx,
		cmd.Username,
	)
	if err != nil {
		return err
	}
	if existUser != nil {
		return ErrUserExists
	}

	existEmail, err := s.repo.GetByEmail(
		ctx,
		cmd.Email,
	)
	if err != nil {
		return err
	}
	if existEmail != nil {
		return ErrEmailExists
	}

	hash, err := password.Hash(cmd.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: cmd.Username,
		Email:    cmd.Email,
		Password: string(hash),
		Role:     "user",
		Status:   1,
	}

	return s.repo.Create(ctx, user)
}

func (s *Service) Login(
	ctx context.Context,
	cmd *usercmd.LoginCommand,
) (*usercmd.LoginResult, error) {
	if cmd == nil {
		return nil, ErrInvalidCommand
	}

	user, err := s.repo.GetByEmail(
		ctx,
		cmd.Email,
	)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if !password.Verify(
		user.Password,
		cmd.Password,
	) {
		return nil, ErrInvalidPassword
	}

	token, err := jwtpkg.GenerateToken(
		user.ID,
		s.jwtConfig.Secret,
		s.jwtConfig.Expire,
	)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	user.LastLoginAt = &now
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &usercmd.LoginResult{
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Service) GetProfile(
	ctx context.Context,
	userID uint64,
) (*userquery.ProfileResult, error) {
	user, err := s.repo.GetByID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return &userquery.ProfileResult{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Status:      user.Status,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *Service) UpdateProfile(
	ctx context.Context,
	userID uint64,
	cmd *usercmd.UpdateProfileCommand,
) error {
	if cmd == nil {
		return ErrInvalidCommand
	}

	user, err := s.repo.GetByID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	user.Nickname = cmd.Nickname
	user.Avatar = cmd.Avatar

	return s.repo.Update(
		ctx,
		user,
	)
}

func (s *Service) ChangePassword(
	ctx context.Context,
	userID uint64,
	cmd *usercmd.ChangePasswordCommand,
) error {
	if cmd == nil {
		return ErrInvalidCommand
	}

	user, err := s.repo.GetByID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	if !password.Verify(
		user.Password,
		cmd.OldPassword,
	) {
		return ErrInvalidPassword
	}

	hash, err := password.Hash(
		cmd.NewPassword,
	)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	return s.repo.Update(
		ctx,
		user,
	)
}
