package ctrl

import (
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	user, pwd, ok := r.BasicAuth()
	if !ok {
		errStr := "no basic credentials in Authorization header"
		log.Println(errStr)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(errStr))
		return
	}

	if err := c.auth.Basic(user, pwd); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	jwt, err := c.auth.GenerateJWT(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("login failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"token\":\"%s\"}", jwt)))
	return
}
