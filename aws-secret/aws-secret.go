package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kuyint/secrets/secrets/utilities"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type secrets struct {
}

func (s *secrets) Run(config string) any {
	a := utilities.File{
		FilePath: config,
	}
	var b JsonConf
	a.ReadConfig(&b)
	fmt.Println(b.AwsSecrets)
	fmt.Println(b.AwsSecrets[0].Name)
	createSecret(b.AwsSecrets[0].Name, b.AwsSecrets[0].Secret)
	return b
}

// Exported
var Secrets secrets

type JsonConf struct {
	AwsSecrets []Config `json:"aws_secret"`
}

type Config struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func createSecret(name string, SecretString string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := secretsmanager.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
		Name:         &name,
		SecretString: &SecretString,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	log.Println(output)
}
