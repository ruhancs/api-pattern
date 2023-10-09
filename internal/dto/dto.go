package dto

type ProductCreateDto struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UserCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtOutputDto struct {
	AccessToken string `json:"access_token"`
}
