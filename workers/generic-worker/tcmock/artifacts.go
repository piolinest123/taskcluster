package tcmock

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
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
	if err = artifact.WriteToDisk(file); err != nil {
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

func (artifact *Artifact) WriteToDisk(file string) (err error) {
	if artifact.contentEncoding == "gzip" {
		var zr *gzip.Reader
		var f *os.File
		f, err = os.Create(file)
		if err != nil {
			return
		}
		defer func() {
			err2 := f.Close()
			if err == nil {
				err = err2
			}
		}()
		zr, err = gzip.NewReader(bytes.NewBuffer(artifact.data))
		_, err = io.Copy(f, zr)
	} else {
		err = ioutil.WriteFile(file, artifact.data, 0777)
	}
	return
}

/////////////////////////////////////////////////

func NewArtifacts(t *testing.T) *Artifacts {
	return &Artifacts{
		t:         t,
		artifacts: map[string]*Artifact{},
	}
}
