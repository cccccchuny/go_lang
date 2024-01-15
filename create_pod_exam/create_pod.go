package main

import (
    "os"
    "context"
    "flag"
    "fmt"
    "path/filepath"
    v1 "k8s.io/api/core/v1"
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

    pod := &v1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "mypod",
            Namespace: "default",
        },
        Spec: v1.PodSpec{
            Containers: []v1.Container{
                {
                    Name:  "mycontainer",
                    Image: "busybox",
                    Command: []string{
                        "sleep",
                        "3600",
                    },
                },
            },
        },
    }

    _, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
    if err != nil {
        panic(err)
    }

    fmt.Println("Pod is created successfully")
}