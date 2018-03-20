package ec2

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/cloudlibz/gocloud/auth"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// prepareRunParams base on vps or legacy params
func prepareRunParams(options RunInstances) map[string]string {
	if options.SubnetId != "" || len(options.NetworkInterfaces) > 0 {
		return makeParamsVPC("RunInstances")
	} else {
		return makeParams("RunInstances")
	}
}

func makeParams(action string) map[string]string {
	return makeParamsWithVersion(action, legacyAPIVersion)
}

func makeParamsVPC(action string) map[string]string {
	return makeParamsWithVersion(action, vpcAPIVersion)
}

//add version to params
func makeParamsWithVersion(action, version string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = version
	return params
}

// prepare Block Devices add into request param
func prepareBlockDevices(params map[string]string, blockDevs []BlockDeviceMapping) {
	for i, b := range blockDevs {
		n := strconv.Itoa(i + 1)
		prefix := "BlockDeviceMapping." + n
		if b.DeviceName != "" {
			params[prefix+".DeviceName"] = b.DeviceName
		}
		if b.VirtualName != "" {
			params[prefix+".VirtualName"] = b.VirtualName
		}
		if b.SnapshotId != "" {
			params[prefix+".Ebs.SnapshotId"] = b.SnapshotId
		}
		if b.VolumeType != "" {
			params[prefix+".Ebs.VolumeType"] = b.VolumeType
		}
		if b.VolumeSize > 0 {
			params[prefix+".Ebs.VolumeSize"] = strconv.FormatInt(b.VolumeSize, 10)
		}
		if b.IOPS > 0 {
			params[prefix+".Ebs.Iops"] = strconv.FormatInt(b.IOPS, 10)
		}
		if b.DeleteOnTermination {
			params[prefix+".Ebs.DeleteOnTermination"] = "true"
		}
	}
}

// prepareNetworkInterfaces add to request param

func prepareNetworkInterfaces(params map[string]string, nics []RunNetworkInterface) {
	for i, ni := range nics {
		n := strconv.Itoa(i)
		prefix := "NetworkInterface." + n
		if ni.Id != "" {
			params[prefix+".NetworkInterfaceId"] = ni.Id
		}
		params[prefix+".DeviceIndex"] = strconv.Itoa(ni.DeviceIndex)
		if ni.SubnetId != "" {
			params[prefix+".SubnetId"] = ni.SubnetId
		}
		if ni.Description != "" {
			params[prefix+".Description"] = ni.Description
		}
		for j, gid := range ni.SecurityGroupIds {
			k := strconv.Itoa(j + 1)
			params[prefix+".SecurityGroupId."+k] = gid
		}
		if ni.DeleteOnTermination {
			params[prefix+".DeleteOnTermination"] = "true"
		}
		if ni.SecondaryPrivateIPCount > 0 {
			val := strconv.Itoa(ni.SecondaryPrivateIPCount)
			params[prefix+".SecondaryPrivateIpAddressCount"] = val
		}
		for j, ip := range ni.PrivateIPs {
			k := strconv.Itoa(j)
			subprefix := prefix + ".PrivateIpAddresses." + k
			params[subprefix+".PrivateIpAddress"] = ip.Address
			params[subprefix+".Primary"] = strconv.FormatBool(ip.IsPrimary)
		}
	}
}

// use to create params to start stop and reboot, terminate instance
func addParamsList(params map[string]string, label string, ids []string) {
	for i, id := range ids {
		params[label+"."+strconv.Itoa(i+1)] = id
	}
}

func clientToken() (string, error) {

	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
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

//pass the param to query and add signature to it base on secret key and acces key

func (ec2 *EC2) PrepareSignatureV2query(params map[string]string, Region string, response map[string]interface{}) error {

	EC2Endpoint := "https://ec2." + Region + ".amazonaws.com"

	req, err := http.NewRequest("GET", EC2Endpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", timeNow().In(time.UTC).Format(time.RFC3339))

	req.URL.RawQuery = query.Encode()
	auth := map[string]string{"AccessKey": auth.Config.AWSAccessKeyID, "SecretKey": auth.Config.AWSSecretKey}
	awsauth.SignatureV2(req, auth)
	r, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	response["body"] = string(body)
	response["status"] = r.StatusCode

	return err
}
