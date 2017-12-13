package domain

type RouteConfig struct {
	Connect *Response
	Delete  *Response
	Get     *Response
	Head    *Response
	Options *Response
	Patch   *Response
	Post    *Response
	Put     *Response
	Trace   *Response
}
