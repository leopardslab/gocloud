package awsdns

import (
	"fmt"
)

type Awsdns struct {
}


type Vpc struct {
	VPCID string `xml:"VPCId,omitempty"`
	VPCRegion string `xml:"VPCRegion,omitempty"`
}


type hostedZoneConfig struct {
			Comment string `xml:"Comment,omitempty"`
			PrivateZone string `xml:"PrivateZone,omitempty"`
		}



type CreateHostedZoneRequest struct {
		HostedZoneConfig  *hostedZoneConfig`xml:"HostedZoneConfig,omitempty"`
		Xmlns string `xml:"xmlns,omitempty"`
		CallerReference string `xml:"CallerReference,omitempty"`
		DelegationSetID string `xml:"DelegationSetId,omitempty"`
		Name string `xml:"Name,omitempty"`
		VPC *Vpc `xml:"VPC,omitempty"`
	}



func (awsdns *Awsdns) Listdns(request interface{}) (resp interface{}, err error) {

	return
}


func (awsdns *Awsdns) Createdns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var option CreateHostedZoneRequest

	for key, value := range param {
		switch key {

		case "name":
			nameV, _ := value.(string)
			option.Name = nameV

		case "delegationSetID":
			DelegationSetIDV, _ := value.(string)
			option.DelegationSetID = DelegationSetIDV

		case "xmlns":
			XmlnsV, _ := value.(string)
			fmt.Println(XmlnsV)
			option.Xmlns = "https://route53.amazonaws.com/doc/2013-04-01/"

		case "callerReference":
			CallerReferenceV, _ := value.(string)
			option.CallerReference = CallerReferenceV

		case "hostedZoneConfig":
			hostedZoneConfigV, _ := value.(map[string]string)
      option.HostedZoneConfig = &hostedZoneConfig{Comment: hostedZoneConfigV["Comment"], PrivateZone:hostedZoneConfigV["PrivateZone"]}

		case "VPC":
			VPCV, _ := value.(map[string]string)
			option.VPC = &Vpc{VPCID : VPCV["VPCID"], VPCRegion : VPCV["VPCRegion"]}

		}
	}

	fmt.Println(option.Name)
	return

}
