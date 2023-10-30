package interfaces

type (
	Book struct {
		ID       int64
		Name     string
		Price    int64
		Desc     string
		CreateAt int64
		UpdateAt int64
	}

	CreateBookReq struct {
		Name  string
		Price int64
		Desc  string
	}

	UpdateBookReq struct {
		ID    int64
		Name  string
		Price int64
		Desc  string
	}

	GetBooksReq struct {
		Name     *string
		Price    *int64
		OrderBy  *string
		Page     int64
		PageSize int64
	}
)
