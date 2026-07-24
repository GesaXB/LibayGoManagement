package responsedto

type AuthorResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
