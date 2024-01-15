package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    homeDir, _ := os.UserHomeDir()
    kubeconfig := flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    flag.Parse()

    config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

    clientset, _ := kubernetes.NewForConfig(config)

    deletePolicy := metav1.DeletePropagationForeground
    err := clientset.BatchV1beta1().CronJobs("test").Delete(context.TODO(), "mycronjob", metav1.DeleteOptions{
        PropagationPolicy: &deletePolicy,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println("CronJob is deleted successfully")
}
