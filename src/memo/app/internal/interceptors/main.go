package interceptors

import (
	"memo/app/internal/services"

	"github.com/revel/revel"
)

//Interceptor struct
type Interceptor struct {
	*revel.Controller
}

//Before method
func (i *Interceptor) Before() revel.Result {

	services.ConDB()

	return nil
}

//After method
func (i *Interceptor) After() revel.Result {
	return nil
}

//Panic method
func (i *Interceptor) Panic() revel.Result {
	return nil
}

func init() {
	revel.InterceptMethod((*Interceptor).Before, revel.BEFORE) // コントローラ実行前
	revel.InterceptMethod((*Interceptor).After, revel.AFTER)   // コントローラ実行後
	revel.InterceptMethod((*Interceptor).Panic, revel.PANIC)   // panic実行後
}