package jobs

import "time"

type JobType string
type JobStatus string
type JobPriority string

type Payload struct {
}

const (
	// Status
	JobStatusPending   JobStatus = "pending"
	JobStatusQueued    JobStatus = "queued"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusRetrying  JobStatus = "retrying"
	JobStatusCancelled JobStatus = "cancelled"

	// Priority
	JobPriorityLow      JobPriority = "low"
	JobPriorityNormal   JobPriority = "normal"
	JobPriorityHigh     JobPriority = "high"
	JobPriorityCritical JobPriority = "critical"

	// Types
	JobTypeSleep   JobType = "sleep"
	JobTypeWebhook JobType = "webhook"
	JobTypeFile    JobType = "file"
	JobTypeReport  JobType = "report"
	JobTypeEmail   JobType = "email"
	JobTypeCleanup JobType = "cleanup"
	JobTypeBackup  JobType = "backup"
)

type Job struct {
	ID   string
	Type JobType
	Payload

	Status   JobStatus
	Priority JobPriority

	Retires    int
	MaxRetires int

	CreatedAt  time.Time
	UpdatedAt  time.Time
	FinishedAt time.Time

	Error string
}
