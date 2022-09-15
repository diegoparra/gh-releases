package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Repo struct {
	Owner      string
	Repository string
}

// func MakeCall() {

// }

func main() {

	repositories := []Repo{
		{
			Owner:      "aws",
			Repository: "karpenter",
		},
		{
			Owner:      "prometheus",
			Repository: "prometheus",
		},
	}

	for _, v := range repositories {

		client := github.NewClient(nil)

		// releases, _, err := client.Repositories.GetLatestRelease(context.Background(), "aws", "karpenter")
		releases, _, err := client.Repositories.GetLatestRelease(context.Background(), v.Owner, v.Repository)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Repository: ", v.Repository)
		fmt.Println("Repository Version: ", releases.GetTagName())
	}
}
