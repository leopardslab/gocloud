package awsdns

import (
)


type Awsdns struct {
}


type ListResourcednsRecordSets struct {
	Name       string
	Type       string
	Identifier string
	MaxItems   int
}

type ChangeResourceRecordSetsRequest struct {
	Comment string   `xml:"ChangeBatch>Comment,omitempty"`
	Changes []Change `xml:"ChangeBatch>Changes>Change"`
}

type Change struct {
	Action string            `xml:"Action"`
	Record ResourceRecordSet `xml:"ResourceRecordSet"`
}

type AliasTarget struct {
	HostedZoneId         string
	DNSName              string
	EvaluateTargetHealth bool
}

type HostedZone struct {
	ID              string `xml:"Id"`
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"Config>Comment"`
	ResourceCount   int    `xml:"ResourceRecordSetCount"`
}

type ChangeInfo struct {
	ID          string `xml:"Id"`
	Status      string `xml:"Status"`
	SubmittedAt string `xml:"SubmittedAt"`
}

type DelegationSet struct {
	NameServers []string `xml:"NameServers>NameServer"`
}

type ResourceRecordSet struct {
	Name          string       `xml:"Name"`
	Type          string       `xml:"Type"`
	TTL           int          `xml:"TTL"`
	Records       []string     `xml:"ResourceRecords>ResourceRecord>Value,omitempty"`
	SetIdentifier string       `xml:"SetIdentifier,omitempty"`
	Weight        int          `xml:"Weight,omitempty"`
	HealthCheckId string       `xml:"HealthCheckId,omitempty"`
	Region        string       `xml:"Region,omitempty"`
	Failover      string       `xml:"Failover,omitempty"`
	AliasTarget   *AliasTarget `xml:"AliasTarget,omitempty"`
	RecordsXML    string       `xml:",innerxml"`
}

type CreateHostedZoneRequest struct {
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"HostedZoneConfig>Comment"`
}

type CreateHostedZoneResponse struct {
	HostedZone    HostedZone    `xml:"HostedZone"`
	ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

type DeleteHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

type ListHostedZonesResponse struct {
	HostedZones []HostedZone `xml:"HostedZones>HostedZone"`
	Marker      string       `xml:"Marker"`
	IsTruncated bool         `xml:"IsTruncated"`
	NextMarker  string       `xml:"NextMarker"`
	MaxItems    int          `xml:"MaxItems"`
}

type ChangeResourceRecordSetsResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

type ListResourceRecordSetsResponse struct {
	Records              []ResourceRecordSet `xml:"ResourceRecordSets>ResourceRecordSet"`
	IsTruncated          bool                `xml:"IsTruncated"`
	MaxItems             int                 `xml:"MaxItems"`
	NextRecordName       string              `xml:"NextRecordName"`
	NextRecordType       string              `xml:"NextRecordType"`
	NextRecordIdentifier string              `xml:"NextRecordIdentifier"`
}
