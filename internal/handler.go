package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/pretty"
	"log"
	"net/http"
	"online-json-diff/api"
	"strings"
)

func baseHandler[REQ api.Request, RES any](req REQ, handlerFunc func(ctx *gin.Context, req REQ) (RES, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		err := ctx.Bind(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		res, err := handlerFunc(ctx, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			panic(err)
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (a *App) handleDiffTarget(ctx *gin.Context, req *api.DiffTargetRequest) (*api.DiffTargetResponse, error) {
	headers := make(map[string]string)
	if req.HeaderKeys != "" {
		for _, headerKey := range strings.Split(req.HeaderKeys, ",") {
			headers[headerKey] = ctx.GetHeader(headerKey)
		}
	}
	res, err := a.serviceContainer.HttpService().MultiGet([]string{req.OriginURL, req.CompareURL}, headers)
	if err != nil {
		return nil, err
	}
	bytes := res.([][]byte)
	if err != nil {
		return nil, err
	}
	return &api.DiffTargetResponse{
		Left:  string(pretty.Pretty(bytes[0])),
		Right: string(pretty.Pretty(bytes[1])),
	}, nil
}

func (a *App) handleDiff(ctx *gin.Context, req *api.DiffTargetRequest) (*api.DiffTargetResponse, error) {
	headers := make(map[string]string)
	if req.HeaderKeys != "" {
		for _, headerKey := range strings.Split(req.HeaderKeys, ",") {
			headers[headerKey] = ctx.GetHeader(headerKey)
		}
	}
	res, err := a.serviceContainer.HttpService().MultiGet([]string{req.OriginURL, req.CompareURL}, headers)
	if err != nil {
		return nil, err
	}
	bytes := res.([][]byte)
	if err != nil {
		return nil, err
	}
	diff, err := a.serviceContainer.DiffService().Diff(bytes[0], bytes[1])
	if err != nil {
		return nil, err
	}
	return &api.DiffTargetResponse{
		Left:  string(pretty.Pretty(bytes[0])),
		Right: diff,
	}, nil
}
