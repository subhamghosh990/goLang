package main

import (
	"fmt"
	"time"
	"wait"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	appLister "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type controler struct {
	clientSet    kubernetes.Interface
	depLister    appLister.DeploymentLister
	depCacheSync cache.InformerSynced
	queue        workqueue.RateLimitingInterface
}

func CreateController(_clientset kubernetes.Interface) *controler {
	depInfo := informers.NewSharedInformerFactory(_clientset, 10*time.Minute)
	depInfo.Apps().V1().Deployments()
	con := &controler{
		clientSet:    _clientset,
		depLister:    depInfo.Lister(),
		depCacheSync: depInfo.Informer().HasSynced(),
		queue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), ""),
	}

	depInfo.Informer().AddEventHandler(
		cache.ResourceEventHandler{
			OnAdd: con.OnAdd,
		},
	)
	return con
}

func (c *controler) Run(ch <-chan struct{}) {
	if !cache.WaitForCacheSync(ch, c.depCacheSync) {
		fmt.Println("error")
	}

	wait.Until(c.worker, 1*time.Second, ch)
	<-ch
}
func (c *controler) OnAdd(obj interface{}) {
	c.queue.Add(obj)
	key, _ := cache.MetaNamespaceKeyFunc(obj)
	ns, name, _ := cache.SplitMetaNamespaceKey(key)
	dep := c.depLister.Deployments(ns).Get(name)
	svc := core.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dep.Name + "-svc",
			Namespace: ns,
			Labels: map[string]string{
				"svc": "adad",
			},
		},
		Spec: core.ServiceSpec{
			Ports: []core.ServicePort{
				core.ServicePort{
					Name:       "http",
					Port:       8080,
					TargetPort: 8080,
				},
			},
		},
	}
	c.clientSet.CoreV1().Services(ns).Create(ctx, &svc, metav1.CreateOptions{})
}

func (c *controler) worker() {

}
