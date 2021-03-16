package kubernetes

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient interface {
	GetConfigMap(namespace string, configMap string) (*v1.ConfigMap, error)
	CreateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error
	UpdateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error
	UpdateAPIServerURL(apiServerURL string) error
}

type KubeClient struct {
	ctx    context.Context
	client kubernetes.Interface
}

func NewKubernetesClient() (KubernetesClient, error) {
	client, err := newClient()
	if err != nil {
		return nil, fmt.Errorf("failed init kubeclient: %s", err)
	}
	return &KubeClient{
		ctx:    context.Background(),
		client: client,
	}, nil
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}

func (c *KubeClient) GetConfigMap(namespace string, configMap string) (*v1.ConfigMap, error) {
	cm, err := c.client.CoreV1().ConfigMaps(namespace).Get(c.ctx, configMap, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get configmap %s/%s: %s", namespace, configMap, err)
	}
	return cm, err
}

func (c *KubeClient) CreateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error {
	var (
		cm  *v1.ConfigMap
		err error
	)
	if dryrun {
		cm, err = c.client.CoreV1().ConfigMaps(namespace).Create(c.ctx, configMap, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		fmt.Printf("Updated configmap '%s'(DryRun)\n", cm.Name)
	} else {
		cm, err = c.client.CoreV1().ConfigMaps(namespace).Create(c.ctx, configMap, metav1.CreateOptions{})
		fmt.Printf("Updated configmap '%s'\n", cm.Name)
	}
	if err != nil {
		return fmt.Errorf("update configmap %s/%s: %s", namespace, configMap.Name, err)
	}
	return nil
}

func (c *KubeClient) UpdateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error {
	var (
		cm  *v1.ConfigMap
		err error
	)
	if dryrun {
		cm, err = c.client.CoreV1().ConfigMaps(namespace).Update(c.ctx, configMap, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
		fmt.Printf("Updated configmap '%s'(DryRun)\n", cm.Name)
	} else {
		cm, err = c.client.CoreV1().ConfigMaps(namespace).Update(c.ctx, configMap, metav1.UpdateOptions{})
		fmt.Printf("Updated configmap '%s'\n", cm.Name)
	}
	if err != nil {
		return fmt.Errorf("update configmap %s/%s: %s", namespace, configMap.Name, err)
	}
	return nil
}

func (c *KubeClient) UpdateAPIServerURL(apiServerURL string) error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(apiServerURL, clientcmd.RecommendedHomeFile)
	fmt.Println("Host:" + kubeConfig.Host)
	if err != nil {
		return fmt.Errorf("updating APIServerURL %s: %s", apiServerURL, err)
	}
	c.client, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return fmt.Errorf("failed to recreate kubeconfig: %s", err)
	}
	return nil
}
