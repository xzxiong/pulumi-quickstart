package main

import (
	"path"

	//appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	//corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	//metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		//appLabels := pulumi.StringMap{
		//	"app": pulumi.String("nginx"),
		//}
		//deployment, err := appsv1.NewDeployment(ctx, "app-dep", &appsv1.DeploymentArgs{
		//	Spec: appsv1.DeploymentSpecArgs{
		//		Selector: &metav1.LabelSelectorArgs{
		//			MatchLabels: appLabels,
		//		},
		//		Replicas: pulumi.Int(1),
		//		Template: &corev1.PodTemplateSpecArgs{
		//			Metadata: &metav1.ObjectMetaArgs{
		//				Labels: appLabels,
		//			},
		//			Spec: &corev1.PodSpecArgs{
		//				Containers: corev1.ContainerArray{
		//					corev1.ContainerArgs{
		//						Name:  pulumi.String("nginx"),
		//						Image: pulumi.String("nginx"),
		//					}},
		//			},
		//		},
		//	},
		//})
		//if err != nil {
		//	return err
		//}
		//
		// ctx.Export("name", deployment.Metadata.Name())
		ctx.Export("author", pulumi.String(cfg.Require("author")))

		funcName := cfg.Get("get-output-function")
		if funcName == "" || funcName == "stackRefGetOutput" {
			if err := stackRefGetOutput(ctx, cfg); err != nil {
				return err
			}
		} else {
			if err := stackRefGetStringOutput(ctx, cfg); err != nil {
				return err
			}
		}

		return nil
	})
}

func stackRefGetStringOutput(ctx *pulumi.Context, cfg *config.Config) error {

	ctx.Log.Info("run stackRefGetStringOutput(not-exist-key) will get error", nil)

	dstStack := "dev"
	if ctx.Stack() == "dev" {
		dstStack = "qa"
	}

	stack := path.Join(ctx.Organization(), ctx.Project(), dstStack)
	s, err := pulumi.NewStackReference(ctx, stack, nil)
	if err != nil {
		return err
	}

	output := s.GetStringOutput(pulumi.String("not-exist-key")).ApplyT(func(v string) string {
		return v
	})
	ctx.Export("output", output)
	return nil
}

func stackRefGetOutput(ctx *pulumi.Context, cfg *config.Config) error {

	ctx.Log.Info("run stackRefGetOutput(not-exist-key) can run ok", nil)

	dstStack := "dev"
	if ctx.Stack() == "dev" {
		dstStack = "qa"
	}

	stack := path.Join(ctx.Organization(), ctx.Project(), dstStack)
	s, err := pulumi.NewStackReference(ctx, stack, nil)
	if err != nil {
		return err
	}

	output := s.GetOutput(pulumi.String("not-exist-key")).ApplyT(func(v interface{}) string {
		if v == nil {
			return "not-exist"
		}
		s, ok := v.(string)
		if !ok {
			return "not-string-val"
		}
		return s
	})
	ctx.Export("output", output)
	return nil
}
