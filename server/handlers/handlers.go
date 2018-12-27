package handlers

import (
	"server/database"
	"server/errors"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func GetProductByIdHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	productID := ctx.UserValue("product_id")

	result, err := database.Instance.GetProductByIdHelper(productID.(string))

	switch err {
	case nil:
		ctx.SetStatusCode(fasthttp.StatusOK) // 200
		jsonProduct, err := json.Marshal(result)
		if err != nil {
			fmt.Println("handler GetProductByIdHandler:", err)
		}
		ctx.SetBody(jsonProduct)
	case errors.ProductNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound) // 404
		errorResponce := errors.Error{
			Message: fmt.Sprintf("Can't find product with product_id: %s", productID),
		}
		jsonError, err := json.Marshal(errorResponce)
		if err != nil {
			fmt.Println("handler GetProductByIdHandler:", err)
		}
		ctx.SetBody(jsonError)
	default:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError) // 500
		ctx.SetBodyString(err.Error())
	}
}

func GetProductsAllHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	result, err := database.Instance.GetProductsAllHelper()

	switch err {
	case nil:
		ctx.SetStatusCode(fasthttp.StatusOK) // 200
		jsonProduct, err := json.Marshal(result)
		if err != nil {
			fmt.Println("handler GetProductsAllHandler:", err)
		}
		ctx.SetBody(jsonProduct)
	case errors.ProductsNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound) // 404
		errorResponce := errors.Error{
			Message: fmt.Sprintln("Can't find any products"),
		}
		jsonError, err := json.Marshal(errorResponce)
		if err != nil {
			fmt.Println("handler GetProductsAllHandler:", err)
		}
		ctx.SetBody(jsonError)
	default:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError) // 500
		ctx.SetBodyString(err.Error())
	}
}