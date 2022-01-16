package domain

type User struct {
	ID        string     `json:"id,omitempty" bson:"_id,omitempty"`
	AccAddr   string     `json:"address"`
	Identity  *Identity  `json:"identity" bson:"identity,omitempty"`
	Paiements []Paiement `json:"paiements" bson:"paiements,omitempty"`
}
type Identity struct {
	FirstName string `json:"firstname" bson:"firstname,omitempty"`
	LastName  string `json:"lastname" bson:"lastname,omitempty"`
	BankCard  string `json:"card" bson:"card,omitempty"`
}
