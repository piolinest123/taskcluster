package tcmock

import (
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/taskcluster/taskcluster/v29/clients/client-go/tcqueue"
)

type Queue struct {
	t     *testing.T
	tasks map[string]*tcqueue.TaskDefinitionAndStatus
}

/////////////////////////////////////////////////

func (queue *Queue) ClaimWork(provisionerId, workerType string, payload *tcqueue.ClaimWorkRequest) (*tcqueue.ClaimWorkResponse, error) {
	log.Printf("ClaimWork called - tasks map size %v", len(queue.tasks))
	maxTasks := payload.Tasks
	tasks := []tcqueue.TaskClaim{}
	for _, j := range queue.tasks {
		if j.Task.WorkerType == workerType && j.Task.ProvisionerID == provisionerId && j.Status.State == "pending" {
			j.Status.State = "running"
			j.Status.Runs = []tcqueue.RunInformation{
				{
					RunID:         0,
					ReasonCreated: "scheduled",
				},
			}
			tasks = append(
				tasks,
				tcqueue.TaskClaim{
					Task:   j.Task,
					Status: j.Status,
				},
			)
			if len(tasks) == int(maxTasks) {
				break
			}
		} else {
			log.Printf("%v / %v / %v != %v / %v / %v", j.Task.WorkerType, j.Task.ProvisionerID, j.Status.State, workerType, provisionerId, "pending")
		}
	}
	log.Printf("Returning tasks: %v", tasks)
	return &tcqueue.ClaimWorkResponse{
		Tasks: tasks,
	}, nil
}

func (queue *Queue) CreateArtifact(taskId, runId, name string, payload *tcqueue.PostArtifactRequest) (*tcqueue.PostArtifactResponse, error) {
	return &tcqueue.PostArtifactResponse{}, nil
}

func (queue *Queue) CreateTask(taskId string, payload *tcqueue.TaskDefinitionRequest) (*tcqueue.TaskStatusResponse, error) {
	log.Printf("CreateTask called - tasks map size %v", len(queue.tasks))
	queue.tasks[taskId] = &tcqueue.TaskDefinitionAndStatus{
		Status: tcqueue.TaskStatusStructure{
			TaskID: taskId,
			State:  "pending",
		},
		Task: tcqueue.TaskDefinitionResponse{
			Created:       payload.Created,
			Deadline:      payload.Deadline,
			Dependencies:  payload.Dependencies,
			Expires:       payload.Expires,
			Extra:         payload.Extra,
			Metadata:      payload.Metadata,
			Payload:       payload.Payload,
			Priority:      payload.Priority,
			ProvisionerID: payload.ProvisionerID,
			Requires:      payload.Requires,
			Retries:       payload.Retries,
			Routes:        payload.Routes,
			SchedulerID:   payload.SchedulerID,
			Scopes:        payload.Scopes,
			Tags:          payload.Tags,
			TaskGroupID:   payload.TaskGroupID,
			WorkerType:    payload.WorkerType,
		},
	}
	log.Printf("Updated tasks map size %v", len(queue.tasks))
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
	queue.tasks[taskId].Status.Runs[0].ReasonResolved = "completed"
	queue.tasks[taskId].Status.Runs[0].State = "completed"
	return queue.Status(taskId)
}

func (queue *Queue) ReportException(taskId, runId string, payload *tcqueue.TaskExceptionRequest) (*tcqueue.TaskStatusResponse, error) {
	queue.tasks[taskId].Status.Runs[0].ReasonResolved = payload.Reason
	queue.tasks[taskId].Status.Runs[0].State = "exception"
	return queue.Status(taskId)
}

func (queue *Queue) ReportFailed(taskId, runId string) (*tcqueue.TaskStatusResponse, error) {
	queue.tasks[taskId].Status.Runs[0].ReasonResolved = "failed"
	queue.tasks[taskId].Status.Runs[0].State = "failed"
	return queue.Status(taskId)
}

func (queue *Queue) Status(taskId string) (*tcqueue.TaskStatusResponse, error) {
	return &tcqueue.TaskStatusResponse{
		Status: queue.tasks[taskId].Status,
	}, nil
}

func (queue *Queue) Task(taskId string) (*tcqueue.TaskDefinitionResponse, error) {
	return &queue.tasks[taskId].Task, nil
}

/////////////////////////////////////////////////

func NewQueue(t *testing.T) *Queue {
	q := &Queue{
		t:     t,
		tasks: map[string]*tcqueue.TaskDefinitionAndStatus{},
	}
	return q
}
