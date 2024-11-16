package module

type TokenC struct {
	Token string `json:"token" bson:"token"`
}

func (t *TokenC) C() string {
	return "token"
}
