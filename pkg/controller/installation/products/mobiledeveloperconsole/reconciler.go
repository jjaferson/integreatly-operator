package mobiledeveloperconsole

import (
	"context"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/marketplace"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/products/config"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	defaultInstallationNamespace = "mobile-developer-console"
	defaultSubscriptionName      = "integreatly-mdc"
	resourceName                 = "mobile-developer-console"
)

type Reconciler struct {
	Config        *config.MobileDeveloperConsole
	ConfigManager config.ConfigReadWriter
	mpm           marketplace.MarketplaceInterface
	logger        *logrus.Entry
	*resources.Reconciler
}

func NewReconciler(configManager config.ConfigReadWriter, instance *v1alpha1.Installation, mpm marketplace.MarketplaceInterface) (*Reconciler, error) {
	config, err := configManager.ReadMobileDeveloperConsole()
	if err != nil {
		return nil, errors.Wrap(err, "could not read mobile developer console")
	}

	if config.GetNamespace() == "" {
		config.SetNamespace(instance.Spec.NamespacePrefix + defaultInstallationNamespace)
	}
	if err = config.Validate(); err != nil {
		return nil, errors.Wrap(err, "nexus config is not valid")
	}

	logger := logrus.WithFields(logrus.Fields{"product": config.GetProductName()})

	return &Reconciler{
		ConfigManager: configManager,
		Config:        config,
		mpm:           mpm,
		logger:        logger,
		Reconciler:    resources.NewReconciler(mpm),
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, inst *v1alpha1.Installation, product *v1alpha1.InstallationProductStatus, client pkgclient.Client) (v1alpha1.StatusPhase, error) {

	phase, err := r.ReconcileNamespace(ctx, r.Config.GetNamespace(), inst, client)
	if err != nil || phase != v1alpha1.PhaseCompleted {
		return phase, err
	}

}
