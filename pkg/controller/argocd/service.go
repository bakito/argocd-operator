package argocd

import (
	"context"

	argoproj "github.com/jmckind/argocd-operator/pkg/apis/argoproj/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// newService retuns a new Service instance.
func newService(name string, component string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app.kubernetes.io/component": component,
				"app.kubernetes.io/name":      name,
				"app.kubernetes.io/part-of":   "argocd",
			},
		},
	}
}

func (r *ReconcileArgoCD) reconcileDexService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-dex-server", "dex-server")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-dex-server",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "http",
			Port:       5556,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(5556),
		}, {
			Name:       "grpc",
			Port:       5557,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(5557),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileMetricsService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-metrics", "metrics")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-application-controller",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "metrics",
			Port:       8082,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8082),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileRedisService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-redis", "redis")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-redis",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "tcp-redis",
			Port:       6379,
			TargetPort: intstr.FromInt(6379),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileRepoService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-repo-server", "repo-server")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-repo-server",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "server",
			Port:       8081,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8081),
		}, {
			Name:       "metrics",
			Port:       8084,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8084),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileServerMetricsService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-server-metrics", "server")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-server",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "metrics",
			Port:       8083,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8083),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileServerService(cr *argoproj.ArgoCD) error {
	svc := newService("argocd-server", "server")
	found := r.isObjectFound(types.NamespacedName{Namespace: cr.Namespace, Name: svc.Name}, svc)
	if found {
		// Service found, do nothing
		return nil
	}

	svc.Spec.Selector = map[string]string{
		"app.kubernetes.io/name": "argocd-server",
	}

	svc.Spec.Ports = []corev1.ServicePort{
		{
			Name:       "http",
			Port:       80,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8080),
		}, {
			Name:       "https",
			Port:       443,
			Protocol:   corev1.ProtocolTCP,
			TargetPort: intstr.FromInt(8080),
		},
	}
	return r.client.Create(context.TODO(), svc)
}

func (r *ReconcileArgoCD) reconcileServices(cr *argoproj.ArgoCD) error {
	err := r.reconcileDexService(cr)
	if err != nil {
		return err
	}

	err = r.reconcileMetricsService(cr)
	if err != nil {
		return err
	}

	err = r.reconcileRedisService(cr)
	if err != nil {
		return err
	}

	err = r.reconcileRepoService(cr)
	if err != nil {
		return err
	}

	err = r.reconcileServerMetricsService(cr)
	if err != nil {
		return err
	}

	err = r.reconcileServerService(cr)
	if err != nil {
		return err
	}

	return nil
}
