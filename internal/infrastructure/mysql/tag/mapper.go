package tag

import domain "galhub/internal/domain/tag"

// ToDomain 将模型转换为领域实体
func ToDomain(
	m *TagModel,
) *domain.Tag {

	return &domain.Tag{
		ID: m.ID,

		Name: m.Name,

		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// ToModel 将领域实体转换为模型
func ToModel(
	t *domain.Tag,
) *TagModel {

	return &TagModel{
		ID: t.ID,

		Name: t.Name,

		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
