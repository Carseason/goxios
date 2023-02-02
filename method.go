package goxios

type MethodType string

func (m *MethodType) String() string {
	return string(*m)
}

const (
	GET_    MethodType = "GET"
	POST_   MethodType = "POST"
	PUT_    MethodType = "PUT"
	PATCH_  MethodType = "PATCH"
	DELETE_ MethodType = "DELETE"
)
