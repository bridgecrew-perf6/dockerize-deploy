package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

func Remap(context context.Context, domain string, newNameservers []string) {

	client := getClient(context)
	domainNameservers := getDomainNameservers(client, context, domain)
	
	match := nameserversMatch(domainNameservers, newNameservers)
	if match {
		fmt.Println("nameservers match :)")
		return
	}

	updateResult := updateDomainNameservers(client, context, domain, newNameservers)
	fmt.Printf("complete (operationId: %s)\n", *updateResult.OperationId)
}

func updateDomainNameservers(client *route53domains.Client, context context.Context, domain string, newNameservers []string) *route53domains.UpdateDomainNameserversOutput {
	updateResult, err := client.UpdateDomainNameservers(
		context,
		&route53domains.UpdateDomainNameserversInput{
			DomainName: &domain,
			Nameservers: []types.Nameserver{
				{Name: &newNameservers[0]},
				{Name: &newNameservers[1]},
				{Name: &newNameservers[2]},
				{Name: &newNameservers[3]},
			},
		})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return updateResult
}

func getDomainNameservers(client *route53domains.Client, context context.Context, domain string) []types.Nameserver {
	getResult, err := client.GetDomainDetail(
		context,
		&route53domains.GetDomainDetailInput{
			DomainName: aws.String(domain),
		})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return getResult.Nameservers
}

func getClient(context context.Context) *route53domains.Client {
	cfg, err := config.LoadDefaultConfig(context, config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	client := route53domains.NewFromConfig(cfg)
	return client
}

func nameserversMatch(domainNameservers []types.Nameserver, newNameservers []string) bool {
	for i := range newNameservers {
		fmt.Printf("%s <<< %s\n", *domainNameservers[i].Name, newNameservers[i])
		if *domainNameservers[i].Name != newNameservers[i] {
			return false
		}
	}
	return true
}
