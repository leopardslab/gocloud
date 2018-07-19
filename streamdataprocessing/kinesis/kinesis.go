package kinesis

//DeleteStream Delete Stream
func (kinesis *Kinesis) DeleteStream(request interface{}) (resp interface{}, err error) {
  param := request.(map[string]interface{})
  var streamName, Region string

  for key, value := range param {
    switch key {
    case "StreamName":
      streamNameV, _ := value.(string)
      streamName = streamNameV

    case "Region":
      RegionV, _ := value.(string)
      Region = RegionV
    }
  }

  params := make(map[string]string)
  preparedeletestream(params, Region)

  deletestreamjsonmap := make(map[string]interface{})
  preparedeletestreamdict(deletestreamjsonmap,streamName)

  response := make(map[string]interface{})
  err = kinesis.PrepareSignatureV4query(params, deletestreamjsonmap, response)
  resp = response
  return resp, err
}



func preparedeletestreamdict(deletestreamjsonmap map[string]interface{},streamName string){

  if streamName != "" {
		deletestreamjsonmap["StreamName"] = streamName
	}

  if limit != "" {
		deletestreamjsonmap["Limit"] = limit
	}
}

func preparedeletestream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.DeleteStream"
}




//DescribeStream describe stream
func (kinesis *Kinesis) DescribeStream(request interface{}) (resp interface{}, err error) {

  param := request.(map[string]interface{})
  var streamName,limit, exclusiveStartShardId, Region string

  for key, value := range param {
    switch key {
    case "StreamName":
      streamNameV, _ := value.(string)
      streamName = streamNameV

    case "Limit":
      limitV, _ := value.(string)
      limit = limitV

    case "ExclusiveStartShardId":
      exclusiveStartShardIdV, _ := value.(string)
      exclusiveStartShardId = exclusiveStartShardIdV

    case "Region":
      RegionV, _ := value.(string)
      Region = RegionV
    }
  }

  params := make(map[string]string)
  preparedescribestream(params, Region)


  describestreamjsonmap := make(map[string]interface{})
  preparedescribestreamdict(describestreamjsonmap,streamName,limit, exclusiveStartShardId)

  response := make(map[string]interface{})
  err = kinesis.PrepareSignatureV4query(params, describestreamjsonmap, response)
  resp = response
  return resp, err
}


func preparecreatestreamdict(describestreamjsonmap map[string]interface{},streamName string,limit string, exclusiveStartShardId string){

  if exclusiveStartShardId != "" {
		describestreamjsonmap["ExclusiveStartShardId"] = exclusiveStartShardId
	}

  if streamName != "" {
		describestreamjsonmap["StreamName"] = streamName
	}

  if limit != "" {
    describestreamjsonmap["Limit"] = limit
  }
}

func preparecreatestream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.DescribeStream"
}



//CreateStream Create Stream
func (kinesis *Kinesis) CreateStream(request interface{}) (resp interface{}, err error) {
  param := request.(map[string]interface{})
  var streamName, Region string

  for key, value := range param {
    switch key {
    case "StreamName":
      streamNameV, _ := value.(string)
      streamName = streamNameV

    case "Region":
      RegionV, _ := value.(string)
      Region = RegionV
    }
  }

  params := make(map[string]string)
  preparedeletestream(params, Region)

  deletestreamjsonmap := make(map[string]interface{})
  preparedeletestreamdict(deletestreamjsonmap,streamName)

  response := make(map[string]interface{})
  err = kinesis.PrepareSignatureV4query(params, deletestreamjsonmap, response)
  resp = response
  return resp, err
}



func preparecreatestreamdict(liststreamjsonmap map[string]interface{},streamName string,shardCount string){

  if shardCount != "" {
		createstreamjsonmap["ShardCount"] = shardCount
	}

  if streamName != "" {
		createstreamjsonmap["StreamName"] = streamName
	}
}

func preparecreatestream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.ListStreams"
}


//ListStream List Stream
func (kinesis *Kinesis) ListStream(request interface{}) (resp interface{}, err error) {

  	param := request.(map[string]interface{})
  	var exclusiveStartStreamName, limit, Region string

  	for key, value := range param {
  		switch key {
  		case "ExclusiveStartStreamName":
  			exclusiveStartStreamNameV, _ := value.(string)
  			exclusiveStartStreamName = exclusiveStartStreamNameV

      case "Limit":
        limitV, _ := value.(string)
        limit = limitV

  		case "Region":
  			RegionV, _ := value.(string)
  			Region = RegionV
  		}
  	}

  	params := make(map[string]string)
  	prepareliststream(params, Region)

  	liststreamjsonmap := make(map[string]interface{})
    prepareliststreamdict(liststreamjsonmap,exclusiveStartStreamName,limit)

  	response := make(map[string]interface{})
  	err = kinesis.PrepareSignatureV4query(params, liststreamjsonmap, response)
  	resp = response
  	return resp, err
}


func prepareliststreamdict(liststreamjsonmap map[string]interface{},exclusiveStartStreamName string,limit string){

  if exclusiveStartStreamName != "" {
		liststreamjsonmap["ExclusiveStartStreamName"] = exclusiveStartStreamName
	}

  if limit != "" {
		liststreamjsonmap["Limit"] = limit
	}
}

func prepareliststream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.ListStreams"
}
