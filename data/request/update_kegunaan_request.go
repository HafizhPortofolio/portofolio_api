package request

type UpdateKegunaansRequest struct {
	Id   int    `validate:"required"`
	Nama string `validate:"required,min=1,max=200" json:"nama"`
}
