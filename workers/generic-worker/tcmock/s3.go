package tcmock

type S3 struct {
}

func (s3 *S3) Publish(putURL, contentType, contentEncoding, file string) error {
	return nil
}
