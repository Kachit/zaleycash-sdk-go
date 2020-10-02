package zaleycash_sdk

import "time"

type MyTarget struct {
	*ResourceAbstract
}

func (m *MyTarget) GetToken(accountId string) (*Response, error) {
	body := make(map[string]interface{})
	body["account_id"] = accountId
	return m.Post("api/v2/my_target/token", body, nil)
}

type MyTargetToken struct {
	AccessToken string  `json:"accessToken"`
	ExpiresIn   float64 `json:"expiresIn"`
}

func (t *MyTargetToken) IsValid() bool {
	return t.IsNotExpired() && t.AccessToken != ""
}

func (t *MyTargetToken) IsNotExpired() bool {
	ts := time.Now().Unix()
	return int64(t.ExpiresIn) > ts
}
