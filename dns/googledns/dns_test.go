package googledns

import "testing"
import "fmt"

func TestCreatedns(t *testing.T) {
	var googledns Googledns

	createdns := map[string]interface{}{
		"Project":     "sheltermap-1493101612061",
		"Kind":        "dns#managedZone",
		"Description": "dns",
		"DnsName":     "rootmonk.me.",
		"Name":        "gocloud",
	}
	_, err := googledns.Createdns(createdns)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestListdns(t *testing.T) {
	var googledns Googledns

	listdns := map[string]string{
		"Project": "sheltermap-1493101612061",
	}

	_, err := googledns.Listdns(listdns)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestListResourcednsRecordSets(t *testing.T) {
	var googledns Googledns

	listResourcednsRecordSets := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.ListResourcednsRecordSets(listResourcednsRecordSets)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestDeletedns(t *testing.T) {
	var googledns Googledns

	deletedns := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.Deletedns(deletedns)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
