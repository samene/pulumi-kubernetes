// Copyright 2021, Pulumi Corporation.  All rights reserved.

package provider

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-kubernetes/provider/v4/pkg/clients"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeConfig is a RESTClientGetter interface implementation
type KubeConfig struct {
	restConfig   *rest.Config
	clientConfig clientcmd.ClientConfig
}

// ToDiscoveryClient implemented interface method
func (k *KubeConfig) ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	c := rest.CopyConfig(k.restConfig)

	// The more groups you have, the more discovery requests you need to make.
	// given 25 groups (our groups + a few custom resources) with one-ish version each, discovery needs to make 50 requests
	// double it just so we don't end up here again for a while.  This config is only used for discovery.
	if discBurst := os.Getenv("PULUMI_K8S_CLIENT_DISCOVERY_BURST"); discBurst != "" {
		burst, err := strconv.Atoi(discBurst)
		if err != nil {
			return nil, fmt.Errorf("invalid value specified for PULUMI_K8S_CLIENT_DISCOVERY_BURST: %w", err)
		}
		c.Burst = burst
	} else {
		c.Burst = 200
	}

	return clients.NewCachedDiscoveryClient(discovery.NewDiscoveryClientForConfigOrDie(c), "~/.kube/cache/discovery/"+k.restConfig.Host, 20*time.Minute), nil
}

// ToRESTConfig implemented interface method
func (k *KubeConfig) ToRESTConfig() (*rest.Config, error) {
	return rest.CopyConfig(k.restConfig), nil
}

// ToRESTMapper implemented interface method
func (k *KubeConfig) ToRESTMapper() (meta.RESTMapper, error) {
	discoveryClient, err := k.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
	expander := restmapper.NewShortcutExpander(mapper, discoveryClient, nil)
	return expander, nil
}

// ToRawKubeConfigLoader implemented interface method
func (k *KubeConfig) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	return k.clientConfig
}

func NewKubeConfig(config *rest.Config, clientConfig clientcmd.ClientConfig) *KubeConfig {
	return &KubeConfig{restConfig: config, clientConfig: clientConfig}
}
