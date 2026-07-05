package handlers

type errorResponseType string

const (
	INVALID_JSON        errorResponseType = "invalid_json"
	INVALID_JOB_TYPE    errorResponseType = "invalid_job_type"
	INVALID_JOB_ID      errorResponseType = "invalid_job_id"
	INVALID_JOB_PAYLOAD errorResponseType = "invalid_job_payload"

	JOB_NOT_FOUND errorResponseType = "job_not_found"

	JOB_CREATION_FAILED errorResponseType = "job_creation_failed"
	JOB_FETCH_FAILED    errorResponseType = "job_fetch_failed"
	JOB_UPDATE_FAILED   errorResponseType = "job_update_failed"
	JOB_DELETION_FAILED errorResponseType = "job_deletion_failed"

	INTERNAL_SERVER_ERROR errorResponseType = "internal_server_error"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}
type Error struct {
	Code    errorResponseType `json:"code"`
	Message string            `json:"message"`
}
