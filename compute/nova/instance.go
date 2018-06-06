package nova

import (
	"fmt"
)

//WIP
func (nova *Nova) Startnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

func (nova *Nova) Stopnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

func (nova *Nova) Rebootnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

func (nova *Nova) Deletenode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//WIP
func (nova *Nova) Createnode(request interface{}) (resp interface{}, err error) {
	var options CreateServer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	fmt.Println("param", param)

	for key, value := range param {

		switch key {
		case "Server":
			ServerV, _ := value.(map[string]interface{})

			var server Server

			for serverkey, servervalue := range ServerV {

				switch serverkey {

				case "Name":
					nameV, _ := servervalue.(string)
					server.Name = nameV

				case "AccessIPv4":
					accessIPv4V, _ := servervalue.(string)
					server.AccessIPv4 = accessIPv4V

				case "AccessIPv6":
					accessIPv6V, _ := servervalue.(string)
					server.AccessIPv6 = accessIPv6V

				case "ImageRef":
					imageRefV, _ := servervalue.(string)
					server.ImageRef = imageRefV

				case "FlavorRef":
					flavorRefV, _ := servervalue.(string)
					server.FlavorRef = flavorRefV

				case "AvailabilityZone":
					availabilityZoneV, _ := servervalue.(string)
					server.FlavorRef = availabilityZoneV

				case "OSDCFDiskConfig":
					oSDCFDiskConfigV, _ := servervalue.(string)
					server.OSDCFDiskConfig = oSDCFDiskConfigV

				case "Personality":
					personalityV, _ := servervalue.([]map[string]string)

					for i := 0; i < len(personalityV); i++ {
						var personality Personality
						personality.Path = personalityV[i]["Path"]
						personality.Contents = personalityV[i]["Contents"]
						server.personality = append(server.personality, personality)
					}

				case "UserData":
					userDataV, _ := servervalue.(string)
					server.UserData = userDataV

				case "metadata":
					metadataV, _ := servervalue.(map[string]string)
					var metadata Metadata
					metadata.MyServerName = metadataV["MyServerName"]
					server.metadata = metadata

				case "Networks":
					networksV, _ := servervalue.(string)
					server.Networks = networksV

				case "securityGroups":
					securityGroupsV, _ := servervalue.([]map[string]string)
					for i := 0; i < len(securityGroupsV); i++ {
						var securityGroup SecurityGroups
						securityGroup.Name = securityGroupsV[i]["Name"]
						server.securityGroups = append(server.securityGroups, securityGroup)
					}

				}

			}

			options.server = server

		case "oSSCHHNTSchedulerHints":
			oSSCHHNTSchedulerHintsv, _ := value.(string)
			options.oSSCHHNTSchedulerHints.SameHost = oSSCHHNTSchedulerHintsv
		}
	}
	fmt.Println(options)

	return resp, err

}
