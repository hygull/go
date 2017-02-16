 package main

 import (
         "fmt"
         "github.com/aws/aws-sdk-go/aws"
         "github.com/aws/aws-sdk-go/aws/credentials"
         "github.com/aws/aws-sdk-go/service/ec2"
 )

 func main() {

         aws_access_key_id := "AKIAJHP2ZBAQDK2OA44A"
         aws_secret_access_key := "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"

         // If you're working with temporary security credentials,
         // you can also keep the session token in AWS_SESSION_TOKEN.

         token := ""

         creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

         ec2client := ec2.New(&aws.Config{
                 Region:      "ap-south-1", // get from your AWS console, click "Properties"
                 Credentials: creds,
              
         })

         regions, err := ec2client.DescribeRegions(&ec2.DescribeRegionsInput{})

         if err != nil {
                 panic(err)
         }

         // see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
         for _, region := range regions.Regions {
                 fmt.Println("Region Name : ", *region.RegionName)
                 fmt.Println("Region Endpoint : ", *region.Endpoint)
         }

 }