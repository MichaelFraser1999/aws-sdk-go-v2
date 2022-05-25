// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The configuration used to authenticate a speaker during a session.
type AuthenticationConfiguration struct {

	// The minimum threshold needed to successfully authenticate a speaker.
	//
	// This member is required.
	AcceptanceThreshold *int32

	noSmithyDocumentSerde
}

// The authentication result produced by Voice ID, processed against the current
// session state and streamed audio of the speaker.
type AuthenticationResult struct {

	// A timestamp indicating when audio aggregation ended for this authentication
	// result.
	AudioAggregationEndedAt *time.Time

	// A timestamp indicating when audio aggregation started for this authentication
	// result.
	AudioAggregationStartedAt *time.Time

	// The unique identifier for this authentication result. Because there can be
	// multiple authentications for a given session, this field helps to identify if
	// the returned result is from a previous streaming activity or a new result. Note
	// that in absence of any new streaming activity, AcceptanceThreshold changes, or
	// SpeakerId changes, Voice ID always returns cached Authentication Result for this
	// API.
	AuthenticationResultId *string

	// The AuthenticationConfiguration used to generate this authentication result.
	Configuration *AuthenticationConfiguration

	// The client-provided identifier for the speaker whose authentication result is
	// produced. Only present if a SpeakerId is provided for the session.
	CustomerSpeakerId *string

	// The authentication decision produced by Voice ID, processed against the current
	// session state and streamed audio of the speaker.
	Decision AuthenticationDecision

	// The service-generated identifier for the speaker whose authentication result is
	// produced.
	GeneratedSpeakerId *string

	// The authentication score for the speaker whose authentication result is
	// produced. This value is only present if the authentication decision is either
	// ACCEPT or REJECT.
	Score *int32

	noSmithyDocumentSerde
}

// Contains all the information about a domain.
type Domain struct {

	// The Amazon Resource Name (ARN) for the domain.
	Arn *string

	// The timestamp at which the domain is created.
	CreatedAt *time.Time

	// The client-provided description of the domain.
	Description *string

	// The service-generated identifier for the domain.
	DomainId *string

	// The current status of the domain.
	DomainStatus DomainStatus

	// The client-provided name for the domain.
	Name *string

	// The server-side encryption configuration containing the KMS Key Identifier you
	// want Voice ID to use to encrypt your data.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// The timestamp showing the domain's last update.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Contains a summary of information about a domain.
type DomainSummary struct {

	// The Amazon Resource Name (ARN) for the domain.
	Arn *string

	// The timestamp showing when the domain is created.
	CreatedAt *time.Time

	// The client-provided description of the domain.
	Description *string

	// The service-generated identifier for the domain.
	DomainId *string

	// The current status of the domain.
	DomainStatus DomainStatus

	// The client-provided name for the domain.
	Name *string

	// The server-side encryption configuration containing the KMS Key Identifier you
	// want Voice ID to use to encrypt your data..
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// The timestamp showing the domain's last update.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Contains configurations defining enrollment behavior for the batch job.
type EnrollmentConfig struct {

	// The action to take when the specified speaker is already enrolled in the
	// specified domain. The default value is SKIP, which skips the enrollment for the
	// existing speaker. Setting the value to OVERWRITE replaces the existing voice
	// prints and enrollment audio stored for that speaker with new data generated from
	// the latest audio.
	ExistingEnrollmentAction ExistingEnrollmentAction

	// The fraud detection configuration to use for the speaker enrollment job.
	FraudDetectionConfig *EnrollmentJobFraudDetectionConfig

	noSmithyDocumentSerde
}

// The configuration defining the action to take when a speaker is flagged by the
// fraud detection system during a batch speaker enrollment job, and the risk
// threshold to use for identification.
type EnrollmentJobFraudDetectionConfig struct {

	// The action to take when the given speaker is flagged by the fraud detection
	// system. The default value is FAIL, which fails the speaker enrollment. Changing
	// this value to IGNORE results in the speaker being enrolled even if they are
	// flagged by the fraud detection system.
	FraudDetectionAction FraudDetectionAction

	// Threshold value for determining whether the speaker is a high risk to be
	// fraudulent. If the detected risk score calculated by Voice ID is greater than or
	// equal to the threshold, the speaker is considered a fraudster.
	RiskThreshold *int32

	noSmithyDocumentSerde
}

// Contains error details for a failed batch job.
type FailureDetails struct {

	// A description of the error that caused the batch job failure.
	Message *string

	// An HTTP status code representing the nature of the error.
	StatusCode *int32

	noSmithyDocumentSerde
}

// The configuration used for performing fraud detection over a speaker during a
// session.
type FraudDetectionConfiguration struct {

	// Threshold value for determining whether the speaker is a fraudster. If the
	// detected risk score calculated by Voice ID is higher than the threshold, the
	// speaker is considered a fraudster.
	//
	// This member is required.
	RiskThreshold *int32

	noSmithyDocumentSerde
}

// The fraud detection result produced by Voice ID, processed against the current
// session state and streamed audio of the speaker.
type FraudDetectionResult struct {

	// A timestamp indicating when audio aggregation ended for this fraud detection
	// result.
	AudioAggregationEndedAt *time.Time

	// A timestamp indicating when audio aggregation started for this fraud detection
	// result.
	AudioAggregationStartedAt *time.Time

	// The FraudDetectionConfiguration used to generate this fraud detection result.
	Configuration *FraudDetectionConfiguration

	// The fraud detection decision produced by Voice ID, processed against the current
	// session state and streamed audio of the speaker.
	Decision FraudDetectionDecision

	// The unique identifier for this fraud detection result. Given there can be
	// multiple fraud detections for a given session, this field helps in identifying
	// if the returned result is from previous streaming activity or a new result. Note
	// that in the absence of any new streaming activity or risk threshold changes,
	// Voice ID always returns cached Fraud Detection result for this API.
	FraudDetectionResultId *string

	// The reason speaker was flagged by the fraud detection system. This is only be
	// populated if fraud detection Decision is HIGH_RISK, and only has one possible
	// value: KNOWN_FRAUDSTER.
	Reasons []FraudDetectionReason

	// Details about each risk analyzed for this speaker.
	RiskDetails *FraudRiskDetails

	noSmithyDocumentSerde
}

// Details regarding various fraud risk analyses performed against the current
// session state and streamed audio of the speaker.
type FraudRiskDetails struct {

	// The details resulting from 'Known Fraudster Risk' analysis of the speaker.
	//
	// This member is required.
	KnownFraudsterRisk *KnownFraudsterRisk

	noSmithyDocumentSerde
}

// Contains all the information about a fraudster.
type Fraudster struct {

	// The timestamp when Voice ID identified the fraudster.
	CreatedAt *time.Time

	// The identifier for the domain containing the fraudster.
	DomainId *string

	// The service-generated identifier for the fraudster.
	GeneratedFraudsterId *string

	noSmithyDocumentSerde
}

// Contains all the information about a fraudster registration job.
type FraudsterRegistrationJob struct {

	// A timestamp showing the creation time of the fraudster registration job.
	CreatedAt *time.Time

	// The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to
	// access customer's buckets to read the input manifest file and write the job
	// output file.
	DataAccessRoleArn *string

	// The identifier of the domain containing the fraudster registration job.
	DomainId *string

	// A timestamp showing when the fraudster registration job ended.
	EndedAt *time.Time

	// Contains details that are populated when an entire batch job fails. In cases of
	// individual registration job failures, the batch job as a whole doesn't fail; it
	// is completed with a JobStatus of COMPLETED_WITH_ERRORS. You can use the job
	// output file to identify the individual registration requests that failed.
	FailureDetails *FailureDetails

	// The input data config containing an S3 URI for the input manifest file that
	// contains the list of fraudster registration job requests.
	InputDataConfig *InputDataConfig

	// The service-generated identifier for the fraudster registration job.
	JobId *string

	// The client-provied name for the fraudster registration job.
	JobName *string

	// Shows the completed percentage of registration requests listed in the input
	// file.
	JobProgress *JobProgress

	// The current status of the fraudster registration job.
	JobStatus FraudsterRegistrationJobStatus

	// The output data config containing the S3 location where you want Voice ID to
	// write your job output file; you must also include a KMS Key ID in order to
	// encrypt the file.
	OutputDataConfig *OutputDataConfig

	// The registration config containing details such as the action to take when a
	// duplicate fraudster is detected, and the similarity threshold to use for
	// detecting a duplicate fraudster.
	RegistrationConfig *RegistrationConfig

	noSmithyDocumentSerde
}

// Contains a summary of information about a fraudster registration job.
type FraudsterRegistrationJobSummary struct {

	// A timestamp showing when the fraudster registration job is created.
	CreatedAt *time.Time

	// The identifier of the domain containing the fraudster registration job.
	DomainId *string

	// A timestamp showing when the fraudster registration job ended.
	EndedAt *time.Time

	// Contains details that are populated when an entire batch job fails. In cases of
	// individual registration job failures, the batch job as a whole doesn't fail; it
	// is completed with a JobStatus of COMPLETED_WITH_ERRORS. You can use the job
	// output file to identify the individual registration requests that failed.
	FailureDetails *FailureDetails

	// The service-generated identifier for the fraudster registration job.
	JobId *string

	// The client-provied name for the fraudster registration job.
	JobName *string

	// Shows the completed percentage of registration requests listed in the input
	// file.
	JobProgress *JobProgress

	// The current status of the fraudster registration job.
	JobStatus FraudsterRegistrationJobStatus

	noSmithyDocumentSerde
}

// The configuration containing input file information for a batch job.
type InputDataConfig struct {

	// The S3 location for the input manifest file that contains the list of individual
	// enrollment or registration job requests.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Indicates the completion progress for a batch job.
type JobProgress struct {

	// Shows the completed percentage of enrollment or registration requests listed in
	// the input file.
	PercentComplete *int32

	noSmithyDocumentSerde
}

// Contains details produced as a result of performing known fraudster risk
// analysis on a speaker.
type KnownFraudsterRisk struct {

	// The score indicating the likelihood the speaker is a known fraudster.
	//
	// This member is required.
	RiskScore *int32

	// The identifier of the fraudster that is the closest match to the speaker. If
	// there are no fraudsters registered in a given domain, or if there are no
	// fraudsters with a non-zero RiskScore, this value is null.
	GeneratedFraudsterId *string

	noSmithyDocumentSerde
}

// The configuration containing output file information for a batch job.
type OutputDataConfig struct {

	// The S3 path of the folder to which Voice ID writes the job output file, which
	// has a *.out extension. For example, if the input file name is input-file.json
	// and the output folder path is s3://output-bucket/output-folder, the full output
	// file path is s3://output-bucket/output-folder/job-Id/input-file.json.out.
	//
	// This member is required.
	S3Uri *string

	// the identifier of the KMS key you want Voice ID to use to encrypt the output
	// file of the fraudster registration job.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// The configuration definining the action to take when a duplicate fraudster is
// detected, and the similarity threshold to use for detecting a duplicate
// fraudster during a batch fraudster registration job.
type RegistrationConfig struct {

	// The action to take when a fraudster is identified as a duplicate. The default
	// action is SKIP, which skips registering the duplicate fraudster. Setting the
	// value to REGISTER_AS_NEW always registers a new fraudster into the specified
	// domain.
	DuplicateRegistrationAction DuplicateRegistrationAction

	// The minimum similarity score between the new and old fraudsters in order to
	// consider the new fraudster a duplicate.
	FraudsterSimilarityThreshold *int32

	noSmithyDocumentSerde
}

// The configuration containing information about the customer-managed KMS Key used
// for encrypting customer data.
type ServerSideEncryptionConfiguration struct {

	// The identifier of the KMS Key you want Voice ID to use to encrypt your data.
	//
	// This member is required.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// Contains all the information about a speaker.
type Speaker struct {

	// A timestamp showing when the speaker is created.
	CreatedAt *time.Time

	// The client-provided identifier for the speaker.
	CustomerSpeakerId *string

	// The identifier of the domain that contains the speaker.
	DomainId *string

	// The service-generated identifier for the speaker.
	GeneratedSpeakerId *string

	// The timestamp when the speaker was last accessed for enrollment, re-enrollment
	// or a successful authentication. This timestamp is accurate to one hour.
	LastAccessedAt *time.Time

	// The current status of the speaker.
	Status SpeakerStatus

	// A timestamp showing the speaker's last update.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Contains all the information about a speaker enrollment job.
type SpeakerEnrollmentJob struct {

	// A timestamp showing the creation of the speaker enrollment job.
	CreatedAt *time.Time

	// The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to
	// access customer's buckets to read the input manifest file and write the job
	// output file.
	DataAccessRoleArn *string

	// The identifier of the domain that contains the speaker enrollment job.
	DomainId *string

	// A timestamp showing when the speaker enrollment job ended.
	EndedAt *time.Time

	// The configuration that defines the action to take when the speaker is already
	// enrolled in Voice ID, and the FraudDetectionConfig to use.
	EnrollmentConfig *EnrollmentConfig

	// Contains details that are populated when an entire batch job fails. In cases of
	// individual registration job failures, the batch job as a whole doesn't fail; it
	// is completed with a JobStatus of COMPLETED_WITH_ERRORS. You can use the job
	// output file to identify the individual registration requests that failed.
	FailureDetails *FailureDetails

	// The input data config containing an S3 URI for the input manifest file that
	// contains the list of speaker enrollment job requests.
	InputDataConfig *InputDataConfig

	// The service-generated identifier for the speaker enrollment job.
	JobId *string

	// The client-provided name for the speaker enrollment job.
	JobName *string

	// Provides details on job progress. This field shows the completed percentage of
	// registration requests listed in the input file.
	JobProgress *JobProgress

	// The current status of the speaker enrollment job.
	JobStatus SpeakerEnrollmentJobStatus

	// The output data config containing the S3 location where Voice ID writes the job
	// output file; you must also include a KMS Key ID to encrypt the file.
	OutputDataConfig *OutputDataConfig

	noSmithyDocumentSerde
}

// Contains a summary of information about a speaker enrollment job.
type SpeakerEnrollmentJobSummary struct {

	// A timestamp showing the creation time of the speaker enrollment job.
	CreatedAt *time.Time

	// The identifier of the domain that contains the speaker enrollment job.
	DomainId *string

	// A timestamp showing when the speaker enrollment job ended.
	EndedAt *time.Time

	// Contains details that are populated when an entire batch job fails. In cases of
	// individual registration job failures, the batch job as a whole doesn't fail; it
	// is completed with a JobStatus of COMPLETED_WITH_ERRORS. You can use the job
	// output file to identify the individual registration requests that failed.
	FailureDetails *FailureDetails

	// The service-generated identifier for the speaker enrollment job.
	JobId *string

	// The client-provided name for the speaker enrollment job.
	JobName *string

	// Provides details regarding job progress. This field shows the completed
	// percentage of enrollment requests listed in the input file.
	JobProgress *JobProgress

	// The current status of the speaker enrollment job.
	JobStatus SpeakerEnrollmentJobStatus

	noSmithyDocumentSerde
}

// Contains a summary of information about a speaker.
type SpeakerSummary struct {

	// A timestamp showing the speaker's creation time.
	CreatedAt *time.Time

	// The client-provided identifier for the speaker.
	CustomerSpeakerId *string

	// The identifier of the domain that contains the speaker.
	DomainId *string

	// The service-generated identifier for the speaker.
	GeneratedSpeakerId *string

	// The timestamp when the speaker was last accessed for enrollment, re-enrollment
	// or a successful authentication. This timestamp is accurate to one hour.
	LastAccessedAt *time.Time

	// The current status of the speaker.
	Status SpeakerStatus

	// A timestamp showing the speaker's last update.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// A tag that can be assigned to a Voice ID resource.
type Tag struct {

	// The first part of a key:value pair that forms a tag associated with a given
	// resource. For example, in the tag ‘Department’:’Sales’, the key is 'Department'.
	//
	// This member is required.
	Key *string

	// The second part of a key:value pair that forms a tag associated with a given
	// resource. For example, in the tag ‘Department’:’Sales’, the value is 'Sales'.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde