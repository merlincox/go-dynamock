package dynamock

func (er *errorExpectation) WillReturnError(err error) {
	er.err = err
}
