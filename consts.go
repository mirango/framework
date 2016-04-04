package framework

const (
	TYPE_STRING = iota
	TYPE_INT
	TYPE_INT64
	TYPE_FLOAT
	TYPE_FLOAT64
	TYPE_BOOL

	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"

	SCHEME_HTTP  = "http"
	SCHEME_HTTPS = "https"

	MIME_HTML = "text/html"
	MIME_XML  = "application/xml"
	MIME_JSON = "application/json"

)
