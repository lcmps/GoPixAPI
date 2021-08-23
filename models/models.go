package models

// PixOpts is a configuration struct.
type PixOpts struct {
	// Pix Key (CPF/CNPJ, Email, Cellphone or Random Key)
	Key string
	// Receiver name
	Name string
	// Receiver city
	City string
	// Transaction amount
	Amount float64
	// Transaction description
	Description string
	// Transaction ID
	TransactionID string
}

// QRCodeOpts is a configuration struct.
type QRCodeOpts struct {
	// QR Code content
	Content string
	// Default: 256
	Size int
}

type APIRequestQRCode struct {
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	City            string  `json:"city"`
	Description     string  `json:"description"`
	Transactionid   string  `json:"transactionId"`
	Pixkey          string  `json:"pixKey"`
	Foregroundcolor string  `json:"foregroundColor"`
	Backgroundcolor string  `json:"backgroundColor"`
}

type PathResp struct {
	Path string `json:"path"`
}
