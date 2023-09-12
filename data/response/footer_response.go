package response

type FooterResponse struct {
	Id        int    `json:"id"`
	NoHP      string `json:"noHP"`
	Email     string `json:"email"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	LinkedIn  string `json:"linkedIn"`
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Youtube   string `json:"youtube"`
	UrlFoto   string `json:"urlFoto"`
}
