package user

import (
	"context"
	domain "galhub/internal/domain/user"
	"galhub/internal/infrastructure/config"
	"galhub/internal/pkg/password"

	jwtpkg "galhub/internal/pkg/jwt"
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
