package lecture

import "time"

func Ptr[T any](v T) *T {
	return &v
}

func TimeP(t time.Time) *time.Time {
	return &t
}

func BoolP(b bool) *bool {
	return &b
}

func StringP(s string) *string {
	return &s
}
