package request

type CreateKegunaansRequest struct {
	Nama string `validate:"required,min=1,max=200" json:"nama"`
}
