package account

type Account struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	LastConnexion string `json:"lastConnexion"`
}
