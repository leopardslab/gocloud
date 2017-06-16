package ec2

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//Sign v2 method for Authenticating request

func SignV2(req *http.Request, auth Auth) (err error) {
	queryVals := req.URL.Query()
	queryVals.Set("AWSAccessKeyId", auth.AccessKey)
	queryVals.Set("SignatureVersion", "2")
	queryVals.Set("SignatureMethod", "HmacSHA256")

	queryStr, err := canonicalQueryString(queryVals)
	if err != nil {
		return err
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	payload := new(bytes.Buffer)
	if err := errorCollector(
		fprintfWrapper(payload, "%s\n", requestMethodVerb(req.Method)),
		fprintfWrapper(payload, "%s\n", req.Host),
		fprintfWrapper(payload, "%s\n", path),
		fprintfWrapper(payload, "%s", queryStr),
	); err != nil {
		return err
	}

	hash := hmac.New(sha256.New, []byte(auth.SecretKey))
	hash.Write(payload.Bytes())
	signature := make([]byte, base64.StdEncoding.EncodedLen(hash.Size()))
	base64.StdEncoding.Encode(signature, hash.Sum(nil))

	queryVals.Set("Signature", string(signature))
	req.URL.RawQuery = queryVals.Encode()

	return nil
}

func canonicalQueryString(queryVals url.Values) (string, error) {

	return strings.Replace(queryVals.Encode(), "+", "%20", -1), nil
}

func clientToken() (string, error) {

	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func fprintfWrapper(w io.Writer, format string, vals ...interface{}) func() error {
	return func() error {
		_, err := fmt.Fprintf(w, format, vals...)
		return err
	}
}

func errorCollector(writers ...func() error) error {
	for _, writer := range writers {
		if err := writer(); err != nil {
			return err
		}
	}

	return nil
}

var timeFormats = []string{
	time.RFC822,
	ISO8601BasicFormat,
	time.RFC1123,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
}

func requestTime(req *http.Request) (time.Time, error) {

	var date string
	if date = req.Header.Get("x-amz-date"); date == "" {
		if date = req.Header.Get("date"); date == "" {
			return time.Time{}, fmt.Errorf(`Could not retrieve a request date. Please provide one in either "x-amz-date", or "date".`)
		}
	}

	for _, format := range timeFormats {
		if parsedTime, err := time.Parse(format, date); err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf(
		"Could not parse the given date. Please utilize on of the following formats: %s",
		strings.Join(timeFormats, ","),
	)
}

func requestMethodVerb(rawMethod string) (verb string) {
	verbPlus := strings.SplitN(rawMethod, " ", 2)
	switch {
	case len(verbPlus) == 0:
		verb = "GET"
	default:
		verb = verbPlus[0]
	}
	return verb
}

func buildError(r *http.Response) error {
	errors := xmlErrors{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err Error
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

type xmlErrors struct {
	RequestId string  `xml:"RequestID"`
	Errors    []Error `xml:"Errors>Error"`
}

var timeNow = time.Now

func (err *Error) Error() string {
	if err.Code == "" {
		return err.Message
	}

	return fmt.Sprintf("%s (%s)", err.Message, err.Code)
}

const (
	ISO8601BasicFormat      = "20060102T150405Z"
	ISO8601BasicFormatShort = "20060102"
)
