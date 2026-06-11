package company

import domain "galhub/internal/domain/company"

func ToDomain(
	m *CompanyModel,
) *domain.Company {

	return &domain.Company{
		ID: m.ID,

		Name: m.Name,

		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToModel(
	c *domain.Company,
) *CompanyModel {

	return &CompanyModel{
		ID: c.ID,

		Name: c.Name,

		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
