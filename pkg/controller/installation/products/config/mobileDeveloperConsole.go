package config

import (
	"errors"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
)

type MobileDeveloperConsole struct {
	config ProductConfig
}

func NewMobileDeveloperConsole(config ProductConfig) *MobileDeveloperConsole {
	return &MobileDeveloperConsole{config: config}
}

func (mdc *MobileDeveloperConsole) GetNamespace() string {
	return mdc.config["NAMESPACE"]
}

func (mdc *MobileDeveloperConsole) Read() ProductConfig {
	return mdc.config
}

func (mdc *MobileDeveloperConsole) SetNamespace(newNamespace string) {
	mdc.config["NAMESPACE"] = newNamespace
}

func (mdc *MobileDeveloperConsole) GetProductName() v1alpha1.ProductName {
	return v1alpha1.ProductMobileDeveloperConsole
}

func (mdc *MobileDeveloperConsole) GetProductVersion() v1alpha1.ProductVersion {
	return v1alpha1.VersionMobileDeveloperConsole
}

func (mdc *MobileDeveloperConsole) GetHost() string {
	return mdc.config["HOST"]
}

func (mdc *MobileDeveloperConsole) SetHost(newHost string) {
	mdc.config["HOST"] = newHost
}

func (mdc *MobileDeveloperConsole) Validate() error {
	if mdc.GetProductName() == "" {
		return errors.New("config product name is not defined")
	}
	if mdc.GetNamespace() == "" {
		return errors.New("config namespace is not defined")
	}
	return nil
}
