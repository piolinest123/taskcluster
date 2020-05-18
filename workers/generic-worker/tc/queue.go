package tc

import (
	"net/url"
	"time"

	"github.com/taskcluster/taskcluster/v29/clients/client-go/tcqueue"
)

type Queue interface {
	ClaimWork(provisionerId, workerType string, payload *tcqueue.ClaimWorkRequest) (*tcqueue.ClaimWorkResponse, error)
	CreateArtifact(taskId, runId, name string, payload *tcqueue.PostArtifactRequest) (*tcqueue.PostArtifactResponse, error)
	CreateTask(taskId string, payload *tcqueue.TaskDefinitionRequest) (*tcqueue.TaskStatusResponse, error)
	GetLatestArtifact_SignedURL(taskId, name string, duration time.Duration) (*url.URL, error)
	ListArtifacts(taskId, runId, continuationToken, limit string) (*tcqueue.ListArtifactsResponse, error)
	ReclaimTask(taskId, runId string) (*tcqueue.TaskReclaimResponse, error)
	ReportCompleted(taskId, runId string) (*tcqueue.TaskStatusResponse, error)
	ReportException(taskId, runId string, payload *tcqueue.TaskExceptionRequest) (*tcqueue.TaskStatusResponse, error)
	ReportFailed(taskId, runId string) (*tcqueue.TaskStatusResponse, error)
	Status(taskId string) (*tcqueue.TaskStatusResponse, error)
	Task(taskId string) (*tcqueue.TaskDefinitionResponse, error)
}
