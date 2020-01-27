package config

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/ownerutil"

	"gopkg.in/yaml.v2"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type ProductConfig map[string]string

func NewManager(ctx context.Context, client k8sclient.Client, namespace string, configMapName string, installation *integreatlyv1alpha1.Installation) (*Manager, error) {
	cfgmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      configMapName,
		},
	}
	err := client.Get(ctx, k8sclient.ObjectKey{Name: configMapName, Namespace: namespace}, cfgmap)
	if !errors.IsNotFound(err) && err != nil {
		return nil, err
	}
	return &Manager{Client: client, Namespace: namespace, cfgmap: cfgmap, context: ctx, installation: installation}, nil
}

//go:generate moq -out ConfigReadWriter_moq.go . ConfigReadWriter
type ConfigReadWriter interface {
	readConfigForProduct(product integreatlyv1alpha1.ProductName) (ProductConfig, error)
	GetOauthClientsSecretName() string
	GetGHOauthClientsSecretName() string
	GetBackupsSecretName() string
	WriteConfig(config ConfigReadable) error
	ReadAMQStreams() (*AMQStreams, error)
	ReadRHSSO() (*RHSSO, error)
	ReadRHSSOUser() (*RHSSOUser, error)
	ReadCodeReady() (*CodeReady, error)
	ReadThreeScale() (*ThreeScale, error)
	ReadFuse() (*Fuse, error)
	ReadFuseOnOpenshift() (*FuseOnOpenshift, error)
	ReadAMQOnline() (*AMQOnline, error)
	GetOperatorNamespace() string
	ReadSolutionExplorer() (*SolutionExplorer, error)
	ReadMonitoring() (*Monitoring, error)
	ReadProduct(product integreatlyv1alpha1.ProductName) (ConfigReadable, error)
	ReadUps() (*Ups, error)
	ReadCloudResources() (*CloudResources, error)
}

//go:generate moq -out ConfigReadable_moq.go . ConfigReadable
type ConfigReadable interface {
	Read() ProductConfig
	GetProductName() integreatlyv1alpha1.ProductName
	GetProductVersion() integreatlyv1alpha1.ProductVersion
	GetOperatorVersion() integreatlyv1alpha1.OperatorVersion
	GetHost() string
	GetWatchableCRDs() []runtime.Object
}

type Manager struct {
	Client       k8sclient.Client
	Namespace    string
	cfgmap       *corev1.ConfigMap
	context      context.Context
	installation *integreatlyv1alpha1.Installation
}

func (m *Manager) ReadProduct(product integreatlyv1alpha1.ProductName) (ConfigReadable, error) {
	switch product {
	case integreatlyv1alpha1.Product3Scale:
		return m.ReadThreeScale()
	case integreatlyv1alpha1.ProductAMQOnline:
		return m.ReadAMQOnline()
	case integreatlyv1alpha1.ProductRHSSO:
		return m.ReadRHSSO()
	case integreatlyv1alpha1.ProductRHSSOUser:
		return m.ReadRHSSOUser()
	case integreatlyv1alpha1.ProductAMQStreams:
		return m.ReadAMQStreams()
	case integreatlyv1alpha1.ProductCodeReadyWorkspaces:
		return m.ReadCodeReady()
	case integreatlyv1alpha1.ProductFuse:
		return m.ReadFuse()
	case integreatlyv1alpha1.ProductFuseOnOpenshift:
		return m.ReadFuseOnOpenshift()
	case integreatlyv1alpha1.ProductSolutionExplorer:
		return m.ReadSolutionExplorer()
	case integreatlyv1alpha1.ProductUps:
		return m.ReadUps()
	case integreatlyv1alpha1.ProductCloudResources:
		return m.ReadCloudResources()
	}

	return nil, fmt.Errorf("no config found for product %v", product)
}

func (m *Manager) ReadSolutionExplorer() (*SolutionExplorer, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductSolutionExplorer)
	if err != nil {
		return nil, err
	}
	return NewSolutionExplorer(config), nil
}

func (m *Manager) GetOperatorNamespace() string {
	return m.Namespace
}

func (m *Manager) GetOauthClientsSecretName() string {
	return "oauth-client-secrets"
}

func (m *Manager) GetBackupsSecretName() string {
	return "backups-s3-credentials"
}

func (m *Manager) GetGHOauthClientsSecretName() string {
	return "github-oauth-secret"
}

func (m *Manager) ReadAMQStreams() (*AMQStreams, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductAMQStreams)
	if err != nil {
		return nil, err
	}
	return NewAMQStreams(config), nil
}

func (m *Manager) ReadThreeScale() (*ThreeScale, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.Product3Scale)
	if err != nil {
		return nil, err
	}
	return NewThreeScale(config), nil
}

func (m *Manager) ReadCodeReady() (*CodeReady, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductCodeReadyWorkspaces)
	if err != nil {
		return nil, err
	}
	return NewCodeReady(config), nil
}

func (m *Manager) ReadFuse() (*Fuse, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductFuse)
	if err != nil {
		return nil, err
	}
	return NewFuse(config), nil
}

func (m *Manager) ReadFuseOnOpenshift() (*FuseOnOpenshift, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductFuseOnOpenshift)
	if err != nil {
		return nil, err
	}
	return NewFuseOnOpenshift(config), nil
}

func (m *Manager) ReadRHSSO() (*RHSSO, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductRHSSO)
	if err != nil {
		return nil, err
	}
	return NewRHSSO(config), nil
}

func (m *Manager) ReadRHSSOUser() (*RHSSOUser, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductRHSSOUser)
	if err != nil {
		return nil, err
	}
	return NewRHSSOUser(config), nil
}

func (m *Manager) ReadAMQOnline() (*AMQOnline, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductAMQOnline)
	if err != nil {
		return nil, err
	}
	return NewAMQOnline(config), nil
}

func (m *Manager) ReadMonitoring() (*Monitoring, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductMonitoring)
	if err != nil {
		return nil, err
	}
	return NewMonitoring(config), nil
}

func (m *Manager) ReadUps() (*Ups, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductUps)
	if err != nil {
		return nil, err
	}

	return NewUps(config), nil
}

func (m *Manager) ReadCloudResources() (*CloudResources, error) {
	config, err := m.readConfigForProduct(integreatlyv1alpha1.ProductCloudResources)
	if err != nil {
		return nil, err
	}
	return NewCloudResources(config), nil
}

func (m *Manager) WriteConfig(config ConfigReadable) error {
	stringConfig, err := yaml.Marshal(config.Read())
	err = m.Client.Get(m.context, k8sclient.ObjectKey{Name: m.cfgmap.Name, Namespace: m.Namespace}, m.cfgmap)
	if errors.IsNotFound(err) {
		m.cfgmap.Data = map[string]string{string(config.GetProductName()): string(stringConfig)}
		ownerutil.EnsureOwner(m.cfgmap, m.installation)
		return m.Client.Create(m.context, m.cfgmap)
	} else {
		if m.cfgmap.Data == nil {
			m.cfgmap.Data = map[string]string{}
		}
		m.cfgmap.Data[string(config.GetProductName())] = string(stringConfig)
		ownerutil.EnsureOwner(m.cfgmap, m.installation)
		return m.Client.Update(m.context, m.cfgmap)
	}
}

func (m *Manager) readConfigForProduct(product integreatlyv1alpha1.ProductName) (ProductConfig, error) {
	config := m.cfgmap.Data[string(product)]
	decoder := yaml.NewDecoder(strings.NewReader(config))
	retConfig := ProductConfig{}
	if config == "" {
		return retConfig, nil
	}
	if err := decoder.Decode(retConfig); err != nil {
		return nil, fmt.Errorf("failed to decode product config for %v: %w", product, err)
	}
	return retConfig, nil
}
