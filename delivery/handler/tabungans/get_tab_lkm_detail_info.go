package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/statuscode"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTabunganLkmInfo(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form (Randy's Framework implementations)
	payload := web.KodeLKMUri{}
	httpio.BindUri(&payload)

	usecase := usecase.NewTabunganUsecase()
	getUser, er := usecase.GetTabDetailInfo(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, getUser)
	}

}
