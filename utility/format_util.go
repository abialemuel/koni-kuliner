package utility

import (
	"net/http"
)

func Btoi(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

func Filter(r *http.Request, searchList []string) map[string][]string {
	f := map[string][]string{}
	for _, params := range searchList {
		if len(r.URL.Query()[params]) > 0 {
			f[params] = r.URL.Query()[params]
		}
	}
	return f
}
