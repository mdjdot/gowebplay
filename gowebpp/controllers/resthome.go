package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RESTAssets 获取资源
func RESTAssets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprint(w, uid)
}
