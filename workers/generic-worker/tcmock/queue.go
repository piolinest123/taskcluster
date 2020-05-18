package tcmock

import (
	"net/url"
	"testing"
	"time"

	"github.com/taskcluster/taskcluster/v29/clients/client-go/tcqueue"
)

type Queue struct {
	t     *testing.T
	tasks map[string]*tcqueue.TaskDefinitionRequest
}

/////////////////////////////////////////////////

func (queue *Queue) ClaimWork(provisionerId, workerType string, payload *tcqueue.ClaimWorkRequest) (*tcqueue.ClaimWorkResponse, error) {
	return &tcqueue.ClaimWorkResponse{}, nil
}

func (queue *Queue) CreateArtifact(taskId, runId, name string, payload *tcqueue.PostArtifactRequest) (*tcqueue.PostArtifactResponse, error) {
	return &tcqueue.PostArtifactResponse{}, nil
}

func (queue *Queue) CreateTask(taskId string, payload *tcqueue.TaskDefinitionRequest) (*tcqueue.TaskStatusResponse, error) {
	tsr := &tcqueue.TaskStatusResponse{
		Status: tcqueue.TaskStatusStructure{
			Deadline:      payload.Deadline,
			Expires:       payload.Expires,
			ProvisionerID: payload.ProvisionerID,
			RetriesLeft:   payload.Retries,
			Runs:          []tcqueue.RunInformation{},
			SchedulerID:   payload.SchedulerID,
			State:         "pending",
			TaskGroupID:   payload.TaskGroupID,
			TaskID:        taskId,
			WorkerType:    payload.WorkerType,
		},
	}
	return tsr, nil
}

func (queue *Queue) GetLatestArtifact_SignedURL(taskId, name string, duration time.Duration) (*url.URL, error) {
	return &url.URL{}, nil
}

func (queue *Queue) ListArtifacts(taskId, runId, continuationToken, limit string) (*tcqueue.ListArtifactsResponse, error) {
	return &tcqueue.ListArtifactsResponse{}, nil
}

func (queue *Queue) ReclaimTask(taskId, runId string) (*tcqueue.TaskReclaimResponse, error) {
	return &tcqueue.TaskReclaimResponse{}, nil
}

func (queue *Queue) ReportCompleted(taskId, runId string) (*tcqueue.TaskStatusResponse, error) {
	return &tcqueue.TaskStatusResponse{}, nil
}

func (queue *Queue) ReportException(taskId, runId string, payload *tcqueue.TaskExceptionRequest) (*tcqueue.TaskStatusResponse, error) {
	return &tcqueue.TaskStatusResponse{}, nil
}

func (queue *Queue) ReportFailed(taskId, runId string) (*tcqueue.TaskStatusResponse, error) {
	return &tcqueue.TaskStatusResponse{}, nil
}

func (queue *Queue) Status(taskId string) (*tcqueue.TaskStatusResponse, error) {
	return &tcqueue.TaskStatusResponse{}, nil
}

func (queue *Queue) Task(taskId string) (*tcqueue.TaskDefinitionResponse, error) {
	return &tcqueue.TaskDefinitionResponse{}, nil
}

/////////////////////////////////////////////////

func NewQueue(t *testing.T) *Queue {
	q := &Queue{
		t:     t,
		tasks: map[string]*tcqueue.TaskDefinitionRequest{},
	}
	return q
}
