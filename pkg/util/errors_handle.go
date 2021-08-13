package util

func HandleError(message string, err error) {
	if err != nil {
		Error(message, err)
	}
}

func HandleFatal(message string, err error) {
	if err != nil {
		Fatal(message, err)
	}
}
