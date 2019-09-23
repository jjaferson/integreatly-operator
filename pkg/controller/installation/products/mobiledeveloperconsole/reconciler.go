package mobiledeveloperconsole

import (
	"context"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/marketplace"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/products/config"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	defaultInstallationNamespace = "mobile-developer-console"
	defaultSubscriptionName      = "integreatly-mobile-developer-console"
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

	err = config.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "mdc config is not valid")
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

	phase, err = r.ReconcileSubscription(ctx, inst, marketplace.Target{Pkg: defaultSubscriptionName, Channel: marketplace.IntegreatlyChannel, Namespace: r.Config.GetNamespace()}, client)
	if err != nil || phase != v1alpha1.PhaseCompleted {
		return phase, err
	}

	phase, err = r.reconcileCustomResource(ctx, client)
	if err != nil || phase != v1alpha1.PhaseCompleted {
		return phase, err
	}

	r.logger.Infof("%s has reconciled successfully", r.Config.GetProductName())
	return v1alpha1.PhaseCompleted, nil
}

func (r *Reconciler) GetPreflightObject(ns string) runtime.Object {
	return nil
}

func (r *Reconciler) reconcileCustomResource(ctx context.Context, client pkgclient.Client) (v1alpha1.StatusPhase, error) {
	r.logger.Debug("reconciling mobile-developer-console custom resource")

	cr := &mdc.MobileDeveloperConsole{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: r.Config.GetNamespace(),
			Name:      resourceName,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: nexus.SchemeGroupVersion.String(),
			Kind:       nexus.NexusKind,
		},
		Spec: nexus.NexusSpec{
			NexusVolumeSize:    "10Gi",
			NexusSSL:           true,
			NexusImageTag:      "latest",
			NexusCPURequest:    1,
			NexusCPULimit:      2,
			NexusMemoryRequest: "2Gi",
			NexusMemoryLimit:   "2Gi",
		},
	}

	// attempt to create the custom resource
	if err := client.Create(ctx, cr); err != nil && !k8serr.IsAlreadyExists(err) {
		return v1alpha1.PhaseFailed, errors.Wrap(err, "failed to get or create a nexus custom resource")
	}

	// if there are no errors, the phase is complete
	return v1alpha1.PhaseCompleted, nil
}
