package tcmock

import (
	"testing"
	"time"

	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/tclog"
)

type Artifacts struct {
}

/////////////////////////////////////////////////

func (a *Artifacts) Publish(taskId string, runId uint, name, putURL, contentType, contentEncoding, file string) error {
	return nil
}

func (a *Artifacts) GetLatest(taskId, name, file string, timeout time.Duration, logger tclog.Logger) (sha256, contentEncoding, contentType string, err error) {
	return "", "", "", nil
}

/////////////////////////////////////////////////

func NewArtifacts(t *testing.T) *Artifacts {
	a := &Artifacts{}
	return a
}
