package awscontainer

type Creatcontainerresp struct {
	Cluster struct {
		ActiveServicesCount               int    `json:"activeServicesCount"`
		ClusterArn                        string `json:"clusterArn"`
		ClusterName                       string `json:"clusterName"`
		PendingTasksCount                 int    `json:"pendingTasksCount"`
		RegisteredContainerInstancesCount int    `json:"registeredContainerInstancesCount"`
		RunningTasksCount                 int    `json:"runningTasksCount"`
		Status                            string `json:"status"`
	} `json:"cluster"`
}
