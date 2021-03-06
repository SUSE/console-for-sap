package cloud

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	azure = "azure"
)

type CloudInstance struct {
	Provider string      `mapstructure:"provider,omitempty"`
	Metadata interface{} `mapstructure:"metadata,omitempty"`
}

type CustomCommand func(name string, arg ...string) *exec.Cmd

var customExecCommand CustomCommand = exec.Command

func IdentifyCloudProvider() (string, error) {
	log.Info("Identifying if the VM is running in a cloud environment...")
	output, err := customExecCommand("dmidecode", "-s", "bios-vendor").Output()
	if err != nil {
		return "", err
	}

	provider := strings.TrimSpace(string(output))
	log.Debugf("dmidecode output: %s", provider)

	switch string(provider) {
	case "Microsoft Corporation":
		log.Infof("VM is running on %s", azure)
		return azure, nil
	default:
		log.Info("VM is not running in any recognized cloud provider")
		return "", nil
	}

}

func NewCloudInstance() (*CloudInstance, error) {
	var err error
	var cloudMetadata interface{}

	provider, err := IdentifyCloudProvider()
	if err != nil {
		return nil, err
	}

	cInst := &CloudInstance{
		Provider: provider,
	}

	switch provider {
	case azure:
		cloudMetadata, err = NewAzureMetadata()
		if err != nil {
			return nil, err
		}
	}

	cInst.Metadata = cloudMetadata

	return cInst, nil

}
