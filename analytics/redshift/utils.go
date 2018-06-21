package redshift

import (
	"fmt"
	awsAuth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
)



func preparedefaultDescribeClusterspram(describeclusterspram map[string]string) {

	describeclusterspram["Action"] = "DescribeClusters"
	describeclusterspram["Version"] = "2012-12-01"
}

func prepareDescribeClusterspram(describeclusterspram map[string]string, describecluster Describecluster) {

	if describecluster.clusterIdentifier != "" {
		describeclusterspram["ClusterIdentifier"] = describecluster.clusterIdentifier
	}

	if describecluster.marker != "" {
		describeclusterspram["Marker"] = describecluster.marker
	}

	if describecluster.maxRecords != 0 {
		describeclusterspram["MaxRecords"] = strconv.Itoa(describecluster.maxRecords)
	}

	if len(describecluster.tagKeys) != 0 {

		for i := 0; i < len(describecluster.tagKeys); i++ {

			n := strconv.Itoa(i)

			prefix := "TagKeys.TagKey." + n

			describeclusterspram[prefix] = describecluster.tagKeys[i]

		}
	}

	if len(describecluster.tagValues) != 0 {

		for i := 0; i < len(describecluster.tagValues); i++ {

			n := strconv.Itoa(i)

			prefix := "TagValues.TagValue." + n

			describeclusterspram[prefix] = describecluster.tagValues[i]
		}
	}
}

func PrepareSignaturequery(describeclusterspram map[string]string, region string, response map[string]interface{}) error {

	service := "redshift"

	endpoint := "https://" + service + "." + region + ".amazonaws.com"

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range describeclusterspram {
		query.Add(varName, varVal)
	}

	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.RawQuery)

	awsAuth.Getrequestsign4(req, region, service)

	r, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	response["body"] = string(body)
	response["status"] = r.StatusCode

	return err
}
