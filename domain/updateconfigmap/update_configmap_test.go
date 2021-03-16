package updateconfigmap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	ikb "github.com/sataga/go-kubernetes-sample/infra/kubernetes"
	v1 "k8s.io/api/core/v1"
)

func TestNewConfigMapUpdater(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	mkbc := ikb.NewMockKubernetesClient(c)

	type args struct {
		kbc ikb.KubernetesClient
	}

	tests := []struct {
		name string
		args args
		want ConfigMapUpdater
	}{
		{
			name: "test constructor",
			args: args{
				kbc: mkbc,
			},
			want: NewConfigMapUpdater(mkbc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigMapUpdater(tt.args.kbc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigMapUpdater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configMapUpdater_UpdateConfigMap(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	type args struct {
		target string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		setupMock func(*ikb.MockKubernetesClient)
	}{
		{
			name: "test tmp-logging-agent configmap not exists",
			args: args{
				target: "tmp-logging-agent",
			},
			wantErr: false,
			setupMock: func(mkbc *ikb.MockKubernetesClient) {
				mkbc.EXPECT().UpdateAPIServerURL(gomock.Any())
				mkbc.EXPECT().GetConfigMap("default", "tmp-logging-agent").Return(nil, fmt.Errorf("not found"))
				mkbc.EXPECT().CreateConfigMap("default", gomock.Any(), false)
			},
		},
		{
			name: "test tmp-logging-agent configmap exists without any no-forward-images",
			args: args{
				target: "tmp-logging-agent",
			},
			wantErr: false,
			setupMock: func(mkbc *ikb.MockKubernetesClient) {
				mkbc.EXPECT().UpdateAPIServerURL(gomock.Any())
				mkbc.EXPECT().GetConfigMap("default", "tmp-logging-agent").Return(&v1.ConfigMap{
					Data: map[string]string{
						"no-forward": "true",
					},
				}, nil)
				mkbc.EXPECT().UpdateConfigMap("default", &v1.ConfigMap{
					Data: map[string]string{
						"no-forward":        "true",
						"no-forward-images": "nghttpx-ingress-controller",
					}}, false)
			},
		},
		{
			name: "test tmp-logging-agent configmap exists without without any .Data",
			args: args{
				target: "tmp-logging-agent",
			},
			wantErr: false,
			setupMock: func(mkbc *ikb.MockKubernetesClient) {
				mkbc.EXPECT().UpdateAPIServerURL(gomock.Any())
				mkbc.EXPECT().GetConfigMap("default", "tmp-logging-agent").Return(&v1.ConfigMap{}, nil)
				mkbc.EXPECT().UpdateConfigMap("default", &v1.ConfigMap{
					Data: map[string]string{
						"no-forward-images": "nghttpx-ingress-controller",
					}}, false)
			},
		},
		{
			name: "test tmp-logging-agent configmap exists and no-forward-images exists",
			args: args{
				target: "tmp-logging-agent",
			},
			wantErr: false,
			setupMock: func(mkbc *ikb.MockKubernetesClient) {
				mkbc.EXPECT().UpdateAPIServerURL(gomock.Any())
				mkbc.EXPECT().GetConfigMap("default", "tmp-logging-agent").Return(&v1.ConfigMap{
					Data: map[string]string{
						"no-forward":        "true",
						"no-forward-images": "nghttpx-ingress-controller",
					},
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// generate mock
			mkbc := ikb.NewMockKubernetesClient(c)
			tt.setupMock(mkbc)
			cmu := &configMapUpdater{
				kbc: mkbc,
			}
			if err := cmu.UpdateConfigMap(tt.args.target, false); (err != nil) != tt.wantErr {
				t.Errorf("configMapUpdater.UpdateConfigMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
