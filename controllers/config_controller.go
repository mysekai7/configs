/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrastructurev1alpha1 "config/api/v1alpha1"
)

// ConfigReconciler reconciles a Config object
type ConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=infrastructure.alauda.io,resources=configs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.alauda.io,resources=configs/status,verbs=get;update;patch

var crt = `-----BEGIN CERTIFICATE-----
MIIC2DCCAcCgAwIBAgIJAP9cGh+QmGS8MA0GCSqGSIb3DQEBBQUAMBIxEDAOBgNV
BAMMB2t1YmUtY2EwHhcNMjAwMjEzMDMzMzE0WhcNNDIwMTA4MDMzMzE0WjASMRAw
DgYDVQQDDAdrdWJlLWNhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
wrnRXrOCf3p0eHjfddOPtPwJt/2Us8IdTU7ni+rYlQFaGArIux3EtSSbeukEXBOV
HXHVRbIl9kqtuhZBoREWviJtX5o+Mmd0sHqqyhfbuc9X1sEbj8a88HUDIugxA/b5
5bzFhVw1IpPeH3z9I2c2udjSpqPbc/a+GQZaZM9MelVf+zwjb/X7ELrRNFK5XLzF
8GdrywyHsfGxHHF1+2piR1R05diRXA0gma54HDug1Yd9recHoOI58dplKYYXT22k
VvNeo1gbEOSQH7+FJQ9GDK146j2axHw3LoPcCzKDuxhEHS70LlPoiNqh/hEVn6uL
pYl3EMiW8o9Sf+QBTqmsTwIDAQABozEwLzAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF
4DAVBgNVHREEDjAMhwTAqGNkhwR2GO8xMA0GCSqGSIb3DQEBBQUAA4IBAQA7iSFz
YDvdwkNjpDX5y2RnfOf+Yct1TqKDbK1DrZrhKzG/uP1hlMSxTeapMHeyR0YNN95s
MVgxTjZnttZ/ZEKuHFJtSnV4rOnFXHGLjDkFrIZo0gpRkh5nzPp0m6J6j7t2WORg
SCNtuqyuB6GLo2hNd7chAn74YkiUb4nUGfO0kzEQujX5HcbGQF33u5DdGkbTtEsd
miDA9mqUa/owz3mldA4CgUrJWmMh/1XGXR3ZSLJiKy2eaLvneluLZBXXBO8tL4Kl
GcwFrW+0O2xyRoGbcueA+kiydSz6sK4DUAg6zWKnhbC5zkyryuaU9azx1WnplYHd
WYKlpw0XEznWQyZ4
-----END CERTIFICATE-----
`

func (r *ConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("config", req.NamespacedName)

	// your logic here
	cfg := infrastructurev1alpha1.Config{}
	err := r.Get(context.Background(), req.NamespacedName, &cfg)
	if err != nil {
		r.Log.Error(err, "get cfg")
	}
	cfg.Spec.Global.TlsSecret.SecretSource.Crt = []byte(crt)

	r.Log.Info(fmt.Sprintf("cfg: %+v", cfg.Spec.Global.TlsSecret.SecretSource.Crt))

	if err := r.Client.Update(context.Background(), &cfg); err != nil {
		r.Log.Error(err, "update cfg")
	}
	r.Log.Info(fmt.Sprintf("cfg.SecretSource: %s", string(cfg.Spec.Global.TlsSecret.SecretSource.Crt)))

	return ctrl.Result{}, nil
}

func (r *ConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrastructurev1alpha1.Config{}).
		Complete(r)
}
