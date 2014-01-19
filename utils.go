package benchmarks


type DevNullWriter struct {

}


func (DevNullWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
