package ec2

// repersents error
type Error struct {
	StatusCode int

	Code string

	Message   string

	RequestId string `xml:"RequestID"`
}
