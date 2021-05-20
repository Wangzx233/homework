package gay

import "net/http"

type Router struct {
handler map[string]HandleFunc
}

//type Handler interface {
//	ServeHTTP(w http.ResponseWriter,r http.Request)
//}

// Run 监听adr下的端口“:端口号”
func (router *Router)Run(adr ...string) error {
	if len(adr)==0 {
		return http.ListenAndServe(":8080",router)
	}
	return http.ListenAndServe(adr[0],router)

}

// New 创建一个实例
func New() *Router {
	return &Router{make(map[string]HandleFunc)}
}