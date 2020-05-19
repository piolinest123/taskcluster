package tcmock

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/fileutil"
	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/tclog"
)

type Artifacts struct {
	t *testing.T
	// artifacts["<taskId>:<name>"]
	artifacts map[string]*Artifact
}

type Artifact struct {
	taskId          string
	runId           uint
	name            string
	contentType     string
	contentEncoding string
	data            []byte
}

/////////////////////////////////////////////////

func (a *Artifacts) Publish(taskId string, runId uint, name, putURL, contentType, contentEncoding, file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		a.t.Fatalf("Could not read file %v for artifact %v of task %v: %v", file, name, taskId, err)
	}
	a.artifacts[taskId+":"+name] = &Artifact{
		taskId:          taskId,
		runId:           runId,
		name:            name,
		contentType:     contentType,
		contentEncoding: contentEncoding,
		data:            b,
	}
	return nil
}

func (a *Artifacts) GetLatest(taskId, name, file string, timeout time.Duration, logger tclog.Logger) (sha256, contentEncoding, contentType string, err error) {
	artifact := a.artifacts[taskId+":"+name]
	err = ioutil.WriteFile(file, artifact.data, 0777)
	if err != nil {
		return
	}
	sha256, err = fileutil.CalculateSHA256(file)
	if err != nil {
		return
	}
	contentEncoding = artifact.contentEncoding
	contentType = artifact.contentType
	return
}

/////////////////////////////////////////////////

func NewArtifacts(t *testing.T) *Artifacts {
	return &Artifacts{
		t:         t,
		artifacts: map[string]*Artifact{},
	}
}
