package util

// StrPointer returns a pointer of the string
func StrPointer(s string) *string {
	return &s
}

// IntPointer returns a pointer of the int
func IntPointer(i int32) *int32 {
	return &i
}

// Int64Pointer returns a pointer of the int64
func Int64Pointer(i int64) *int64 {
	return &i
}

// BoolPointer returns a point of the bool
func BoolPointer(b bool) *bool {
	return &b
}

// PointerToBool returns the bool from a pointer
func PointerToBool(flag *bool) bool {
	if flag == nil {
		return false
	}
	return *flag
}

// PointerToString returns the string from a pointer
func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// PointerToInt32 returns the int32 from a pointer
func PointerToInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
