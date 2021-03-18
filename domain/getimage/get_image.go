package getimage

import (
	"fmt"

	ikb "github.com/sataga/go-kubernetes-sample/infra/kubernetes"
)

type ImageGetter interface {
	GetImage(target string, dryrun bool) error
}

type imageGetter struct {
	kbc ikb.KubernetesClient
}

func NewImageGetter(kbc ikb.KubernetesClient) ImageGetter {
	return &imageGetter{
		kbc: kbc,
	}
}

func (c *imageGetter) GetImage(target string, dryrun bool) error {
	switch target {
	case "nginx-ingress-controller":
		// minikubeを利用しているためIP固定になっています。
		if err := c.kbc.UpdateAPIServerURL("https://192.168.64.5:8443"); err != nil {
			return fmt.Errorf("update api server url: %s", err)
		}
		tmp := c.kbc.GetPod("kube-system")
		fmt.Println(tmp)
		return nil
	// case "tmp-logging-agent":
	// 	var (
	// 		namespace = "default"
	// 		key       = "no-forward-images"
	// 		value     = "nghttpx-ingress-controller"
	// 	)

	// 	tmpLoggingAgent := &v1.ConfigMap{
	// 		ObjectMeta: metav1.ObjectMeta{
	// 			Name:      target,
	// 			Namespace: namespace,
	// 		},
	// 	}

	// 	// minikubeを利用しているためIP固定になっています。
	// 	if err := c.kbc.UpdateAPIServerURL("https://192.168.64.5:8443"); err != nil {
	// 		return fmt.Errorf("update api server url: %s", err)
	// 	}

	// 	cm, err := c.kbc.GetConfigMap(namespace, target)
	// 	if err != nil {
	// 		if strings.Contains(err.Error(), "not found") {
	// 			fmt.Println(err)
	// 		} else {
	// 			return fmt.Errorf("failed to get ConfigMap.: %s", err)
	// 		}
	// 	}
	// 	if cm == nil {
	// 		if err := c.kbc.CreateConfigMap(namespace, tmpLoggingAgent, dryrun); err != nil {
	// 			return fmt.Errorf("failed to create ConfigMap.: %s", err)
	// 		}
	// 		fmt.Printf("%s is created.\n", target)
	// 	} else {
	// 		if _, ok := cm.Data[key]; ok {
	// 			fmt.Printf("%s exists.\n", key)
	// 		} else {
	// 			if cm.Data == nil {
	// 				cm.Data = map[string]string{}
	// 			}
	// 			cm.Data[key] = value
	// 			if err := c.kbc.GetImage(namespace, cm, dryrun); err != nil {
	// 				return fmt.Errorf("failed to update ConfigMap.: %s", err)
	// 			}
	// 			fmt.Printf("%s is updated.\n", target)
	// 		}
	// 	}
	// 	return nil
	default:
		return fmt.Errorf("unsupported configmap target: %s", target)
	}
}
