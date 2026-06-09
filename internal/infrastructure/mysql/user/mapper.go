package user

import domain "galhub/internal/domain/user"

func ToDomain(m *UserModel) *domain.User {

	return &domain.User{
		ID:          m.ID,
		Username:    m.Username,
		Email:       m.Email,
		Password:    m.Password,
		Nickname:    m.Nickname,
		Avatar:      m.Avatar,
		Role:        m.Role,
		Status:      m.Status,
		LastLoginAt: m.LastLoginAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ToModel(u *domain.User) *UserModel {

	return &UserModel{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		Nickname:    u.Nickname,
		Avatar:      u.Avatar,
		Role:        u.Role,
		Status:      u.Status,
		LastLoginAt: u.LastLoginAt,
	}
}
