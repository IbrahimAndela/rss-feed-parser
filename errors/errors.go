package errors

//CheckErr ... checks whether an error has been raised or not and returns a boolean
func CheckErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}
