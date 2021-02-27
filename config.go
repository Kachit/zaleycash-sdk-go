package zaleycash

type Config struct {
	Uri       string
	SecretKey string
	PublicKey string
}

func NewConfig(secretKey string, publicKey string) *Config {
	return &Config{Uri: ProdAPIUrl, SecretKey: secretKey, PublicKey: publicKey}
}
