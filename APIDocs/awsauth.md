package awsauth
    import "github.com/scorelab/gocloud-v2/awsauth"


FUNCTIONS

func SignatureV2(req *http.Request, Auth interface{}) (err error)
    SignatureV2 method for Authenticating request

func SignatureV4(request *http.Request, request_parameters []byte, amztarget string, method string, region string, service string, host string, ContentType string, signedheaders string) (signrequest *http.Request)


