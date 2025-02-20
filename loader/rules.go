package loader

import (
	_ "github.com/aquasecurity/defsec/rules/aws/apigateway"
	_ "github.com/aquasecurity/defsec/rules/aws/athena"
	_ "github.com/aquasecurity/defsec/rules/aws/autoscaling"
	_ "github.com/aquasecurity/defsec/rules/aws/cloudfront"
	_ "github.com/aquasecurity/defsec/rules/aws/cloudtrail"
	_ "github.com/aquasecurity/defsec/rules/aws/cloudwatch"
	_ "github.com/aquasecurity/defsec/rules/aws/codebuild"
	_ "github.com/aquasecurity/defsec/rules/aws/config"
	_ "github.com/aquasecurity/defsec/rules/aws/documentdb"
	_ "github.com/aquasecurity/defsec/rules/aws/dynamodb"
	_ "github.com/aquasecurity/defsec/rules/aws/ebs"
	_ "github.com/aquasecurity/defsec/rules/aws/ec2"
	_ "github.com/aquasecurity/defsec/rules/aws/ecr"
	_ "github.com/aquasecurity/defsec/rules/aws/ecs"
	_ "github.com/aquasecurity/defsec/rules/aws/efs"
	_ "github.com/aquasecurity/defsec/rules/aws/eks"
	_ "github.com/aquasecurity/defsec/rules/aws/elasticache"
	_ "github.com/aquasecurity/defsec/rules/aws/elasticsearch"
	_ "github.com/aquasecurity/defsec/rules/aws/elb"
	_ "github.com/aquasecurity/defsec/rules/aws/iam"
	_ "github.com/aquasecurity/defsec/rules/aws/kinesis"
	_ "github.com/aquasecurity/defsec/rules/aws/kms"
	_ "github.com/aquasecurity/defsec/rules/aws/lambda"
	_ "github.com/aquasecurity/defsec/rules/aws/mq"
	_ "github.com/aquasecurity/defsec/rules/aws/msk"
	_ "github.com/aquasecurity/defsec/rules/aws/neptune"
	_ "github.com/aquasecurity/defsec/rules/aws/rds"
	_ "github.com/aquasecurity/defsec/rules/aws/redshift"
	_ "github.com/aquasecurity/defsec/rules/aws/s3"
	_ "github.com/aquasecurity/defsec/rules/aws/sam"
	_ "github.com/aquasecurity/defsec/rules/aws/sns"
	_ "github.com/aquasecurity/defsec/rules/aws/sqs"
	_ "github.com/aquasecurity/defsec/rules/aws/ssm"
	_ "github.com/aquasecurity/defsec/rules/aws/vpc"
	_ "github.com/aquasecurity/defsec/rules/aws/workspaces"
	_ "github.com/aquasecurity/defsec/rules/azure/appservice"
	_ "github.com/aquasecurity/defsec/rules/azure/authorization"
	_ "github.com/aquasecurity/defsec/rules/azure/compute"
	_ "github.com/aquasecurity/defsec/rules/azure/container"
	_ "github.com/aquasecurity/defsec/rules/azure/database"
	_ "github.com/aquasecurity/defsec/rules/azure/datafactory"
	_ "github.com/aquasecurity/defsec/rules/azure/datalake"
	_ "github.com/aquasecurity/defsec/rules/azure/keyvault"
	_ "github.com/aquasecurity/defsec/rules/azure/monitor"
	_ "github.com/aquasecurity/defsec/rules/azure/network"
	_ "github.com/aquasecurity/defsec/rules/azure/securitycenter"
	_ "github.com/aquasecurity/defsec/rules/azure/storage"
	_ "github.com/aquasecurity/defsec/rules/azure/synapse"
	_ "github.com/aquasecurity/defsec/rules/cloudstack/compute"
	_ "github.com/aquasecurity/defsec/rules/digitalocean/compute"
	_ "github.com/aquasecurity/defsec/rules/digitalocean/spaces"
	_ "github.com/aquasecurity/defsec/rules/general/secrets"
	_ "github.com/aquasecurity/defsec/rules/github/actions"
	_ "github.com/aquasecurity/defsec/rules/github/repositories"
	_ "github.com/aquasecurity/defsec/rules/google/bigquery"
	_ "github.com/aquasecurity/defsec/rules/google/compute"
	_ "github.com/aquasecurity/defsec/rules/google/dns"
	_ "github.com/aquasecurity/defsec/rules/google/gke"
	_ "github.com/aquasecurity/defsec/rules/google/iam"
	_ "github.com/aquasecurity/defsec/rules/google/kms"
	_ "github.com/aquasecurity/defsec/rules/google/sql"
	_ "github.com/aquasecurity/defsec/rules/google/storage"
	_ "github.com/aquasecurity/defsec/rules/kubernetes/network"
	_ "github.com/aquasecurity/defsec/rules/openstack/compute"
	_ "github.com/aquasecurity/defsec/rules/openstack/networking"
	_ "github.com/aquasecurity/defsec/rules/oracle/compute"
)
