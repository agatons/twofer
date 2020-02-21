package config

import (
	"github.com/caarlos0/env/v6"
	"net/url"
	"sync"
)

type Config struct {
	Port int `env:"PORT" envDefault:"43210"`

	FrejaID FrejaID
	BankID  BankID

	QREnabled  bool `env:"QR_ENABLE" envDefault:"TRUE"`
	OTPEnabled bool `env:"OTP_ENABLE" envDefault:"TRUE"`
}

func (c Config) EIDEnabled() bool {
	return c.BankID.Enabled
}

type BankID struct {
	Enabled        bool     `env:"EID_BANKID_ENABLE"`
	URL            *url.URL `env:"EID_BANKID_URL"`
	RootCA         string   `env:"EID_BANKID_ROOT_CA_PEN"`
	RootCAFile     string   `env:"EID_BANKID_ROOT_CA_PEN_FILE"`
	ClientCert     string   `env:"EID_BANKID_CLIENT_CERT"`
	ClientCertFile string   `env:"EID_BANKID_CLIENT_CERT_FILE"`
	ClientKey      string   `env:"EID_BANKID_CLIENT_KEY"`
	ClientKeyFile  string   `env:"EID_BANKID_CLIENT_KEY_FILE"`
}

func (b BankID) GetRootCA() []byte {
	if b.RootCAFile != "" {
		return []byte(b.RootCAFile)
	}
	return []byte(b.RootCA)
}
func (b BankID) GetClientCert() []byte {
	if b.ClientCertFile != "" {
		return []byte(b.ClientCertFile)
	}
	return []byte(b.ClientCert)
}
func (b BankID) GetClientKey() []byte {
	if b.ClientKeyFile != "" {
		return []byte(b.ClientKeyFile)
	}
	return []byte(b.ClientKey)
}

type FrejaID struct {
	Enabled        bool     `env:"EID_FREJA_ENABLE"`
	URL            *url.URL `env:"EID_FREJA_URL"`
	RootCA         string   `env:"EID_FREJA_ROOT_CA_PEN"`
	RootCAFile     string   `env:"EID_FREJA_ROOT_CA_PEN_FILE"`
	ClientCert     string   `env:"EID_FREJA_CLIENT_CERT"`
	ClientCertFile string   `env:"EID_FREJA_CLIENT_CERT_FILE"`
	ClientKey      string   `env:"EID_FREJA_CLIENT_KEY"`
	ClientKeyFile  string   `env:"EID_FREJA_CLIENT_KEY_FILE"`
	JWSCert        string   `env:"EID_FREJA_JWS_CERT"`
	JWSCertFile    string   `env:"EID_FREJA_JWS_CERT_FILE"`
}

func (b FrejaID) GetRootCA() []byte {
	if b.RootCAFile != "" {
		return []byte(b.RootCAFile)
	}
	return []byte(b.RootCA)
}
func (b FrejaID) GetClientCert() []byte {
	if b.ClientCertFile != "" {
		return []byte(b.ClientCertFile)
	}
	return []byte(b.ClientCert)
}
func (b FrejaID) GetClientKey() []byte {
	if b.ClientKeyFile != "" {
		return []byte(b.ClientKeyFile)
	}
	return []byte(b.ClientKey)
}
func (b FrejaID) GetJWSCert() []byte {
	if b.ClientKeyFile != "" {
		return []byte(b.ClientKeyFile)
	}
	return []byte(b.ClientKey)
}

var once sync.Once
var config Config

func Get() Config {
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			//TODO something smart if things fail
			panic(err)
		}
	})

	return config
}
