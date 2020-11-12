package schemes

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ImagePvc         = "image-pvc"
	ImageMountPath   = "/tmp"
	DefaultDindImage = "docker:19.03.0-beta5-dind"
)

func NewDindPod(namespace, pod, container, image string, opts ...PodOption) *corev1.Pod {
	p := dindPod(namespace, pod, container, image)

	for _, opt := range opts {
		opt(p)
	}

	return p
}

type PodOption func(pod *corev1.Pod)

func WithEnv(envs map[string]string) PodOption {
	return func(pod *corev1.Pod) {
		if envs == nil {
			return
		}
		for k, v := range envs {
			pod.Spec.Containers[0].Env = append(
				pod.Spec.Containers[0].Env,
				corev1.EnvVar{
					Name:  k,
					Value: v,
				},
			)
		}
	}
}

func WithPvc(pvcName string) PodOption {
	return func(pod *corev1.Pod) {
		if len(pvcName) == 0 {
			return
		}
		pod.Spec.Volumes = append(pod.Spec.Volumes,
			corev1.Volume{
				Name: ImagePvc,
				VolumeSource: corev1.VolumeSource{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: pvcName,
					},
				},
			})

		pod.Spec.Containers[0].VolumeMounts = append(pod.Spec.Containers[0].VolumeMounts,
			corev1.VolumeMount{
				Name:      ImagePvc,
				MountPath: ImageMountPath,
			},
		)

	}
}

func WithLifeCycle(cmd []string) PodOption {
	return func(pod *corev1.Pod) {
		execCmd := "update-ca-certificates; mkdir -p /root/.docker/trust/private"

		if len(cmd) > 0 {
			cmd = append([]string{execCmd}, cmd...)
			execCmd = strings.Join(cmd, "; ")
		}

		pod.Spec.Containers[0].Lifecycle = &corev1.Lifecycle{
			PostStart: &corev1.Handler{
				Exec: &corev1.ExecAction{
					Command: []string{"/bin/sh", "-c", execCmd},
				},
			},
		}
	}
}

func WithDcjSecret(secretName string) PodOption {
	return func(pod *corev1.Pod) {
		if len(secretName) == 0 {
			return
		}

		const (
			VolName = "dockerconfigjson"
			VolPath = "/home/dockremap"
		)

		pod.Spec.Containers[0].VolumeMounts = append(pod.Spec.Containers[0].VolumeMounts,
			corev1.VolumeMount{
				Name:      VolName,
				MountPath: VolPath,
			},
		)
		pod.Spec.Volumes = append(pod.Spec.Volumes,
			corev1.Volume{
				Name: VolName,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: secretName,
					},
				},
			},
		)
	}
}

func WithCertSecret(secretName string) PodOption {
	return func(pod *corev1.Pod) {
		if len(secretName) == 0 {
			return
		}
		const (
			VolName = "cert"
			VolPath = "/usr/local/share/ca-certificates"
		)

		pod.Spec.Containers[0].VolumeMounts = append(pod.Spec.Containers[0].VolumeMounts,
			corev1.VolumeMount{
				Name:      VolName,
				MountPath: VolPath,
			},
		)
		pod.Spec.Volumes = append(pod.Spec.Volumes,
			corev1.Volume{
				Name: VolName,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: secretName,
					},
				},
			},
		)
	}
}

func dindPod(namespace, pod, container, image string) *corev1.Pod {
	label := map[string]string{}
	label["obj"] = "docker-cli"

	// set dind image name
	if len(image) == 0 {
		image = DefaultDindImage
	}

	privileged := true
	dind := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod,
			Namespace: namespace,
			Labels:    label,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "docker-cli",
					Image: image,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
				},
			},
		},
	}

	return dind
}
