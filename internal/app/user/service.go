package user

import (
	"context"
	domain "galhub/internal/domain/user"
	"galhub/internal/pkg/password"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
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
) (*domain.User, error) {
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
	return user, nil
}
