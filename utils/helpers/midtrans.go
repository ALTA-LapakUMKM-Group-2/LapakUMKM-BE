package helpers

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/Findryankp/snapMidtransGo"
)

var ServerKey = ""

func PostMidtrans(data map[string]any) (string, error) {
	var postData = snapMidtransGo.DataPostMidtrans{
		OrderId:   data["order_id"].(string),
		Nominal:   data["nominal"].(int),
		FirstName: data["firstname"].(string),
		LastName:  data["lastname"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
		ServerKey: ServerKey,
	}

	test, err := snapMidtransGo.SanboxRequestSnapMidtrans(postData)
	return test, err
}

type ResponseFromCallbackMidtrans struct {
	TransactionId     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	GrossAmount       string `json:"gross_amount"`
}

func ValidateSignatureKey(response ResponseFromCallbackMidtrans, orderId, statusCode string) bool {
	str := orderId + statusCode + ServerKey
	hash := sha512.Sum512([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return string(hashStr) == response.SignatureKey
}
