package v1

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	v1 "github.com/kubeflow/common/pkg/apis/common/v1"
)

func TestSetElasticPolicy(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	type args struct {
		job *PyTorchJob
	}
	type result struct {
		expectedMinReplicas *int32
		expectedMaxReplicas *int32
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "minReplicas and maxReplicas to null",
			args: args{
				job: &PyTorchJob{
					Spec: PyTorchJobSpec{
						ElasticPolicy: &ElasticPolicy{},
						PyTorchReplicaSpecs: map[v1.ReplicaType]*v1.ReplicaSpec{
							PyTorchReplicaTypeWorker: {
								Replicas: int32Ptr(1),
							},
						},
					},
				},
			},
			result: result{
				expectedMinReplicas: int32Ptr(1),
				expectedMaxReplicas: int32Ptr(1),
			},
		},
		{
			name: "minReplicas and maxReplicas to 1",
			args: args{
				job: &PyTorchJob{
					Spec: PyTorchJobSpec{
						ElasticPolicy: &ElasticPolicy{
							MaxReplicas: int32Ptr(1),
							MinReplicas: int32Ptr(1),
						},
						PyTorchReplicaSpecs: map[v1.ReplicaType]*v1.ReplicaSpec{
							PyTorchReplicaTypeWorker: {
								Replicas: int32Ptr(1),
							},
						},
					},
				},
			},
			result: result{
				expectedMinReplicas: int32Ptr(1),
				expectedMaxReplicas: int32Ptr(1),
			},
		},
		{
			name: "minReplicas and maxReplicas to 1",
			args: args{
				job: &PyTorchJob{
					Spec: PyTorchJobSpec{
						ElasticPolicy: &ElasticPolicy{
							MaxReplicas: int32Ptr(1),
							MinReplicas: int32Ptr(1),
						},
						PyTorchReplicaSpecs: map[v1.ReplicaType]*v1.ReplicaSpec{
							PyTorchReplicaTypeWorker: {
								Replicas: int32Ptr(1),
							},
						},
					},
				},
			},
			result: result{
				expectedMinReplicas: int32Ptr(1),
				expectedMaxReplicas: int32Ptr(1),
			},
		},
		{
			name: "minReplicas to null, maxRepliacs to 1",
			args: args{
				job: &PyTorchJob{
					Spec: PyTorchJobSpec{
						ElasticPolicy: &ElasticPolicy{
							MaxReplicas: int32Ptr(1),
							MinReplicas: nil,
						},
						PyTorchReplicaSpecs: map[v1.ReplicaType]*v1.ReplicaSpec{
							PyTorchReplicaTypeWorker: {
								Replicas: int32Ptr(1),
							},
						},
					},
				},
			},
			result: result{
				expectedMinReplicas: int32Ptr(1),
				expectedMaxReplicas: int32Ptr(1),
			},
		},
		{
			name: "maxRepliacs to null, minReplicas to 1",
			args: args{
				job: &PyTorchJob{
					Spec: PyTorchJobSpec{
						ElasticPolicy: &ElasticPolicy{
							MaxReplicas: nil,
							MinReplicas: int32Ptr(1),
						},
						PyTorchReplicaSpecs: map[v1.ReplicaType]*v1.ReplicaSpec{
							PyTorchReplicaTypeWorker: {
								Replicas: int32Ptr(1),
							},
						},
					},
				},
			},
			result: result{
				expectedMinReplicas: int32Ptr(1),
				expectedMaxReplicas: int32Ptr(1),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setElasticPolicy(test.args.job)
			if test.result.expectedMinReplicas != nil {
				gomega.Expect(test.args.job.Spec.ElasticPolicy.MinReplicas).
					To(gomega.Equal(test.result.expectedMinReplicas))
			} else {
				gomega.Expect(test.args.job.Spec.ElasticPolicy.MinReplicas).
					To(gomega.BeNil())
			}

			if test.result.expectedMaxReplicas != nil {
				gomega.Expect(test.args.job.Spec.ElasticPolicy.MaxReplicas).
					To(gomega.Equal(test.result.expectedMaxReplicas))
			} else {
				gomega.Expect(test.args.job.Spec.ElasticPolicy.MaxReplicas).
					To(gomega.BeNil())
			}
		})
	}
}

func int32Ptr(n int) *int32 {
	val := int32(n)
	return &val
}
