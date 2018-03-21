package ec2

// Error represents error
type Error struct {
	StatusCode int
	Code string
	Message string
	RequestID string `xml:"RequestID"`
}
