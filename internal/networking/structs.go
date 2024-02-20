package networking

// Ifconfig is a struct that represents the response from https://ifconfig.co/json
type Ifconfig struct {
	IP string `json:"ip"`
}
