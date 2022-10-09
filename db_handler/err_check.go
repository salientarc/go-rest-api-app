package dbhandler

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
