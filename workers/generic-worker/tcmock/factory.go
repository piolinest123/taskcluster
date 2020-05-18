package tcmock

import (
	"testing"

	tcclient "github.com/taskcluster/taskcluster/v29/clients/client-go"
	"github.com/taskcluster/taskcluster/v29/workers/generic-worker/tc"
)

type ServiceFactory struct {
	T *testing.T
}

func (sf *ServiceFactory) Auth(creds *tcclient.Credentials, rootURL string) tc.Auth {
	return NewAuth(sf.T)
}

func (sf *ServiceFactory) Queue(creds *tcclient.Credentials, rootURL string) tc.Queue {
	return NewQueue(sf.T)
}

func (sf *ServiceFactory) Secrets(creds *tcclient.Credentials, rootURL string) tc.Secrets {
	return NewSecrets(sf.T)
}

func (sf *ServiceFactory) PurgeCache(creds *tcclient.Credentials, rootURL string) tc.PurgeCache {
	return NewPurgeCache(sf.T)
}

func (sf *ServiceFactory) WorkerManager(creds *tcclient.Credentials, rootURL string) tc.WorkerManager {
	return NewWorkerManager(sf.T)
}
