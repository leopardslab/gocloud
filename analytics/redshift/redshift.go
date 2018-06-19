package redshift

//CreateDatasets Create Datasets.
func (redshift *Redshift) CreateDatasets(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//DeleteDatasets delete Datasets.
func (redshift *Redshift) DeleteDatasets(request interface{}) (resp interface{}, err error) {

	return resp, err
}


type Describecluster struct{
	clusterIdentifier string
	marker string
	maxRecords int
	tagKeys []string
	tagValues []string
}

//GetDatasets get Datasets.
func (redshift *Redshift) GetDatasets(request interface{}) (resp interface{}, err error) {

	var region string

	var describecluster describecluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "ClusterIdentifier":
			clusterIdentifier, _ := value.(string)
			describecluster.clusterIdentifier = clusterIdentifier

		case "Marker":
			marker, _ := value.(string)
			describecluster.marker = marker

		case "MaxRecords":
			maxRecords, _ := value.(int)
			describecluster.maxRecords = maxRecords

		case "TagKeys":
			tagKeys, _ := value.([]string)
			describecluster.tagKeys = tagKeys

		case "TagValues":
			tagValues, _ := value.([]string)
			describecluster.TagValues = tagValues
		}
	}


	describeclusterspram := make(map[string]string)

	preparedefaultDescribeClusterspram(describeclusterspram)
	prepareDescribeClusterspram(describeclusterspram,describecluster)

 	return resp, err
}


func preparedefaultDescribeClusterspram(describeclusterspram map[string]string){

	describeclusterspram["Action"]="DescribeClusters"
  describeclusterspram["Version"]="2012-12-01"
}

func prepareDescribeClusterspram(describeclusterspram map[string]string,describecluster Describecluster){

	describeclusterspram["Action"]="DescribeClusters"
  describeclusterspram["Version"]="2012-12-01"

	if describecluster.clusterIdentifier != ""{
			describeclusterspram["ClusterIdentifier"] = describecluster.clusterIdentifier
	}


		if describecluster.marker != ""{
				describeclusterspram["Marker"] = describecluster.marker
	}

	if describecluster.maxRecords != 0 {
				describeclusterspram["MaxRecords"] = strconv.Itoa(describecluster.maxRecords)
	}


	if len(describecluster.tagKeys) != 0 {

		for i := 0 ;i< len(describecluster.tagKeys); i++{

			n := strconv.Itoa(i)

			prefix = "TagKeys.TagKey." + n

			describeclusterspram["prefix"] = describecluster.tagKeys[i]

		}
	}

	if len(describecluster.tagValues) != 0 {

		for i := 0 ; i< len(describecluster.tagValues); i++{

			n := strconv.Itoa(i)

			prefix = "TagValues.TagValue." + n

			describeclusterspram["prefix"] = describecluster.tagValues[i]
		}
	}
}


func (ec2 *EC2) PrepareSignaturequery(describeclusterspram map[string]string, region string, response map[string]interface{}) error {

	EC2Endpoint := "https://ec2." + regin + ".amazonaws.com"

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

	awsAuth.SignatureV2(req, auth)

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



//UpdateDatasets  Update Datasets.
func (redshift *Redshift) UpdateDatasets(request interface{}) (resp interface{}, err error) {

	return resp, err
}


//ListDatasets  list Datasets.
func (bigquery *Bigquery) ListDatasets(request interface{}) (resp interface{}, err error) {

}
