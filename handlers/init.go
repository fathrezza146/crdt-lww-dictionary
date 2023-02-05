package handlers

import usecase "lww-elem-set-go/usecase"

var (
	// LWW ...
	LWW usecase.LWW
)

func init() {
	LWW = *usecase.InitLWW()
}

type GeneralResponse struct {
	Success bool        "json:'success'"
	Data    interface{} "json:'data'"
	Message string      "json:'message'"
}
