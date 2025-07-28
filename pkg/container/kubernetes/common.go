package kubernetes

import (
	"github.com/stacklok/toolhive/pkg/logger"
	"os"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// extra kinds
const (
	// defaultRetries is the number of times a resource discovery is retried
	defaultRetries = 10

	//defaultRetryInterval is the interval to wait before retring a resource discovery
	defaultRetryInterval = 3 * time.Second
)

var isOpenShift *bool

func DetectOpenShiftWith(config *rest.Config) (bool, error) {
	if isOpenShift == nil {
		value, ok := os.LookupEnv("OPERATOR_OPENSHIFT")
		if ok {
			//ctrl.Log.V(1).Info("Set by env-var 'OPERATOR_OPENSHIFT': " + value)
			logger.Infof("OpenShift set by env var 'OPERATOR_OPENSHIFT': " + value)
			return strings.ToLower(value) == "true", nil
		}

		var err error
		discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
		if err != nil {
			return false, err
		}

		var isOpenShiftResourcePresent bool
		for i := 0; i < defaultRetries; i++ {
			isOpenShiftResourcePresent, err = discovery.IsResourceEnabled(discoveryClient,
				schema.GroupVersionResource{
					Group:    "route.openshift.io",
					Version:  "v1",
					Resource: "routes",
				})

			if err == nil {
				break
			}

			time.Sleep(defaultRetryInterval)
		}

		if err != nil {
			return false, err
		}

		isOpenShift = &isOpenShiftResourcePresent
		if isOpenShiftResourcePresent {
			logger.Infof("OpenShift detected by route resource check.")
		}
	}
	return *isOpenShift, nil
}
