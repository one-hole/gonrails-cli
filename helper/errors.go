package helper

// PanicError - panic with error
func PanicError(err error) {
	if nil != err {
		panic(err)
	}
}
