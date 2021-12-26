package JsonMiddleware

import "net/http"

func Handle(f http.HandlerFunc, r http.Request) http.HandlerFunc {
	accept := r.Header.Get("Accept")
	if accept != "" && accept == "application/json" {
		return f
	}
	panic("hata")
}
