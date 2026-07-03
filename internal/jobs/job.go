package jobs

import (
	"errors"
	"time"
)

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
	JobTypeSleep JobType = "sleep"

	JobTypeSendEmail        JobType = "send_email"
	JobTypeSendNotification JobType = "send_notification"

	JobTypeGenerateThumbnail JobType = "generate_thumbnail"
	JobTypeResizeImage       JobType = "resize_image"
	JobTypeCompressFiles     JobType = "compress_files"

	JobTypeSearchText      JobType = "search_text"
	JobTypeBackupDirectory JobType = "backup_directory"
	JobTypeExportUserData  JobType = "export_user_data"

	JobTypeCleanupStorage  JobType = "cleanup_storage"
	JobTypeGenerateInvoice JobType = "generate_invoice"
)

var (
	ErrMaxRetiresExhausted = errors.New("maximum retires reached")
)

type Job struct {
	ID   string
	Type JobType
	Payload

	Status   JobStatus
	Priority JobPriority

	Retries    int
	MaxRetries int

	CreatedAt  time.Time
	StartedAt  time.Time
	UpdatedAt  time.Time
	FinishedAt time.Time

	Error string
}
