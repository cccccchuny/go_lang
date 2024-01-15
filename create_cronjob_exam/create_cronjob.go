package main

import (
    "context"
    "flag"
    "fmt"
    "os"
	batchv1 "k8s.io/api/batch/v1"
    "path/filepath"
    batchv1beta1 "k8s.io/api/batch/v1beta1"
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

    cronJob := &batchv1beta1.CronJob{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "mycronjob",
            Namespace: "test",
        },
        Spec: batchv1beta1.CronJobSpec{
            Schedule: "*/1 * * * *", // Run every minute
            JobTemplate: batchv1beta1.JobTemplateSpec{
                Spec: batchv1.JobSpec{
                    Template: v1.PodTemplateSpec{
                        Spec: v1.PodSpec{
                            Containers: []v1.Container{
                                {
                                    Name:  "mycontainer",
                                    Image: "go_test_image:1.0",
                                    Command: []string{
										"echo",
										"test",
                                    },
                                },
                            },
                            RestartPolicy: "OnFailure",
                        },
                    },
                },
            },
        },
    }

    _, err := clientset.BatchV1beta1().CronJobs("test").Create(context.TODO(), cronJob, metav1.CreateOptions{})
    if err != nil {
        panic(err)
    }

    fmt.Println("CronJob is created successfully") 
}
