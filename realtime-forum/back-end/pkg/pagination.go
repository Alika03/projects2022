package pkg

type Pagination struct {
	totalItems int
	limit      int
}

func NewPagination(totalItems, limit int) *Pagination {
	return &Pagination{
		totalItems: totalItems,
		limit:      limit,
	}
}

func (p *Pagination) PerPage() int {
	return p.limit
}

func (p *Pagination) PagesCount() int {
	if p.totalItems%p.limit != 0 {
		return (p.totalItems / p.limit) + 1
	}
	return p.totalItems / p.limit
}
