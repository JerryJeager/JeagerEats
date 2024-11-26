package categories

type CategorySv interface {
}

type CategoryServ struct {
	repo CategoryStore
}

func NewCategoryService(repo CategoryStore) *CategoryServ {
	return &CategoryServ{repo: repo}
}
