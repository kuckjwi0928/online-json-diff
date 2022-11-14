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
	mrs := a.serviceContainer.HttpService().MultiRequest([]string{req.OriginURL, req.CompareURL}, req.Method, []byte(req.BodyJSON), headers)
	for _, v := range mrs {
		if v.Err != nil {
			return nil, v.Err
		}
	}
	return &api.DiffTargetResponse{
		Left:  string(pretty.Pretty(mrs[req.OriginURL].Response)),
		Right: string(pretty.Pretty(mrs[req.CompareURL].Response)),
	}, nil
}

func (a *App) handleDiff(ctx *gin.Context, req *api.DiffTargetRequest) (*api.DiffTargetResponse, error) {
	headers := make(map[string]string)
	if req.HeaderKeys != "" {
		for _, headerKey := range strings.Split(req.HeaderKeys, ",") {
			headers[headerKey] = ctx.GetHeader(headerKey)
		}
	}
	mrs := a.serviceContainer.HttpService().MultiRequest([]string{req.OriginURL, req.CompareURL}, req.Method, []byte(req.BodyJSON), headers)
	for _, v := range mrs {
		if v.Err != nil {
			return nil, v.Err
		}
	}
	diff, err := a.serviceContainer.DiffService().Diff(mrs[req.OriginURL].Response, mrs[req.CompareURL].Response)
	if err != nil {
		return nil, err
	}
	return &api.DiffTargetResponse{
		Left:  string(pretty.Pretty(mrs[req.OriginURL].Response)),
		Right: diff,
	}, nil
}
