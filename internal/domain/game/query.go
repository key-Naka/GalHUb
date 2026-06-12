package game

type ListQuery struct {
	Offset int
	Limit  int

	Keyword string

	TagID *uint64

	CompanyID *uint64

	Status *int8
}
