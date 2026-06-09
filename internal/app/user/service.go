package user

import (
	"context"
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

func (s *Service) Register(ctx context.Context, cmd *RegisterCommand) error {
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
	hash, err := password.Hash(
		cmd.Password,
	)
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
	cmd LoginCommand,
) (*LoginResult, error) {
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

	_ = s.repo.Update(
		ctx,
		user,
	)
	return &LoginResult{
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
func (s *Service) GetProfile(
	ctx context.Context,
	userID uint64,
) (*domain.User, error) {

	return s.repo.GetByID(
		ctx,
		userID,
	)
}
func (s *Service) UpdateProfile(
	ctx context.Context,
	userID uint64,
	cmd UpdateProfileCommand,
) error {

	user, err := s.repo.GetByID(
		ctx,
		userID,
	)

	if err != nil {
		return err
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
	cmd ChangePasswordCommand,
) error {

	user, err := s.repo.GetByID(
		ctx,
		userID,
	)

	if err != nil {
		return err
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

	user.Password = hash

	return s.repo.Update(
		ctx,
		user,
	)
}
