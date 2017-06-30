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
	fmt.Println("##########################################")

	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

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
	attachdisk := map[string]interface{}{
		"projectid": "sheltermap-1493101612061",
		"instance":  "sumesh-10",
		"Zone":      "us-east4-c",
		"Source":    "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
	}
	googlecloud.Attachdisk(attachdisk)

}
