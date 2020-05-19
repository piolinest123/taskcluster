package tcmock

import (
	"testing"

	tcclient "github.com/taskcluster/taskcluster/v29/clients/client-go"
	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/s3"
	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/tc"
)

type ServiceFactory struct {
	t             *testing.T
	auth          tc.Auth
	queue         tc.Queue
	secrets       tc.Secrets
	purgeCache    tc.PurgeCache
	workerManager tc.WorkerManager
	s3            s3.Publisher
}

func NewServiceFactory(t *testing.T) *ServiceFactory {
	publisher := &S3{}
	return &ServiceFactory{
		auth:          NewAuth(t),
		queue:         NewQueue(t, publisher),
		secrets:       NewSecrets(t),
		purgeCache:    NewPurgeCache(t),
		workerManager: NewWorkerManager(t),
		s3:            publisher,
	}
}

func (sf *ServiceFactory) Auth(creds *tcclient.Credentials, rootURL string) tc.Auth {
	return sf.auth
}

func (sf *ServiceFactory) Queue(creds *tcclient.Credentials, rootURL string) tc.Queue {
	return sf.queue
}

func (sf *ServiceFactory) Secrets(creds *tcclient.Credentials, rootURL string) tc.Secrets {
	return sf.secrets
}

func (sf *ServiceFactory) PurgeCache(creds *tcclient.Credentials, rootURL string) tc.PurgeCache {
	return sf.purgeCache
}

func (sf *ServiceFactory) WorkerManager(creds *tcclient.Credentials, rootURL string) tc.WorkerManager {
	return sf.workerManager
}

func (sf *ServiceFactory) S3() s3.Publisher {
	return sf.s3
}
