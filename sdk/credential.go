package sdk

// Credential is used to sign the request
type Credential struct {
	MerchantId int64
	SecretKey  string
}

func NewCredentials(mid int64, secretKey string) *Credential {
	return &Credential{mid, secretKey}
}
