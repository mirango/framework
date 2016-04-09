package framework

import (
	"io"
)

type Renderer interface {
	MimeType() string
	Render(Context, interface{}) ([]byte, error)
	RenderToWriter(io.Writer, Context, interface{}) error
}

type Renderable interface {
	Render(Context) ([]byte, error)
	RenderToWriter(io.Writer, Context) error
}

type Renderers []Renderer

func (r Renderers) Get(mimeType string) Renderer {
	for _, re := range r {
		if re.MimeType() == mimeType {
			return re
		}
	}
	return nil
}

func (r *Renderers) Append(nr ...Renderer) {
	if nr == nil {
		return
	}
	for _, nre := range nr {
		found := false
		for i, re := range *r {
			if re.MimeType() == nre.MimeType() {
				(*r)[i] = nre
				found = true
				break
			}
		}
		if !found {
			*r = append(*r, nre)
		}
	}
}
