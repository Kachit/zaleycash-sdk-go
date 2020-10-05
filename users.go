package zaleycash_sdk

type UsersResource struct {
	*ResourceAbstract
}

/**
 * @unmarshal User
 */
func (u *UsersResource) GetByEmail(email string) (*Response, error) {
	query := make(map[string]interface{})
	if email != "" {
		query["email"] = email
	}
	return u.Get("api/v2/user", query)
}

/**
 * @unmarshal Balance
 */
func (u *UsersResource) GetBalance() (*Response, error) {
	return u.Get("api/v2/user/balance", nil)
}

type Balance struct {
	User  *User          `json:"user"`
	Items []*BalanceItem `json:"balances"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type BalanceItem struct {
	Currency        int     `json:"currency"`
	Balance         float64 `json:"balance"`
	CreditLimit     float64 `json:"creditLimit"`
	OriginalBalance float64 `json:"originalBalance"`
}
