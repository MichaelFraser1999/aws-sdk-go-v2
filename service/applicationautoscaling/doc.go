// Code generated by smithy-go-codegen DO NOT EDIT.

// Package applicationautoscaling provides the API client, operations, and
// parameter types for Application Auto Scaling.
//
// With Application Auto Scaling, you can configure automatic scaling for the
// following resources:
//
//     * Amazon ECS services
//
//     * Amazon EC2 Spot Fleet
// requests
//
//     * Amazon EMR clusters
//
//     * Amazon AppStream 2.0 fleets
//
//     *
// Amazon DynamoDB tables and global secondary indexes throughput capacity
//
//     *
// Amazon Aurora Replicas
//
//     * Amazon SageMaker endpoint variants
//
//     * Custom
// resources provided by your own applications or services
//
//     * Amazon Comprehend
// document classification and entity recognizer endpoints
//
//     * AWS Lambda
// function provisioned concurrency
//
//     * Amazon Keyspaces (for Apache Cassandra)
// tables
//
//     * Amazon Managed Streaming for Apache Kafka cluster storage
//
// API
// Summary The Application Auto Scaling service API includes three key sets of
// actions:
//
//     * Register and manage scalable targets - Register AWS or custom
// resources as scalable targets (a resource that Application Auto Scaling can
// scale), set minimum and maximum capacity limits, and retrieve information on
// existing scalable targets.
//
//     * Configure and manage automatic scaling -
// Define scaling policies to dynamically scale your resources in response to
// CloudWatch alarms, schedule one-time or recurring scaling actions, and retrieve
// your recent scaling activity history.
//
//     * Suspend and resume scaling -
// Temporarily suspend and later resume automatic scaling by calling the
// RegisterScalableTarget
// (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
// API action for any Application Auto Scaling scalable target. You can suspend and
// resume (individually or in combination) scale-out activities that are triggered
// by a scaling policy, scale-in activities that are triggered by a scaling policy,
// and scheduled scaling.
//
// To learn more about Application Auto Scaling, including
// information about granting IAM users required permissions for Application Auto
// Scaling actions, see the Application Auto Scaling User Guide
// (https://docs.aws.amazon.com/autoscaling/application/userguide/what-is-application-auto-scaling.html).
package applicationautoscaling