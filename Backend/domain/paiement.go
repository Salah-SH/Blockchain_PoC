package domain

type CreatePaiementResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Token  string  `json:"token"`
		Amount float64 `json:"amount"`
	} `json:"data"`
}
type CreatePaiement struct {
	Vender string `json:"vendor"`
	Amount string `json:"amount"`
	Note   string `json:"note"`
}
type Paiement struct {
	Paiementtoken string `json:"Paiementtoken" bson:"Paiementtoken"`
	Amount        string `json:"amount" bson:"amount"`
	Status        string `json:"status" bson:"status"`
	CagnotteName  string `json:"cagnottename" bson:"cagnottename"`
}
type PaymeeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		PaymentStatus bool    `json:"payment_status"`
		Token         string  `json:"token"`
		Amount        float64 `json:"amount"`
		TransactionID int     `json:"transaction_id"`
		BuyerID       int     `json:"buyer_id"`
	} `json:"data"`
}
