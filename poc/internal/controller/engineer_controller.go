/*
Copyright 2023.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	log "sigs.k8s.io/controller-runtime/pkg/log"

	kiratechv1 "tutorial.statu/api/v1"
)

// EngineerReconciler reconciles a Engineer object
type EngineerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kiratech.tutorial.status,resources=engineers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kiratech.tutorial.status,resources=engineers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kiratech.tutorial.status,resources=engineers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Engineer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile\

var ()

func (r *EngineerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("reconciling Enginner resource")
	var engineer kiratechv1.Engineer
	if err := r.Get(ctx, req.NamespacedName, &engineer); err != nil {
		log.Error(err, "unable to fetch Engineer")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.

		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	var company = engineer.Spec.Company

	log.Info("Engineer company", "company", company)

	var isHappy bool = false
	if company == "Kiratech" {
		isHappy = true
	}
	engineer.Status.Happy = isHappy
	if err := r.Status().Update(ctx, &engineer); err != nil {
		log.Error(err, "unable to update Engineer happy status", "status", isHappy)
		return ctrl.Result{}, err
	}
	log.Info("Engineer's status updated", "status", isHappy)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EngineerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kiratechv1.Engineer{}).
		Complete(r)
}
