package awsmachinelearning

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
)


func(awsmachinelearning *Awsmachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error) {

/*




	param := request.(map[string]interface{})

	var Region string
  var createMLModel CreateMLModel

	for key, value := range param {
		switch key {
			case "MLModelId":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV

			case "MLModelName":
				MLModelNameV, _ := value.(string)
				createMLModel.MLModelName = MLModelNameV

			case "MLModelType":
				MLModelTypeV, _ := value.(string)
				createMLModel.MLModelType = MLModelTypeV

			case "Recipe":
				RecipeV, _ := value.(string)
				createMLModel.Recipe = RecipeV


			case "RecipeURI":
				RecipeURIV, _ := value.(string)
				createMLModel.RecipeURI = RecipeURIV


			case "TrainingDataSourceID":
				TrainingDataSourceIDV, _ := value.(string)
				createMLModel.TrainingDataSourceID = TrainingDataSourceIDV


			case "String":
				StringV, _ := value.(string)
				createMLModel.parameters.String = StringV
		}
	}

	params := make(map[string]string)

	preparecreateMLModelparams(params, createMLModel, Region)

	createMLModeljsonmap := make(map[string]interface{})

	preparecreateMLModelparamsdict(createMLModeljsonmap, createMLModel)

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, createMLModeljsonmap, response)
	resp = response
	*/
	return resp, err
}


func preparestarttaskparamsdict(createMLModeljsonmap map[string]interface{}, createMLModel CreateMLModel) {

	/*
	if createMLModel.MLModelId != "" {
		createMLModeljsonmap["MLModelId"] = createMLModel.MLModelId
	}
	if createMLModel.MLModelName != "" {
		createMLModeljsonmap["MLModelName"] = createMLModel.MLModelName
	}

	if createMLModel.MLModelType != "" {
		createMLModeljsonmap["MLModelType"] = createMLModel.MLModelType
	}
	if createMLModel.Recipe != "" {
		createMLModeljsonmap["Recipe"] = createMLModel.Recipe
	}

	if createMLModel.RecipeURI != "" {
		createMLModeljsonmap["RecipeURI"] = createMLModel.RecipeURI
	}

	if createMLModel.TrainingDataSourceID != "" {
		createMLModeljsonmap["TrainingDataSourceID"] = createMLModel.TrainingDataSourceID
	}


	if createMLModel.Parameters.String != "" {
		parameters := make(map[string]string)
		parameters["string"] = createMLModel.Parameters.String
		createMLModeljsonmap["parameters"] = parameters
	}
*/

}


func preparecreateMLModel(params map[string]string, createMLModel CreateMLModel, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonML_20141212.CreateMLModel"
}


func(awsmachinelearning *Awsmachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				MLModelIdV, _ := value.(string)
				MLModelId = MLModelIdV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	preparedeletemodel(params, MLModelId, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId ,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

func preparedeletemodel(params map[string]string, MLModelId string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.DeleteMLModel"
}

func preparegetmodel(params map[string]string, MLModelId string ,Verbose string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if Verbose != "" {
		params["Verbose"] = Verbose
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.GetMLModel"
}

func prepareupdatemodel(params map[string]string, MLModelId string ,ScoreThreshold string, MLModelName string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if MLModelName != "" {
		params["MLModelName"] = MLModelName
	}

	if ScoreThreshold != "" {
		params["ScoreThreshold"] = ScoreThreshold
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.UpdateMLModel"
}



//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (awsmachinelearning *Awsmachinelearning) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://machinelearning." + params["Region"] + ".amazonaws.com"
	fmt.Println("ECSEndpoint : ",ECSEndpoint)
	service := "machinelearning"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	fmt.Println("host : ",host)
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}

func(awsmachinelearning *Awsmachinelearning) GetMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, Verbose, Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				MLModelIdV, _ := value.(string)
				MLModelId = MLModelIdV

			case "Verbose":
				VerboseV, _ := value.(string)
				Verbose = VerboseV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	preparedeletemodel(params, MLModelId, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId,
		"Verbose" : Verbose,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

func (awsmachinelearning *Awsmachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, MLModelName, ScoreThreshold, Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				MLModelIdV, _ := value.(string)
				MLModelId = MLModelIdV

			case "MLModelName":
				MLModelNameV, _ := value.(string)
				MLModelName = MLModelNameV

			case "ScoreThreshold":
				ScoreThresholdV, _ := value.(string)
				ScoreThreshold = ScoreThresholdV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	prepareupdatemodel(params, MLModelId  ,ScoreThreshold , MLModelName ,Region)


	updatemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId ,
		"MLModelName": MLModelName,
		"ScoreThreshold" : ScoreThreshold,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, updatemodeljsonmap, response)
	resp = response
	return resp, err
}
