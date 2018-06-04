package awsmachinelearning

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
)

//CreateMLModel creates model.
func (awsmachinelearning *Awsmachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var Region string
	var createMLModel CreateMLModel

	for key, value := range param {
		switch key {
		case "MLModelId":
			MLModelIdV, _ := value.(string)
			createMLModel.MLModelID = MLModelIdV

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

	preparecreateMLModel(params, createMLModel, Region)

	createMLModeljsonmap := make(map[string]interface{})

	preparecreateMLModelparamsdict(createMLModeljsonmap, createMLModel)

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, createMLModeljsonmap, response)
	resp = response

	return resp, err
}

//DeleteMLModel delete model.
func (awsmachinelearning *Awsmachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error) {

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
		"MLModelId": MLModelId,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

//GetMLModel describe model.
func (awsmachinelearning *Awsmachinelearning) GetMLModel(request interface{}) (resp interface{}, err error) {

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
		"Verbose":   Verbose,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

//UpdateMLModel update model.
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

	prepareupdatemodel(params, MLModelId, ScoreThreshold, MLModelName, Region)

	updatemodeljsonmap := map[string]interface{}{
		"MLModelId":      MLModelId,
		"MLModelName":    MLModelName,
		"ScoreThreshold": ScoreThreshold,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, updatemodeljsonmap, response)
	resp = response
	return resp, err
}
