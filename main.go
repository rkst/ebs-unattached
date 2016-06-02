package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main () {

	// Set-up the service object for AWS
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	// Setup the params for the request
	params := &ec2.DescribeVolumesInput {
		Filters: []*ec2.Filter {
			{
			// Status = Available
				Name: aws.String("status"),
				Values: []*string {
					aws.String("available"), 
					},
			},
		},
	}

	// Get instances
	vols, err := svc.DescribeVolumes(params)
	if err != nil {
		panic(err)
	}

	// Print it out
	if len(vols.Volumes) > 0 {
		for _, vl := range vols.Volumes {
			fmt.Printf("%s - %s\n", *vl.VolumeId, vl.CreateTime.Format("Jan _2 2006 15:04:05"))
		}
	}
 }