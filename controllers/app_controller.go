/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ingressv1 "github.com/yandongxiao/kubebuilder-demo/api/v1"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ingress.example.com,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ingress.example.com,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ingress.example.com,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var app ingressv1.App
	err := r.Get(ctx, req.NamespacedName, &app)
	if err != nil {
		fmt.Println(err)
	}

	// 即使每次调谐都更新 app（保持内容不变），也不会触发重复调谐。
	//r.Update(ctx, copy)

	// 只要指定的 Foo 是固定值，这种情况也不会导致死循环。
	copy := app.DeepCopy()
	copy.Spec.Foo = "dsa"
	r.Update(ctx, copy)

	fmt.Println(app.Spec.Foo)

	// 直接更新 app.Status.Name 不会生效，需要调用 r.Status.Update 才可以。
	//if app.Status.Name != "" {
	//	fmt.Println("app.Status.Name is not empty", app.Status.Name)
	//	app.Status.Name = app.Spec.Foo
	//} else {
	//	app.Status.Name = app.Spec.Foo
	//}

	// r.Update 是更新 Spec 的
	if err := r.Status().Update(ctx, &app); err != nil {
		fmt.Println(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ingressv1.App{}).
		Complete(r)
}
