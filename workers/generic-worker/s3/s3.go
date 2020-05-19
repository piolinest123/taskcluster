package s3

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/taskcluster/httpbackoff/v3"
)

type Publisher interface {
	Publish(putURL, contentType, contentEncoding, file string) error
}

type BackoffPublisher struct {
}

func (b *BackoffPublisher) Publish(putURL, contentType, contentEncoding, file string) error {
	// perform http PUT to upload to S3...
	httpClient := &http.Client{}
	httpCall := func() (putResp *http.Response, tempError error, permError error) {
		var transferContent *os.File
		transferContent, permError = os.Open(file)
		if permError != nil {
			return
		}
		defer transferContent.Close()
		var fileInfo os.FileInfo
		fileInfo, permError = transferContent.Stat()
		if permError != nil {
			return
		}
		transferContentLength := fileInfo.Size()

		var httpRequest *http.Request
		httpRequest, permError = http.NewRequest("PUT", putURL, transferContent)
		if permError != nil {
			return
		}
		httpRequest.Header.Set("Content-Type", contentType)
		httpRequest.ContentLength = transferContentLength
		if enc := contentEncoding; enc != "" {
			httpRequest.Header.Set("Content-Encoding", enc)
		}
		requestHeaders, dumpError := httputil.DumpRequestOut(httpRequest, false)
		if dumpError != nil {
			log.Print("Could not dump request, never mind...")
		} else {
			log.Print("Request")
			log.Print(string(requestHeaders))
		}
		putResp, tempError = httpClient.Do(httpRequest)
		if tempError != nil {
			return
		}
		// bug 1394557: s3 incorrectly returns HTTP 400 for connection inactivity,
		// which can/should be retried, so explicitly handle...
		if putResp.StatusCode == 400 {
			tempError = fmt.Errorf("S3 returned status code 400 which could be an intermittent issue - see https://bugzilla.mozilla.org/show_bug.cgi?id=1394557")
		}
		return
	}
	putResp, putAttempts, err := httpbackoff.Retry(httpCall)
	log.Printf("%v put requests issued to %v", putAttempts, putURL)
	if putResp != nil {
		defer putResp.Body.Close()
		respBody, dumpError := httputil.DumpResponse(putResp, true)
		if dumpError != nil {
			log.Print("Could not dump response output, never mind...")
		} else {
			log.Print("Response")
			log.Print(string(respBody))
		}
	}
	return err
}
