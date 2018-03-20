package googleauth
    import "github.com/shlokgilda/gocloud/googleauth"


FUNCTIONS

func Sign() (token *oauth2.Token)
    sign() GCE signature it give URL to get Autorization code on which we it
    generate auth token and pass in each request in request header

func SignJWT() (client *http.Client)
    SignJWT reperesnts google service account authentication.

TYPES

type JWT struct {
    Type                    string `json:"type"`
    ProjectID               string `json:"project_id"`
    PrivateKeyID            string `json:"private_key_id"`
    PrivateKey              string `json:"private_key"`
    ClientEmail             string `json:"client_email"`
    ClientID                string `json:"client_id"`
    AuthURI                 string `json:"auth_uri"`
    TokenURI                string `json:"token_uri"`
    AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
    ClientX509CertURL       string `json:"client_x509_cert_url"`
}
    JWT struct reperesnts JWT json.


