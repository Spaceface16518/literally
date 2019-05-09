package literally

func StringPtr(literal string) *string {
	return &literal
}

func BoolPtr(literal bool) *bool {
	return &literal
}

func IntPtr(literal int) *int {
	return &literal
}

func UintPtr(literal uint) *uint {
	return &literal
}
