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

func (mdc *MobileDeveloperConsole) SetNamespace(newNamespace string) {
	mdc.config["NAMESPACE"] = newNamespace
}

func (mdc *MobileDeveloperConsole) GetProductName() v1alpha1.ProductName {
	return v1alpha1.ProductMobileDeveloperConsole
}

func (n *MobileDeveloperConsole) GetProductVersion() string {
	return mdc.config["VERSION"]
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
