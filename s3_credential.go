package paperspace

type S3Credential struct {
	AccessKey string `json:"accessKey"`
	Bucket    string `json:"bucket"`
	Endpoint  string `json:"endpoint,omitempty`
	SecretKey string `json:"secretKey,omitempty"`
}
