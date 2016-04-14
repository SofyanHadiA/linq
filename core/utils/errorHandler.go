package utils

func HandleWarn(err error) bool {
	if err != nil {
		Log.Warn(err.Error())
		return true
	} else {
		return false
	}
}

func HandleFatal(err error) bool {
	if err != nil {
		Log.Fatal(err.Error())
		return true
	} else {
		return false
	}
}
