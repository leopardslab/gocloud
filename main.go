package main

import (
	"fmt"
	"github.com/scorelab/gocloud-v2/gocloud"
)

func main() {

	/*
	    InitializeParams := map[string]string{
	    "SourceImage" : "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
	    "DiskType" :  "1",
	    "DiskSizeGb" :"1",
	  }

	    disk := []map[string]interface{}{
	    {
	    "Type": "usb",
	    "Boot":  true,
	    "Mode": "",
	    "AutoDelete": false,
	    "deviceName": "",
	    "InitializeParams": InitializeParams,
	  },

	    {
	    "Type": "uss",
	    "Boot":  true,
	    "Mode": "",
	    "AutoDelete": false,
	    "deviceName": "",
	    "InitializeParams": InitializeParams,
	  },
	  }

	    AccessConfigs := []map[string]string{{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	    },{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	    }}

	  AccessConfigs2 := []map[string]string{{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	      },}


	    NetworkInterfaces := []map[string]interface{}{{
	    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	    "Subnetwork": "",
	    "AccessConfigs": AccessConfigs,
	  },

	    {
	    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	    "Subnetwork": "hello",
	    "AccessConfigs": AccessConfigs2,
	  },
	  }

	  scheduling := map[string]interface{}{
	    "Preemptible": false,
	    "OnHostMaintenance": "MIGRATE",
	    "AutomaticRestart": true,
	  }


	    gce := map[string]interface{}{
	    "projectid":"sheltermap-1493101612061",
	    "Name": "scorelab",
	    "MachineType":"n1-standard-1",
	    "Zone":"us-east4-c",
	    "disk": disk,
	    "NetworkInterfaces": NetworkInterfaces,
	    "description":"",
	    "scheduling" : scheduling,
	    }
	*/

	/*
	   start := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider, gocloud.Secretkey, gocloud.Secretid)
	//  googlecloud.Startnode(start)

	/*
	   stop := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//  googlecloud.Stopnode(stop)

	/*
	   reboot := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//googlecloud.Rebootnode(reboot)

	/*
	   delete := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",
	   }

	   googlecloud.Deletenode(delete)
	*/
	/*
	    ec2 := map[string]interface{}{
	      "ImageId"     : "ami-ccf405a5",
	      "InstanceType": "t1.micro",
	    }

	    amazoncloud, _ := CloudProvider(Amazonprovider,Secretkey,Secretid)
	    amazoncloud.Createnode(ec2)


	    reboot := []string{
	        "i-0ce8be9072b64fe84",
	    }

	   amazoncloud.Stopnode(stop)

	    start := []string{
	        "i-002ebb744a5fc48e0",
	    }
	    amazoncloud.Startnode(start)


	    amazoncloud.Deletenode(start)

	    amazoncloud.Rebootnode(reboot)
	*/

	/*
		InitializeParams := map[string]string{
			"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
			"DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
			"DiskSizeGb":  "10",
		}

		disk := []map[string]interface{}{
			{
				"Boot":             true,
				"AutoDelete":       false,
				"DeviceName":       "bokya",
				"Type":             "PERSISTENT",
				"Mode":             "READ_WRITE",
				"InitializeParams": InitializeParams,
			},
		}

		AccessConfigs := []map[string]string{{
			"Name": "external-nat",
			"Type": "ONE_TO_ONE_NAT",
		},
		}

		NetworkInterfaces := []map[string]interface{}{
			{
				"Network":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
				"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
				"AccessConfigs": AccessConfigs,
			},
		}

		gce := map[string]interface{}{
			"projectid":         "sheltermap-1493101612061",
			"Name":              "sumesh-10",
			"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
			"Zone":              "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c",
			"disk":              disk,
			"NetworkInterfaces": NetworkInterfaces,
		}

		googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
		googlecloud.Createnode(gce)
	*/
	/*
		ec2 := map[string]interface{}{
			"ImageId":      "ami-ccf405a5",
			"InstanceType": "t1.micro",
		}

		amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

		amazoncloud.Createnode(ec2)



	*/
	/*
	   	start := map[string]string{
	   		"projectid":"sheltermap-1493101612061",
	   		"instance":"sumesh-10",
	   		"Zone": "us-east4-c",
	   	}
	     googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	     googlecloud.Startnode(start)
	*/
	fmt.Println("********************")
	/*
		disk := map[string]interface{}{
			"projectid": "sheltermap-1493101612061",
			"Name":      "disk-11",
			"Type":      "pd-standard",
			"Zone":      "us-east4-c",
		}

		googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
		googlecloud.Createdisk(disk)

		fmt.Println("")

		diskDeletedisk := map[string]string{
			"projectid": "sheltermap-1493101612061",
			"disk":      "disk-10",
			"Zone":      "us-east4-c",
		}

		googlecloud.Deletedisk(diskDeletedisk)
	*/
	/*
		createsnapshot := map[string]interface{}{
			"projectid": "sheltermap-1493101612061",
			"disk":      "disk-11",
			"Zone":      "us-east4-c",
			"Name":      "disk-12",
		}
	*/
	/*
		googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

		//	googlecloud.Createsnapshot(createsnapshot)

		deletesnapshot := map[string]string{
			"projectid": "sheltermap-1493101612061",
			"snapshot":  "gg",
		}

		googlecloud.Deletesnapshot(deletesnapshot)
	*/

	/*
		googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

		attachdisk := map[string]interface{}{
			"projectid": "sheltermap-1493101612061",
			"instance":      "sumesh-10",
			"Zone":      "us-east4-c",
			"Source":"projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
		}
		googlecloud.Attachdisk(attachdisk)

	*/

	/*
	   	ec2 := map[string]interface{}{
	    	 "ImageId"     : "ami-ccf405a5",
	    	 "InstanceType": "t1.micro",
	   	 "Region" :"us-east-1",

	     }

	     amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	     amazoncloud.Createnode(ec2)

	   	start := map[string]string{
	   		"instance-id":	"i-0174bd6f54178e89b",
	   		"Region"     :	"us-east-1",
	   	}
	   	//amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	   	amazoncloud.Startnode(start)

	   	stop := map[string]string{
	   		"instance-id":	"i-0ec61e05211ceadad",
	   		"Region"     :	"us-east-1",
	   	}
	     amazoncloud.Stopnode(stop)




	   	InitializeParams := map[string]string{
	   		"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
	   		"DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
	   		"DiskSizeGb":  "10",
	   	}

	   	disk := []map[string]interface{}{
	   		{
	   			"Boot":             true,
	   			"AutoDelete":       false,
	   			"DeviceName":       "bokya",
	   			"Type":             "PERSISTENT",
	   			"Mode":             "READ_WRITE",
	   			"InitializeParams": InitializeParams,
	   		},
	   	}

	   	AccessConfigs := []map[string]string{{
	   		"Name": "external-nat",
	   		"Type": "ONE_TO_ONE_NAT",
	   	},
	   	}

	   	NetworkInterfaces := []map[string]interface{}{
	   		{
	   			"Network":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	   			"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
	   			"AccessConfigs": AccessConfigs,
	   		},
	   	}

	   	gce := map[string]interface{}{
	   		"projectid":         "sheltermap-1493101612061",
	   		"Name":              "scorelab-instance-11",
	   		"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
	   		"Zone":              "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c",
	   		"disk":              disk,
	   		"NetworkInterfaces": NetworkInterfaces,
	   	}
	*/
	/*
	   	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	   //	googlecloud.Createnode(gce)


	   	start := map[string]string{
	   		"projectid":"sheltermap-1493101612061",
	   		"instance":"sumesh-10",
	   		"Zone": "us-east4-c",
	   	}

	     googlecloud.Startnode(start)
	*/
	/*
		createdisk := map[string]interface{}{
		//	"IOPS":       int64(3000),
			"AvailZone":  "us-east-1a",
		//	"VolumeType": "gp2",
		  "VolumeSize": 100,
		}
	*/
	//amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	//amazoncloud.Createdisk(createdisk)

	/*
		deletedisk := map[string]string{
			"VolumeId": "vol-05371175f10d2e6a4",
		}

		amazoncloud.Deletedisk(deletedisk)
	*/
	/*
		createsnapshot := map[string]string{
			"VolumeId": "vol-047d011f7536d2b7c",
			"Description":"",
		}

		amazoncloud.Createsnapshot(createsnapshot)
	*/

	/*
		deletesnapshot := map[string]string{
			"SnapshotId": "snap-0f0839076356ce6cb",
		}
		amazoncloud.Deletesnapshot(deletesnapshot)

		ec2 := map[string]interface{}{
			"ImageId":      "ami-ccf405a5",
			"InstanceType": "t1.micro",
			"Region":       "us-east-1",
		}

		amazoncloud.Createnode(ec2)

		attachdisk := map[string]string{
			"VolumeId":   "vol-049426a70587418d7",
			"InstanceId": "i-07737a0121fba7b0c",
			"Device":     "/dev/sdh",
		}

		amazoncloud.Attachdisk(attachdisk)

		detachdisk := map[string]string{
			"VolumeId":   "vol-049426a70587418d7",
			"InstanceId": "i-07737a0121fba7b0c",
			"Device":     "/dev/sdh",
			"Force":      "true",
		}
		amazoncloud.Detachdisk(detachdisk)
	*/
	//	fmt.Println("##########################################")

	//	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	/*
		detachdiskg := map[string]string{
			"projectid":  "sheltermap-1493101612061",
			"instance":   "sumesh-10",
			"Zone":       "us-east4-c",
			"deviceName": "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
		}

		googlecloud.Detachdisk(detachdiskg)
	*/

	/*
		attachdisk := map[string]interface{}{
			"projectid": "sheltermap-1493101612061",
			"instance":  "sumesh-10",
			"Zone"    :  "us-east4-c",
			"Source"  :  "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
		}
		googlecloud.Attachdisk(attachdisk)
	*/

	/*	SourceImageEncryptionKeys := map[string]string{
			"RawKey": "1",
			"Sha256": "2",
		}

		Licenses := []string{"GPL", "apache"}

		disk := map[string]interface{}{
			"projectid":                 "sheltermap-1493101612061",
			"Name":                      "disk-21",
			"Type":                      "pd-standard",
			"Zone":                      "us-east4-c",
			"SizeGb":                    "20",
			"Licenses":                  Licenses,
			"SourceImageEncryptionKeys": SourceImageEncryptionKeys,
		}

		googlecloud.Createdisk(disk)
	*/

	/*
		Licenses := []string{"GPL", "apache"}

		SourceImageEncryptionKeys := map[string]string{
			"RawKey": "12",
			"Sha256": "14",
		}

		InitializeParams := map[string]interface{}{
			"DiskName":                  "DiskName",
			"DiskType":                  "DiskType",
			"DiskSizeGb":                "DiskSizeGb",
			"SourceImage":               "SourceImage",
			"SourceImageEncryptionKeys": SourceImageEncryptionKeys,
		}

	*/
	/*
		attachdisk := map[string]interface{}{
			"projectid": "sheltermap-1493101612061",
			"instance":  "sumesh-10",
			"Zone":      "us-east4-c",
			"Source":    "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
		}
		googlecloud.Attachdisk(attachdisk)

		amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

		attachdisk1 := map[string]string {
			"VolumeId":   "vol-049426a70587418d7",
			"InstanceId": "i-07737a0121fba7b0c",
			"Device":     "/dev/sdh",
			"Region":     "us-east-1",
		}

		amazoncloud.Attachdisk(attachdisk1)

		detachdisk := map[string]string{
			"VolumeId":   "vol-049426a70587418d7",
			"InstanceId": "i-07737a0121fba7b0c",
			"Device":     "/dev/sdh",
			"Force":      "true",
			"Region":     "us-east-1",
		}
		amazoncloud.Detachdisk(detachdisk)
	*/
	/*
		Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

		creatloadbalancer := map[string]interface{}{
			"Name":    "my-load-balancer",
			"Subnets": Subnets,
		}

		amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
		//	amazoncloud.Createnode(ec2)

		amazoncloud.Creatloadbalancer(creatloadbalancer)

		/*
			deleteloadbalancer := map[string]string{
				"LoadBalancerArn": "my-load-balancer",
			}
			amazoncloud.Deleteloadbalancer(deleteloadbalancer)

		listloadbalancer := map[string]string{
			"LoadBalancerArn": "arn:aws:elasticloadbalancing:us-east-1:350088531623:loadbalancer/app/my-load-balancer/0ab412f3c9473ea7",
		}

		amazoncloud.Listloadbalancer(listloadbalancer)
	*/
	/*
		Listeners := []map[string]string{{
			"InstancePort":     "80",
			"LoadBalancerPort": "80",
			"Protocol":         "http",
			"InstanceProtocol": "http",
			"SSLCertificateId": "",
		},
		}

		Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

		creatloadbalancer := map[string]interface{}{
			"LoadBalancerName": "my-load-balancer",
			"Listeners":        Listeners,
			"Subnets":          Subnets,
		}

		fmt.Println(creatloadbalancer)

		amazoncloud.Creatloadbalancer(creatloadbalancer)

		deleteloadbalancer := map[string]string{
			"LoadBalancerName": "my-load-balancer",
		}

		amazoncloud.Deleteloadbalancer(deleteloadbalancer)

		amazoncloud.Listloadbalancer(nil)

	*/
	//amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	/*
		attachnodewithloadbalancer := map[string]interface{}{
			"Instances":        []string{"i-05f4f2535c41b680b"},
			"LoadBalancerName": "my-load-balancer",
		}
	*/
	//amazoncloud.Attachnodewithloadbalancer(attachnodewithloadbalancer)
	//amazoncloud.Detachnodewithloadbalancer(attachnodewithloadbalancer)

	/*
		ec2 := map[string]interface{}{
			"ImageId":      "ami-ccf405a5",
			"InstanceType": "t1.micro",
			"Region":       "us-east-1",
		}

		amazoncloud.Createnode(ec2)
	*/
	/*
		creatcontainer := map[string]interface{}{
			"clusterName": "rootmonk",
			"Region":      "us-east-1",
		}
	*/
	//	amazoncloud, _ = gocloud.CloudProvider(gocloud.Amazonprovider)
	//amazoncloud.Createcontainer(creatcontainer)
	//	amazoncloud.Deletecontainer(creatcontainer)
	/*
		LoadBalancers := []map[string]interface{}{{
			"containerName":    "rootmonk",
			"loadBalancerName": "us-east-2",
		}, {
			"containerName":    "rootmonk",
			"loadBalancerName": "us-east-2",
		},
		}
		createservice := map[string]interface{}{
			"serviceName":    "ecs-simple-service",
			"taskDefinition": "ecs-demo",
			"desiredCount":   1,
			"Region":         "us-east-1",
			"LoadBalancers":  LoadBalancers,
		}
	*/
	/*
		environment := []map[string]string{
			{
				"name":  "name1",
				"value": "value1",
			}, {
				"name":  "name2",
				"value": "value2",
			},
		}
		containerOverrides := []map[string]interface{}{{
			"command":           []string{"command1", "command2"},
			"cpu":               2,
			"name":              "bokya",
			"memory":            2,
			"memoryReservation": "memoryReservation",
			"environment":       environment,
		},
		}

		overrides := map[string]interface{}{
			"taskRoleArn":        "taskRoleArn",
			"containerOverrides": containerOverrides,
		}
	*/
	/*
		runtask := map[string]interface{}{
			//	"overrides": overrides,
			"taskDefinition": "string",
			"Region":         "us-east-1",
		}
	*/
	//	amazoncloud.Runtask(runtask)
	/*
		deleteservice := map[string]interface{}{
			"cluster": "cluster",
			"service": "service",
			"Region":  "us-east-1",
		}
		amazoncloud.Deleteservice(deleteservice)
	*/
	/*
		stoptask := map[string]interface{}{
			"cluster": "cluster",
			"reason":  "reason",
			"task":    "task",
			"Region":  "us-east-1",
		}
		amazoncloud.Stoptask(stoptask)

	*/
	/*
		delete := map[string]string{
			"projectid": "sheltermap-1493101612061",
			"instance":  "instance-10",
			"Zone":      "us-west1-c",
		}

		googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
		googlecloud.Deletenode(delete)

	*/
	/*
		amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

		createdns := map[string]interface{}{
			"name":             "rootmonk.me",
			"hostedZoneConfig": "hostedZoneConfig",
		}

		amazoncloud.Createdns(createdns)

		deletedns := map[string]string{
			"ID": "ZOD7SUP0ZGGQQ",
		}
		amazoncloud.Deletedns(deletedns)

		listdns := map[string]interface{}{
			"marker":   "",
			"maxItems": 2,
		}

		amazoncloud.Listdns(listdns)

		for i := 0; i < 5; i++ {
			fmt.Println("****************")
		}
		listResourcednsRecordSets := map[string]interface{}{
			"zone": "ZBNX5TIW033J2",
		}
		amazoncloud.ListResourcednsRecordSets(listResourcednsRecordSets)

		/*
			resourceRecordSet := map[string]interface{}{
				"name":    "http://rootmonk.me",
				"type":    "A",
				"ttl":     300,
				"records": []string{"127.0.0.1"},
			}

			fmt.Println(resourceRecordSet)

			changes := []map[string]interface{}{
				{
					"action": "CREATE",
					//		"record": resourceRecordSet,
				},
				{
					"action": "CREATE",
					//	"record": resourceRecordSet,
				},
			}
			changednsrecordsets := map[string]interface{}{
				"zone":    "Z3DTMMYMFT5XOZ",
				"comment": "Helloworld",
				"changes": changes,
			}

			amazoncloud.Changednsrecordsets(changednsrecordsets)
	*/

	/*
	   	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	   	reboot := map[string]string{
	   	"projectid":"sheltermap-1493101612061",
	   	"instance":"testing-scorelab",
	   	"Zone": "us-east4-c",
	   	}
	      resp,err := googlecloud.Rebootnode(reboot)

	   	 if(err!=nil){
	   		    fmt.Println("Request fail")
	   	 }else{
	   	 		response := resp.(map[string]interface{})
	   	 		fmt.Println(response["body"])
	   	 		fmt.Println(response["status"])
	   			fmt.Println("Request pass")

	    }
	*/
	/*
	    Listeners := []map[string]string{{
	   		"InstancePort":     "80",
	   		"LoadBalancerPort": "80",
	   		"Protocol":         "http",
	   		"InstanceProtocol": "http",
	   		"SSLCertificateId": "",
	   	},
	   	}

	   	Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

	   	creatloadbalancer := map[string]interface{}{
	   		"LoadBalancerName": "my-load-balancer",
	   		"Listeners":        Listeners,
	   		"Subnets":          Subnets,
	   	}

	   	fmt.Println(creatloadbalancer)
	   	amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	   	amazoncloud.Creatloadbalancer(creatloadbalancer)


	   /////use code from here
	   	ec2 := map[string]interface{}{
	   		"ImageId":      "ami-ccf405a5",
	   		"InstanceType": "t1.micro",
	   		"Region" :"us-east-1",
	   	}

	   	amazoncloud, _ = gocloud.CloudProvider(gocloud.Amazonprovider)

	   	amazoncloud.Createnode(ec2)


	   	stop := map[string]string{
	   		"instance-id": "i-0dca1c71599785f4b",
	   		"Region":      "us-east-1",
	   	}

	   	amazoncloud.Stopnode(stop)
	*/

	/* google storage API
	   	disk := map[string]interface{}{
	   		"projectid": "sheltermap-1493101612061",
	   		"Name":      "disk-11",
	   		"Type":      "pd-standard",
	   		"Zone":      "us-east4-c",
	   	}

	   	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	   	resp ,_  := googlecloud.Createdisk(disk)

	   	response := resp.(map[string]interface{})

	    	fmt.Println(response["body"])


	   	createsnapshot := map[string]interface{}{
	   		"projectid": "sheltermap-1493101612061",
	   		"disk":      "disk-11",
	   		"Zone":      "us-east4-c",
	   		"Name":      "disk-12",
	   	}

	     resp ,_  =  googlecloud.Createsnapshot(createsnapshot)

	   	response = resp.(map[string]interface{})

	   	fmt.Println(response["body"])

	   	deletesnapshot := map[string]string{
	   		"projectid": "sheltermap-1493101612061",
	   		"snapshot":  "disk-12",
	   	}


	   	resp ,_  = googlecloud.Deletesnapshot(deletesnapshot)

	   	response = resp.(map[string]interface{})

	    	fmt.Println(response["body"])


	   	attachdisk := map[string]interface{}{
	   		"projectid": "sheltermap-1493101612061",
	   		"instance":      "sumesh-10",
	   		"Zone":      "us-east4-c",
	   		"Source":"projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
	   	}

	     resp, _  = googlecloud.Attachdisk(attachdisk)

	   	response = resp.(map[string]interface{})

	    	fmt.Println(response["body"])

	      detachdisk := map[string]string{
	   		"projectid":  "sheltermap-1493101612061",
	   		"instance":   "sumesh-10",
	   		"Zone":       "us-east4-c",
	   		"deviceName": "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
	   	}

	   	resp, _ = googlecloud.Detachdisk(detachdisk)

	   	response = resp.(map[string]interface{})

	   	fmt.Println(response["body"])


	   	deletedisk := map[string]string{
	   		"VolumeId": "vol-05371175f10d2e6a4",
	   	}

	     resp, _ = googlecloud.Deletedisk(deletedisk)

	   	response = resp.(map[string]interface{})

	    	fmt.Println(response["body"])

	*/

	/*amazon storage API


	  	createdisk := map[string]interface{}{
	  		"AvailZone":  "us-east-1a",
	  	  "VolumeSize" : 100,
	  		"Region":  "us-east-1",
	  	}

	   amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

	   amazoncloud.Createdisk(createdisk)


	  	deletedisk := map[string]string{
	  		"VolumeId": "vol-0996a16ff8f032760",
	  		"Region":  "us-east-1",
	  	}

	  	amazoncloud.Deletedisk(deletedisk)

	  	fmt.Println("I am here")


	  	attachdisk := map[string]string {
	  		"VolumeId":   "vol-049426a70587418d7",
	  		"InstanceId": "i-0050d952f9b8d45d5",
	  		"Device":     "/dev/sdh",
	  		"Region":     "us-east-1",
	  	}

	  	amazoncloud.Attachdisk(attachdisk)

	  	detachdisk := map[string]string{
	  		"VolumeId":   "vol-049426a70587418d7",
	  		"InstanceId": "i-0050d952f9b8d45d5",
	  		"Device":     "/dev/sdh",
	  		"Force":      "true",
	  		"Region":     "us-east-1",
	  	}
	  	amazoncloud.Detachdisk(detachdisk)

	  	createsnapshot := map[string]string{
	  		"VolumeId": "vol-047d011f7536d2b7c",
	  		"Description":"",
	  		"Region":"us-east-1",
	  	}

	  	amazoncloud.Createsnapshot(createsnapshot)

	  	deletesnapshot := map[string]string{
	  		"SnapshotId": "snap-0f0839076356ce6cb",
	  		"Region":     "us-east-1",
	  	}
	  	amazoncloud.Deletesnapshot(deletesnapshot)

	*/

	/*loadbalancer api aws
	Listeners := []map[string]string{{
		"InstancePort":     "80",
		"LoadBalancerPort": "80",
		"Protocol":         "http",
		"InstanceProtocol": "http",
		"SSLCertificateId": "",
	},
	}

	Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

	creatloadbalancer := map[string]interface{}{
		"LoadBalancerName": "my-load-balancer",
		"Listeners":        Listeners,
		"Subnets":          Subnets,
	}

	amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
	resp, _ := amazoncloud.Creatloadbalancer(creatloadbalancer)
	response := resp.(map[string]interface{})

	fmt.Println("***************")
	fmt.Println(response["body"])

	*/
	/*
	   creatloadbalancer := map[string]interface{}{
	   	"Project": "sheltermap-1493101612061",
	   	"name":  "google-loadbalancer",
	   	"Region":"us-east4-c",
	   	"instances": []string{"testing-scorelab"},
	   }

	   googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
	   googlecloud.Creatloadbalancer(creatloadbalancer)

	*/
	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	createdns := map[string]interface{}{
		"Project":     "sheltermap-1493101612061",
		"Kind":        "dns#managedZone",
		"Description": "dns",
		"DnsName":     "rootmonk.me.",
		"Name":        "gocloud3",
	}
	googlecloud.Createdns(createdns)

	listdns := map[string]string{
		"Project": "sheltermap-1493101612061",
	}

	googlecloud.Listdns(listdns)

	listResourcednsRecordSets := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	googlecloud.ListResourcednsRecordSets(listResourcednsRecordSets)

	deletedns := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	googlecloud.Deletedns(deletedns)

}
