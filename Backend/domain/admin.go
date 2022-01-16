package domain

type AdminCreds struct {
	PaymeeAccount string `json:"paymee_account" bson:"paymee_account,omitempty"`
	AuthToken     string `json:"auth_token" bson:"auth_token,omitempty"`
}
