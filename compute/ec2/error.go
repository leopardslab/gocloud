package ec2

// represents error
type Error struct {
	StatusCode int

	Code string

	Message string

	RequestId string `xml:"RequestID"`
}
