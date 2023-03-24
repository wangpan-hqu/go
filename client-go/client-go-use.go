package client_go

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func client_go_user() {

	kubeconfig := flag.String("kubeconfig", "C:/Users/wjyl/Desktop/config", "absolute path")
	flag.Parse()
	// 生成config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	/*
		config.APIPath = ApiPath
		config.GroupVersion = &corev1.SchemeGroupVersion
		config.NegotiatedSerializer = scheme.Codecs
	*/
	clientset, err := clientset(config)
	if err != nil {
		panic(err)
	}
	listNodes(clientset)
	if err != nil {
		fmt.Println(err)
	}

}
func restClient(config *rest.Config) (*rest.RESTClient, error) {
	// 生成restClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// 声明空结构体
	rest := &corev1.PodList{}
	c := context.Background()
	if err = restClient.Get().Namespace("defualt").Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).Do(c).Into(rest); err != nil {
		panic(err)
	}
	for _, v := range rest.Items {
		fmt.Printf("NameSpace: %v  Name: %v  Status: %v \n", v.Namespace, v.Name, v.Status.Phase)
	}

	return restClient, err
}

func clientset(config *rest.Config) (*kubernetes.Clientset, error) {

	// 生成clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return clientSet, err
	}
	return clientSet, nil
}

func ListCm(c *kubernetes.Clientset, ns string) error {
	configMaps, err := c.CoreV1().ConfigMaps(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, cm := range configMaps.Items {
		fmt.Printf("configName: %v, configData: %v \n", cm.Name, cm.Data)
	}
	return nil
}

func listNodes(c *kubernetes.Clientset) error {
	nodeList, err := c.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, node := range nodeList.Items {
		fmt.Printf("nodeName: %v, status: %v", node.GetName(), node.GetCreationTimestamp())
	}
	return nil
}

func ListPods(c *kubernetes.Clientset, ns string) {
	pods, err := c.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, v := range pods.Items {
		fmt.Printf("namespace: %v podname: %v podstatus: %v \n", v.Namespace, v.Name, v.Status.Phase)
	}
}

func ListDeployment(c *kubernetes.Clientset, ns string) error {
	deployments, err := c.AppsV1().Deployments(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, v := range deployments.Items {
		fmt.Printf("deploymentname: %v, available: %v, ready: %v", v.GetName(), v.Status.AvailableReplicas, v.Status.ReadyReplicas)
	}
	return nil
}

func CreatePVC(c *kubernetes.Clientset, ns string) error {

	pvc, err := c.CoreV1().PersistentVolumeClaims("namespace").Create(context.Background(), &corev1.PersistentVolumeClaim{}, metav1.CreateOptions{})

	if err != nil {
		return err
	}
	fmt.Printf("deploymentname: %v, available: %v, ready: %v", pvc.GetName(), pvc.CreationTimestamp)

	return nil
}

/*
func CreatePVC1(c *kubernetes.Clientset, ns string) error {

	pvc, err :=c.StorageV1().StorageClasses().Create()

	if err != nil {
		return err
	}
	fmt.Printf("deploymentname: %v, available: %v, ready: %v", pvc.GetName(), pvc.CreationTimestamp)

	return nil
}
*/
