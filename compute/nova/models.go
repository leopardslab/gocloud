package nova



//WIP
type  CreateServer struct {
  server Server `json:"server"`
	oSSCHHNTSchedulerHints OSSCHHNTSchedulerHints  `json:"OS-SCH-HNT:scheduler_hints"`
}

type OSSCHHNTSchedulerHints struct {
  SameHost string `json:"same_host"`
}


type Server struct {
  AccessIPv4 string `json:"accessIPv4"`

  AccessIPv6 string `json:"accessIPv6"`

  Name string `json:"name"`

  Networks string `json:"networks"`

  ImageRef string `json:"imageRef"`

  FlavorRef string `json:"flavorRef"`

  AvailabilityZone string `json:"availability_zone"`

  OSDCFDiskConfig string `json:"OS-DCF:diskConfig"`

  personality []Personality `json:"personality"`

  UserData string `json:"user_data"`

  metadata Metadata`json:"metadata"`

  securityGroups []SecurityGroups `json:"security_groups"`
}

type Personality struct {
  Path string `json:"path"`
  Contents string `json:"contents"`
}

type Metadata struct {
  MyServerName string `json:"My Server Name"`
}

type SecurityGroups struct {
  Name string `json:"name"`
}
