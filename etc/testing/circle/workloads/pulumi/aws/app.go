package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	secret "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func DeployApp(ctx *pulumi.Context, k8sProvider *kubernetes.Provider, saRole *iam.Role, rdsInstance *rds.Instance, bucket *s3.Bucket) error {
	cfg := config.New(ctx, "")
	enterpriseKey := os.Getenv("ENT_ACT_CODE")
	if enterpriseKey == "" {
		return errors.WithStack(fmt.Errorf("need to supply env var ENT_ACT_CODE"))
	}
	awsSAkey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSAsecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	metricCreds := os.Getenv("BIGQUERY_AUTH_JSON")
	if metricCreds == "" {
		return errors.WithStack(fmt.Errorf("need to supply env var BIGQUERY_AUTH_JSON"))
	}
	jsonKey := []byte(metricCreds)
	encoded := base64.StdEncoding.EncodeToString(jsonKey)
	wpCloudFlareLoadTestAWSKeyID := os.Getenv("CF_WP_LOADTEST_AWSKEYID")
	if wpCloudFlareLoadTestAWSKeyID == "" {
		return errors.WithStack(fmt.Errorf("need to supply env var cloudflare loadtest aws access key id."))
	}
	wpCloudFlareLoadTestEndpoint := os.Getenv("CF_WP_LOADTEST_ENDPOINT_URL")
	if wpCloudFlareLoadTestEndpoint == "" {
		return errors.WithStack(fmt.Errorf("need to supply env var cloudflare loadtest endpoint."))
	}
	wpCloudFlareLoadTestSecretAccessKey := os.Getenv("CF_WP_LOADTEST_AWSACCESSKEY")
	if wpCloudFlareLoadTestSecretAccessKey == "" {
		return errors.WithStack(fmt.Errorf("need to supply env var cloudflare loadtest aws access key."))
	}
	pachdImageTag, err := cfg.Try("pachdVersion")
	if err != nil {
		pachdImageTag = "2.5.3"
	}
	helmChartVersion, err := cfg.Try("helmChartVersion")
	if err != nil {
		helmChartVersion = ""
	}
	pgBouncerMaxConnections, err := cfg.TryInt("pgBouncerMaxConnections")
	if err != nil {
		pgBouncerMaxConnections = 1000
	}
	pgBouncerDefaultPoolSize, err := cfg.TryInt("pgBouncerDefaultPoolSize")
	if err != nil {
		pgBouncerMaxConnections = 20
	}
	etcdStorageClass, err := cfg.Try("etcdStorageClass")
	if err != nil {
		etcdStorageClass = ""
	}
	etcdResourceRequestsCPU, err := cfg.TryInt("etcdResourceLimitsRequestsCPU")
	if err != nil {
		etcdResourceRequestsCPU = 4
	}
	etcdResourceRequestsMemory, err := cfg.Try("etcdResourceLimitsRequestsMemory")
	if err != nil {
		etcdResourceRequestsMemory = "4Gi"
	}
	namespace, err := corev1.NewNamespace(ctx, "test-ns", &corev1.NamespaceArgs{},
		pulumi.Provider(k8sProvider))

	if err != nil {
		return errors.WithStack(fmt.Errorf("error occurred while attempting to create test-ns: %w", err))
	}
	_, err = secret.NewSecret(ctx, "metrics-secret", &secret.SecretArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("metrics-secret"),
			Namespace: namespace.Metadata.Elem().Name(),
		},
		Data: pulumi.StringMap{
			"creds": pulumi.String(encoded),
		},
		Type: pulumi.String("Opaque"),
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return errors.WithStack(fmt.Errorf("error creating metric secret: %w", err))
	}
	_, err = secret.NewSecret(ctx, " transfer-config", &secret.SecretArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("transfer-config"),
			Namespace: namespace.Metadata.Elem().Name(),
		},
		Data: pulumi.StringMap{
			"access_key_id":     pulumi.String(wpCloudFlareLoadTestAWSKeyID),
			"endpoint_url":      pulumi.String(wpCloudFlareLoadTestEndpoint),
			"secret_access_key": pulumi.String(wpCloudFlareLoadTestSecretAccessKey),
		},
		Type: pulumi.String("Opaque"),
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return errors.WithStack(fmt.Errorf("error creating metric secret: %w", err))
	}
	values := pulumi.Map{
		"proxy": pulumi.Map{
			"enabled": pulumi.Bool(true),
			"service": pulumi.Map{
				"type": pulumi.String("LoadBalancer"),
			},
		},
		"console": pulumi.Map{
			"enabled": pulumi.Bool(false),
		},
		"pachd": pulumi.Map{
			"logLevel": pulumi.String("debug"),
			"image": pulumi.Map{
				"tag": pulumi.String(pachdImageTag),
			},
			"storage": pulumi.Map{
				"amazon": pulumi.Map{
					"bucket": bucket.Bucket,
					"region": pulumi.String("us-west-2"),
					"id":     pulumi.String(awsSAkey),
					"secret": pulumi.String(awsSAsecret),
				},
			},
			"externalService": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
			"enterpriseLicenseKey": pulumi.String(enterpriseKey),
			"oauthClientSecret":    pulumi.String("test"),
			"rootToken":            pulumi.String("test"),
			"enterpriseSecret":     pulumi.String("test"),
		},
		"deployTarget": pulumi.String("AMAZON"),
		"global": pulumi.Map{
			"postgresql": pulumi.Map{
				"postgresqlHost":                   rdsInstance.Address,
				"postgresqlUsername":               pulumi.String("postgres"),
				"postgresqlPassword":               cfg.RequireSecret("rdsPGDBPassword"),
				"postgresqlPostgresPassword":       cfg.RequireSecret("rdsPGDBPassword"),
				"identityDatabaseFullNameOverride": pulumi.String("dex"),
			},
		},
		"postgresql": pulumi.Map{
			"enabled": pulumi.Bool(false),
		},
		"pgbouncer": pulumi.Map{
			"maxConnections":  pulumi.Int(pgBouncerMaxConnections),
			"defaultPoolSize": pulumi.Int(pgBouncerDefaultPoolSize),
		},
		"etcd": pulumi.Map{
			"resources": pulumi.Map{
				"requests": pulumi.Map{
					"cpu":    pulumi.Int(etcdResourceRequestsCPU),
					"memory": pulumi.String(etcdResourceRequestsMemory),
				},
			},
			"storageClass": pulumi.String(etcdStorageClass),
		},
	}

	if helmChartVersion == "" {
		_, err = helm.NewRelease(ctx, "pach-release", &helm.ReleaseArgs{
			Namespace: namespace.Metadata.Elem().Name(),
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String("https://helm.pachyderm.com"),
			},
			Chart:   pulumi.String("pachyderm"),
			Timeout: pulumi.Int(1200),
			Values:  values,
		}, pulumi.Provider(k8sProvider))
	} else {
		_, err = helm.NewRelease(ctx, "pach-release", &helm.ReleaseArgs{
			Namespace: namespace.Metadata.Elem().Name(),
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String("https://helm.pachyderm.com"),
			},
			Chart:   pulumi.String("pachyderm"),
			Timeout: pulumi.Int(1200),
			Version: pulumi.String(helmChartVersion),
			Values:  values,
		}, pulumi.Provider(k8sProvider))
	}

	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to successfully helm install: %w", err))
	}

	return nil
}
