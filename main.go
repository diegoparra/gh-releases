package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Repo struct {
	Owner            string
	Repository       string
	CurrentLfVersion string
}

func main() {
	repositories := []Repo{
		// {
		// 	Owner:            "kubernetes-sigs",
		// 	Repository:       "aws-load-balancer-controller",
		// 	CurrentLfVersion: "v2.4.4",
		// },
		// {
		// 	Owner:            "",
		// 	Repository:       "argocd",
		// 	CurrentLfVersion: "",
		// },
		// {
		// 	Owner:            "runatlantis",
		// 	Repository:       "atlantis",
		// 	CurrentLfVersion: "v0.25.0",
		// },
		// {
		// 	Owner:            "kubernetes-sigs",
		// 	Repository:       "aws-ebs-csi-driver",
		// 	CurrentLfVersion: "v1.13.0",
		// },
		// {
		// 	Owner:            "kubernetes-sigs",
		// 	Repository:       "external-dns",
		// 	CurrentLfVersion: "v0.13.5",
		// },
		// {
		// 	Owner:            "external-secrets",
		// 	Repository:       "external-secrets",
		// 	CurrentLfVersion: "helm-chart-0.7.2",
		// },
		// {
		// 	Owner:            "grafana",
		// 	Repository:       "grafana",
		// 	CurrentLfVersion: "v9.4.7",
		// },
		// {
		// 	Owner:            "kubernetes",
		// 	Repository:       "ingress-nginx",
		// 	CurrentLfVersion: "controller-v1.4.0",
		// },
		// {
		// 	Owner:            "aws",
		// 	Repository:       "karpenter",
		// 	CurrentLfVersion: "v0.26.6",
		// },
		// {
		// 	Owner:            "kedacore",
		// 	Repository:       "keda",
		// 	CurrentLfVersion: "v2.9.3",
		// },
		//
		// TODO: This guy is failing
		// it's failing cause the package name is: kube-prometheus-stack
		// and it usually is not in the first release page
		// The function GetLatestRelease brings only the first release page
		// {
		// 	Owner:            "prometheus-community",
		// 	Repository:       "helm-charts",
		// 	CurrentLfVersion: "v45.27.2",
		// },
		// {
		// 	Owner:            "kubecost",
		// 	Repository:       "cost-analyzer-helm-chart",
		// 	CurrentLfVersion: "v1.103.2",
		// },
		// {
		// 	Owner:            "elastic",
		// 	Repository:       "logstash",
		// 	CurrentLfVersion: "v8.10.0",
		// },
		// {
		// 	Owner:            "grafana",
		// 	Repository:       "loki",
		// 	CurrentLfVersion: "v2.8.4",
		// },
		// {
		// 	Owner:            "metabase",
		// 	Repository:       "metabase",
		// 	CurrentLfVersion: "v0.46.2",
		// },
		// {
		// 	Owner:            "kubernetes-sigs",
		// 	Repository:       "prometheus-adapter",
		// 	CurrentLfVersion: "v0.10.0",
		// },
		// {
		// 	Owner:            "prometheus-community",
		// 	Repository:       "helm-charts",
		// 	CurrentLfVersion: "prometheus-elasticsearch-exporter-4.15.1",
		// },
		// {
		// 	Owner:            "rabbitmq",
		// 	Repository:       "rabbitmq-server",
		// 	CurrentLfVersion: "v3.11.10",
		// },
	}

	for _, v := range repositories {

		client := github.NewClient(nil)

		releases, _, err := client.Repositories.GetLatestRelease(
			context.Background(),
			v.Owner,
			v.Repository,
		)
		if err != nil {
			fmt.Println(err)
		}

		if releases.GetTagName() != v.CurrentLfVersion {
			fmt.Printf("Our current version for: %s, is outdated", v.Repository)
			fmt.Println("")
			fmt.Println("Current Version: ", v.CurrentLfVersion)
			fmt.Println("Available Version: ", releases.GetTagName())
			fmt.Println("")
		}

	}
}
