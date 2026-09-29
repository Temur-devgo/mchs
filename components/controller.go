package components

import (
	"database/sql"
	"github.com/astaxie/beego"
	"mchs-api/components/db"
	"strconv"
	"strings"
)

type Controller struct {
	beego.Controller
}

func (this *Controller) Options() {
	this.Data["json"] = "OK"
	this.ServeJSON()
}

func (this *Controller) GetLanguage() int {
	lang := this.Ctx.Input.Header("Accept-Language")
	if lang == "" {
		return 1
	}

	var languageId sql.NullInt32
	row := db.DB.QueryRow("select id from languages where system_title='" + lang + "'")
	err := row.Scan(&languageId)
	if err != nil {
	}

	if languageId.Int32 == 0 {
		return 1
	} else {
		return int(languageId.Int32)
	}
}

func (this *Controller) GetQueryParam(param string, value int) int {
	params, ok := this.Ctx.Request.URL.Query()[param]

	if !ok || len(params) < 1 {
		return value
	}

	paramRes := params[0]
	res, _ := strconv.Atoi(paramRes)

	return res
}

func (this *Controller) GetQueryParamString(param string, value string) string {
	params, ok := this.Ctx.Request.URL.Query()[param]

	if !ok || len(params) < 1 {
		return value
	}

	paramRes := params[0]

	return paramRes
}

func (this *Controller) GetFilterQueryParam(param string, value string) string {
	params := this.GetString(param)

	if params != "" {
		return strings.TrimSpace(params)
	}
	return ""
}

type PaginationData struct {
	Page    int
	PerPage int
}

func (this *Controller) GetPaginationData(pagesType int) PaginationData {
	page := this.GetQueryParam("page", 1)
	perPage := this.GetQueryParam("per-page", 40)

	if pagesType == 2 {
		perPage = this.GetQueryParam("per-page", 3)
	}

	return PaginationData{Page: page, PerPage: perPage}
}