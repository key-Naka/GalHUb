package user

import (
	"context"
	domain "galhub/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, cmd *RegisterCommand) error {
	existUser, _ := s.repo.GetByUsername(
		ctx,
		cmd.Username,
	)
	if existUser != nil {
		return ErrUserExists
	}
	existEmail, _ := s.repo.GetByEmail(
		ctx,
		cmd.Email,
	)
	if existEmail != nil {
		return ErrEmailExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
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
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(cmd.Password),
	)

	if err != nil {
		return nil, ErrInvalidPassword
	}
	return user, nil
}
