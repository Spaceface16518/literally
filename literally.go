package literally

//StringPtr takes a string literal and provides a pointer to it
func StringPtr(literal string) *string {
	return &literal
}

//BoolPtr takes a bool literal and provides a pointer to it
func BoolPtr(literal bool) *bool {
	return &literal
}

//IntPtr takes a int literal and provides a pointer to it
func IntPtr(literal int) *int {
	return &literal
}

//UintPtr takes a uint literal and provides a pointer to it
func UintPtr(literal uint) *uint {
	return &literal
}
