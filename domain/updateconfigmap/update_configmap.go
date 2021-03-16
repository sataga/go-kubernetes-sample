package updateconfigmap

import (
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ikb "github.com/sataga/go-kubernetes-sample/infra/kubernetes"
)

type ConfigMapUpdater interface {
	UpdateConfigMap(target string, dryrun bool) error
}

type configMapUpdater struct {
	kbc ikb.KubernetesClient
}

func NewConfigMapUpdater(kbc ikb.KubernetesClient) ConfigMapUpdater {
	return &configMapUpdater{
		kbc: kbc,
	}
}

func (c *configMapUpdater) UpdateConfigMap(target string, dryrun bool) error {
	switch target {
	case "tmp-logging-agent":
		var (
			namespace = "default"
			key       = "no-forward-images"
			value     = "nghttpx-ingress-controller"
		)

		tmpLoggingAgent := &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      target,
				Namespace: namespace,
			},
		}

		// minikubeを利用しているためIP固定になっています。
		if err := c.kbc.UpdateAPIServerURL("https://192.168.64.5:8443"); err != nil {
			return fmt.Errorf("update api server url: %s", err)
		}

		cm, err := c.kbc.GetConfigMap(namespace, target)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				fmt.Println(err)
			} else {
				return fmt.Errorf("failed to get ConfigMap.: %s", err)
			}
		}
		if cm == nil {
			if err := c.kbc.CreateConfigMap(namespace, tmpLoggingAgent, dryrun); err != nil {
				return fmt.Errorf("failed to create ConfigMap.: %s", err)
			}
			fmt.Printf("%s is created.\n", target)
		} else {
			if _, ok := cm.Data[key]; ok {
				fmt.Printf("%s exists.\n", key)
			} else {
				if cm.Data == nil {
					cm.Data = map[string]string{}
				}
				cm.Data[key] = value
				if err := c.kbc.UpdateConfigMap(namespace, cm, dryrun); err != nil {
					return fmt.Errorf("failed to update ConfigMap.: %s", err)
				}
				fmt.Printf("%s is updated.\n", target)
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported configmap target: %s", target)
	}
}
