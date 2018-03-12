package aliauth

import "time"

type Common struct {
	Format               string
	Version              string
	AccessKeyId          string
	Signature            string
	SignatureMethod      string
	Timestamp            time.Time
	SignatureVersion     string
	SignatureNonce       string
	ResourceOwnerAccount string
	Action               string
}