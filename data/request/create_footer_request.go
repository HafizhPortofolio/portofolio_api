package request

type CreateFootersRequest struct {
	NoHP      string `validate:"required,min=1,max=200" json:"noHP"`
	Email     string `validate:"required,min=1,max=200" json:"email"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	LinkedIn  string `json:"linkedIn"`
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Youtube   string `json:"youtube"`
	UrlFoto   string `json:"urlFoto"`
}
