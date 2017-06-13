package ec2


import(

	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
	//"reflect"

)

func (ec2 *EC2) StartInstances(ids ...string) (resp *StartInstanceResp, err error) {
	params := makeParams("StartInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &StartInstanceResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ec2 *EC2) StopInstances(ids ...string) (resp *StopInstanceResp, err error) {
	params := makeParams("StopInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &StopInstanceResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func (ec2 *EC2) RebootInstances(ids ...string) (resp *SimpleResp, err error) {
	params := makeParams("RebootInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &SimpleResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func (ec2 *EC2) TerminateInstances(instIds []string) (resp *TerminateInstancesResp, err error) {

	var ec2 RunInstances

	param := make(map[string]interface{})

	param = options.(map[string]interface{})

   	for key,value := range param{
		    switch key {
			      case "ImageId":
						    ImageId,_ := value.(string)
							  ec2.ImageId = ImageId

						case "MinCount":
							  MinCount,_ := value.(int)
								ec2.MinCount = MinCount

						case "MaxCount":
								MaxCount,_ := value.(int)
								ec2.MaxCount = MaxCount

						case "KeyName":
								KeyName,_ := value.(string)
								ec2.KeyName = KeyName

						case "KernelId":
								KernelId,_ := value.(string)
								ec2.KernelId = KernelId

						case "InstanceType":
								InstanceType,_ := value.(string)
								ec2.InstanceType = InstanceType

							case "RamdiskId":
							    RamdiskId,_ := value.(string)
								  ec2.RamdiskId = RamdiskId

							case "AvailZone":
								  AvailZone,_ := value.(string)
									ec2.AvailZone = AvailZone

							case "PlacementGroupName":
									PlacementGroupName,_ := value.(string)
									ec2.PlacementGroupName = PlacementGroupName

							case "Monitoring":
									Monitoring,_ := value.(bool)
									ec2.Monitoring = Monitoring

							case "SubnetId":
									SubnetId,_ := value.(string)
									ec2.SubnetId = SubnetId

							case "DisableAPITermination":
									DisableAPITermination,_ := value.(bool)
									ec2.DisableAPITermination = DisableAPITermination

							case "ShutdownBehavior":
									ShutdownBehavior,_ := value.(string)
									ec2.ShutdownBehavior = ShutdownBehavior

							case "PrivateIPAddress":
									PrivateIPAddress,_ := value.(string)
									ec2.PrivateIPAddress = PrivateIPAddress

							case "SecurityGroup":
									SecurityGroupparam,_ := value.([]map[string]string)
									//fmt.Println(SecurityGroupparam)
									for i := 0; i < len(SecurityGroupparam); i++{
										    var securityGroup SecurityGroup
										    securityGroup.Id = SecurityGroupparam[i]["Id"]
											  securityGroup.Name =SecurityGroupparam[i]["Name"]
						 					  ec2.SecurityGroups = append(ec2.SecurityGroups,securityGroup)
								 }

						 case "BlockDevice":
						    BlockDeviceparam,_ := value.([]map[string]interface{})
								for i := 0; i < len(BlockDeviceparam); i++{
									var BlockDeviceMappingParam BlockDeviceMapping
									    for BlockDeviceparamkey,BlockDeviceparamvalue := range BlockDeviceparam[i]{
										     switch BlockDeviceparamkey{
										         case "DeviceName":
											         BlockDeviceMappingParam.DeviceName = BlockDeviceparamvalue.(string)

										         case "VirtualName":
											         BlockDeviceMappingParam.VirtualName = BlockDeviceparamvalue.(string)

														 case "SnapshotId":
											         BlockDeviceMappingParam.SnapshotId = BlockDeviceparamvalue.(string)

										         case "VolumeType":
											         BlockDeviceMappingParam.VolumeType = BlockDeviceparamvalue.(string)

														 case "VolumeSize":
											         BlockDeviceMappingParam.VolumeSize = BlockDeviceparamvalue.(int64)

										         case "DeleteOnTermination":
											         BlockDeviceMappingParam.DeleteOnTermination = BlockDeviceparamvalue.(bool)

														 case "IOPS":
											         BlockDeviceMappingParam.IOPS = BlockDeviceparamvalue.(int64)

												 }
							        }
								 ec2.BlockDeviceMappings = append(ec2.BlockDeviceMappings,BlockDeviceMappingParam)

			         }

        }
    }

		
	return
}


func (ec2 *EC2) RunInstances(options *RunInstances) (resp *RunInstancesResp, err error) {
	params := prepareRunParams(*options)
	params["ImageId"] = options.ImageId
	params["InstanceType"] = options.InstanceType
	var min, max int
	if options.MinCount == 0 && options.MaxCount == 0 {
		min = 1
		max = 1
	} else if options.MaxCount == 0 {
		min = options.MinCount
		max = min
	} else {
		min = options.MinCount
		max = options.MaxCount
	}
	params["MinCount"] = strconv.Itoa(min)
	params["MaxCount"] = strconv.Itoa(max)
	i, j := 1, 1
	for _, g := range options.SecurityGroups {
		if g.Id != "" {
			params["SecurityGroupId."+strconv.Itoa(i)] = g.Id
			i++
		} else {
			params["SecurityGroup."+strconv.Itoa(j)] = g.Name
			j++
		}
	}
	prepareBlockDevices(params, options.BlockDeviceMappings)
	prepareNetworkInterfaces(params, options.NetworkInterfaces)
	token, err := clientToken()
	if err != nil {
		return nil, err
	}
	params["ClientToken"] = token

	if options.KeyName != "" {
		params["KeyName"] = options.KeyName
	}
	if options.KernelId != "" {
		params["KernelId"] = options.KernelId
	}
	if options.RamdiskId != "" {
		params["RamdiskId"] = options.RamdiskId
	}
	if options.UserData != nil {
		userData := make([]byte, base64.StdEncoding.EncodedLen(len(options.UserData)))
		base64.StdEncoding.Encode(userData, options.UserData)
		params["UserData"] = string(userData)
	}
	if options.AvailZone != "" {
		params["Placement.AvailabilityZone"] = options.AvailZone
	}
	if options.PlacementGroupName != "" {
		params["Placement.GroupName"] = options.PlacementGroupName
	}
	if options.Monitoring {
		params["Monitoring.Enabled"] = "true"
	}
	if options.SubnetId != "" {
		params["SubnetId"] = options.SubnetId
	}
	if options.DisableAPITermination {
		params["DisableApiTermination"] = "true"
	}
	if options.ShutdownBehavior != "" {
		params["InstanceInitiatedShutdownBehavior"] = options.ShutdownBehavior
	}
	if options.PrivateIPAddress != "" {
		params["PrivateIpAddress"] = options.PrivateIPAddress
	}

	resp = &RunInstancesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}





func (ec2 *EC2) query(params map[string]string, resp interface{}) error {

	req, err := http.NewRequest("GET", USEast.EC2Endpoint, nil)
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

	auth := Auth{"dummy","dummy"}


	SignV2(req,auth)

	fmt.Println(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	//fmt.Println(string(r.Body))

	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}
	if r.StatusCode != 200 {
		return buildError(r)
	}

	fmt.Println(xml.NewDecoder(r.Body).Decode(resp))
	return xml.NewDecoder(r.Body).Decode(resp)

}















