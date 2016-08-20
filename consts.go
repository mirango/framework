package framework

type ValueType int

const (
	TYPE_STRING ValueType = iota
	TYPE_INT
	TYPE_FLOAT
	TYPE_BOOL
	TYPE_UINT
	TYPE_COMPLEX
	TYPE_STRUCT

	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"

	SCHEME_HTTP  = "http"
	SCHEME_HTTPS = "https"

	MIME_HTML = "text/html"
	MIME_XML  = "application/xml"
	MIME_JSON = "application/json"
	MIME_YAML = "application/x-yaml"

	HEADER_Accept      = "Accept"
	HEADER_ContentType = "Content-Type"
	ENCODING_GZIP      = "gzip"
	ENCODING_DEFLATE   = "deflate"
)
