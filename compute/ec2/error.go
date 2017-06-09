package ec2

type Error struct {
	StatusCode int

	Code string

	Message   string

	RequestId string `xml:"RequestID"`
}

