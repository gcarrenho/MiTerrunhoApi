package api

import "github.com/gin-gonic/gin"

type WinesAPI struct {
	stores streaming.Stores
	Router *gin.RouterGroup
	schema gojsonschema.JSONLoader
}

func NewWinesAPI(stores streaming.Stores, app *gin.RouterGroup) WinesAPI {

	winesAPI := WinesAPI{
		stores: stores,
		schema: gojsonschema.NewStringLoader(schemas.VoucherPrint()),
	}

	winesAPI := app.Group("wines")

	wines.GET("/", winesAPI.list)

	winesAPI.Router = wines

	return winesAPI
}

func (api WineAPI) list(c *gin.Context) {}
