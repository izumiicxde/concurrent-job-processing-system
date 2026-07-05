package jobs

import (
	"encoding/json"
	"errors"
	"time"
)

type JobType string
type JobStatus string
type JobPriority string

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

type CreateJobRequest struct {
	Type       JobType         `json:"type"`
	Payload    json.RawMessage `json:"payload"`
	Priority   JobPriority     `json:"priority"`
	MaxRetries int             `json:"max_retries"`
}

type Job struct {
	ID      string          `json:"id"`
	Type    JobType         `json:"type"`
	Payload json.RawMessage `json:"payload"`

	Status   JobStatus   `json:"status"`
	Priority JobPriority `json:"priority"`

	Retries    int `json:"retries"`
	MaxRetries int `json:"max_retries"`

	CreatedAt  time.Time `json:"created_at"`
	StartedAt  time.Time `json:"started_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	FinishedAt time.Time `json:"finished_at"`

	Error string `json:"error"`
}
