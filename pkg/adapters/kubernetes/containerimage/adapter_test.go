package containerimage

import (
	"errors"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func Test_containerImageAdapter_Validate(t *testing.T) {
	type args struct {
		cfg monitor.AdapterConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "kind_empty",
			args:    args{},
			wantErr: errors.New("invalid k8sContainerImage.kind value  (valid are [Deployment StatefulSet DaemonSet])"),
		},
		{
			name: "kind_invalid",
			args: args{
				cfg: monitor.AdapterConfig{K8sContainerImage: monitor.K8sContainerImage{
					Kind: "test",
				}},
			},
			wantErr: errors.New("invalid k8sContainerImage.kind value test (valid are [Deployment StatefulSet DaemonSet])"),
		},
		{
			name: "namespace_empty",
			args: args{
				cfg: monitor.AdapterConfig{K8sContainerImage: monitor.K8sContainerImage{
					Kind:          KindDeployment,
					Namespace:     "",
					Name:          "test-name",
					ContainerName: "test-container",
				}},
			},
			wantErr: ErrNamespaceEmpty,
		},
		{
			name: "name_empty",
			args: args{
				cfg: monitor.AdapterConfig{K8sContainerImage: monitor.K8sContainerImage{
					Kind:          KindDeployment,
					Namespace:     "test-ns",
					Name:          "",
					ContainerName: "test-container",
				}},
			},
			wantErr: ErrNameEmpty,
		},
		{
			name: "valid_without_container_name",
			args: args{
				cfg: monitor.AdapterConfig{K8sContainerImage: monitor.K8sContainerImage{
					Kind:          KindDeployment,
					Namespace:     "test-ns",
					Name:          "test-ns",
					ContainerName: "",
				}},
			},
			wantErr: nil,
		},
		{
			name: "valid_wit_container_name",
			args: args{
				cfg: monitor.AdapterConfig{K8sContainerImage: monitor.K8sContainerImage{
					Kind:          KindDeployment,
					Namespace:     "test-ns",
					Name:          "test-ns",
					ContainerName: "test-container",
				}},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newContainerImageAdapter(nil, nil)
			if err := a.Validate(tt.args.cfg); (err != nil) != (tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_containerImageAdapter_Fetch(t *testing.T) {
	type args struct {
		cfg monitor.AdapterConfig
	}
	tests := []struct {
		name    string
		args    args
		objects []runtime.Object
		want    string
		wantErr error
	}{
		{
			name: "Deployment_multiple_no_container_set",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDeployment,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo:test-v1.2.3.beta23-1337",
									},
									{
										Name:  "asdf-test-1",
										Image: "test-owner/test-repo:v0.1.555.1234-alpha.49",
									},
								},
							},
						},
					},
				},
			},
			want:    "",
			wantErr: errors.New("failed to load version from k8s test-nstest-name: PodTemplate has more then 1 container but no ContainerName is specified"),
		},
		{
			name: "Deployment_multiple_container_not_found",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDeployment,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-adsf-container",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo:test-v1.2.3.beta23-1337",
									},
									{
										Name:  "asdf-test-1",
										Image: "test-owner/test-repo:v0.1.555.1234-alpha.49",
									},
								},
							},
						},
					},
				},
			},
			want:    "",
			wantErr: errors.New("podTemplate of Deployment:test-ns/test-name:test-adsf-container does not have the desired container"),
		},
		{
			name: "Deployment_fallback_to_latest",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDeployment,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo",
									},
								},
							},
						},
					},
				},
			},
			want:    "latest",
			wantErr: nil,
		},
		{
			name: "Deployment",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDeployment,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo:test-v1.2.3.beta23-1337",
									},
									{
										Name:  "asdf-test-1",
										Image: "test-owner/test-repo:v0.1.555.1234-alpha.49",
									},
								},
							},
						},
					},
				},
			},
			want:    "test-v1.2.3.beta23-1337",
			wantErr: nil,
		},
		{
			name: "Deployment_not_found",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDeployment,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: nil,
			want:    "",
			wantErr: errors.New("failed to load Deployment: deployments.apps \"test-name\" not found"),
		},
		{
			name: "StatefulSet",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindStatefulSet,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.StatefulSet{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.StatefulSetSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo:test-v1.2.3.beta23-1337",
									},
									{
										Name:  "asdf-test-1",
										Image: "test-owner/test-repo:v0.1.555.1234-alpha.49",
									},
								},
							},
						},
					},
				},
			},
			want:    "test-v1.2.3.beta23-1337",
			wantErr: nil,
		},
		{
			name: "StatefulSet_not_found",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindStatefulSet,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: nil,
			want:    "",
			wantErr: errors.New("failed to load StatefulSet: statefulsets.apps \"test-name\" not found"),
		},
		{
			name: "DaemonSet",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDaemonSet,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: []runtime.Object{
				&appsv1.DaemonSet{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-name",
					},
					Spec: appsv1.DaemonSetSpec{
						Template: v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "test-container",
										Image: "test-owner/test-repo:test-v1.2.3.beta23-1337",
									},
									{
										Name:  "asdf-test-1",
										Image: "test-owner/test-repo:v0.1.555.1234-alpha.49",
									},
								},
							},
						},
					},
				},
			},
			want:    "test-v1.2.3.beta23-1337",
			wantErr: nil,
		},
		{
			name: "DaemonSet_not_found",
			args: args{
				cfg: monitor.AdapterConfig{
					K8sContainerImage: monitor.K8sContainerImage{
						Kind:          KindDaemonSet,
						Namespace:     "test-ns",
						Name:          "test-name",
						ContainerName: "test-container",
					},
				},
			},
			objects: nil,
			want:    "",
			wantErr: errors.New("failed to load DaemonSet: daemonsets.apps \"test-name\" not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newContainerImageAdapter(logging.NewTestLogger(t), newTestClient(tt.objects...))
			got, err := a.Fetch(tt.args.cfg)
			if (err != nil) != (tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTestClient(objects ...runtime.Object) kubernetes.Interface {
	return fake.NewSimpleClientset(objects...)
}
