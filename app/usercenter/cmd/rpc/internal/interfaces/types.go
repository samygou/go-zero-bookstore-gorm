package interfaces

type (
	Account struct {
		ID         int64
		CreateTime int64
		UpdateTime int64
		Mobile     string
		Username   string
		Password   string
		Sex        int64
		Avatar     string
		Remark     string
	}

	CreateAccountReq struct {
		Mobile   string
		Username string
		Password string
		Sex      int64
		Avatar   string
		Remark   string
	}

	UpdateAccountReq struct {
		ID       int64
		Mobile   string
		Username string
		Password string
		Sex      int64
		Avatar   string
		Remark   string
	}
)
