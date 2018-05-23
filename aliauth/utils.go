package aliauth

import (
	"reflect"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"time"
	srand "crypto/rand"
	"math/rand"
)

// PutStructIntoMap puts key and value of struct into map[string]interface{}
// if value is string and not empty -> put
// if value is bool                 -> put   NOTE: the default value of origin Ali API 's parameter must be false, if not ,do not use this function
// if value is int and not 0        -> put   NOTE: the optional values of origin Ali API 's parameter must dose not include 0, if not ,do not use this function
func PutStructIntoMap(i interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	e := reflect.ValueOf(i).Elem()
	typeOfOptions := e.Type()
	for i := 0; i < e.NumField(); i++ {
		switch e.Field(i).Type().String() {
		case "string":
			if e.Field(i).Interface() != "" {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		case "bool":
			params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
		case "int":
			if e.Field(i).Interface() != 0 {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		}
	}
	return params
}

//CreateSignature creates signature for string following Ali-cloud rules
func createSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// CreateRandomString create random string
func createRandomString() string {
	b := make([]byte, 32)
	l := len(dictionary)

	_, err := srand.Read(b)

	if err != nil {
		// Fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}
func getEndpointWithRegion(region string) string {
	switch region {
	case "cn-zhangjiakou":
		return region
	case "cn-huhehaote":
		return region
	case "ap-northeast-1":
		return region
	case "ap-southeast-2":
		return region
	case "ap-southeast-3":
		return region
	case "ap-southeast-5":
		return region
	case "ap-south-1":
		return region
	case "me-east-1":
		return region
	case "eu-central-1":
		return region
	default:
		return ""
	}
}
