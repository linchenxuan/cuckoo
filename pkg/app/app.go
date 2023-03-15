package app

import (
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
	"github.com/gin-gonic/gin"
)

type ICuckooApplication interface {
	InitRequestFromBody(req any) cuckoo_error.IErrorStatus
	InitRequestFromQuery(req map[string]*string) cuckoo_error.IErrorStatus
	InitRequestFromParam(req map[string]*string) cuckoo_error.IErrorStatus
	ErrorResponse(status cuckoo_error.IErrorStatus)
	SucessResponse(data any)
	SimpleResponse(status cuckoo_error.IErrorStatus, data any)
	SetUserInfo(info valobj.UserInfo)
	GetUserInfo() *valobj.UserInfo
}

type _CuckooApplication struct {
	Ctx *gin.Context
}

func NewApplication(ctx *gin.Context) ICuckooApplication {
	return &_CuckooApplication{
		Ctx: ctx,
	}
}

func (app *_CuckooApplication) InitRequestFromBody(req any) cuckoo_error.IErrorStatus {
	err := app.Ctx.ShouldBind(req)
	if err != nil {
		return cuckoo_error.InvalidParams
	}

	return nil
}

func (app *_CuckooApplication) InitRequestFromQuery(req map[string]*string) cuckoo_error.IErrorStatus {
	for k := range req {
		*(req[k]) = app.Ctx.Query(k)
	}

	return nil
}

func (app *_CuckooApplication) InitRequestFromParam(req map[string]*string) cuckoo_error.IErrorStatus {
	for k := range req {
		*(req[k]) = app.Ctx.Param(k)
	}
	return nil
}

func (app *_CuckooApplication) ErrorResponse(status cuckoo_error.IErrorStatus) {
	app.Ctx.JSON(status.StatusCode(), gin.H{
		"code":    status.Code(),
		"message": status.Message(),
	})
}

func (app *_CuckooApplication) SucessResponse(data any) {
	status := cuckoo_error.Success
	app.Ctx.JSON(status.StatusCode(), gin.H{
		"code":    status.Code(),
		"message": status.Message(),
		"data":    data,
	})
}

func (app *_CuckooApplication) SimpleResponse(status cuckoo_error.IErrorStatus, data any) {
	if status == nil {
		app.SucessResponse(data)
	} else {
		app.ErrorResponse(status)
	}
}

func (app *_CuckooApplication) SetUserInfo(info valobj.UserInfo) {
	app.Ctx.Set("userInfo", info)
}

func (app *_CuckooApplication) GetUserInfo() *valobj.UserInfo {
	value, ok := app.Ctx.Get("userInfo")
	if !ok {
		return nil
	}

	userInfo, ok := value.(valobj.UserInfo)
	if !ok {
		return nil
	}

	return &userInfo
}
