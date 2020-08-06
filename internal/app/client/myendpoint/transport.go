package myendpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"micro_service/internal/app/client/model"
)

func MakePingEndpoint(s model.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return model.PingReq{
			Ping: RpcResp(),
		}, nil
	}
}

// 解析参数 request
func decodePingReq(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var request model.PingReq
	return request, nil
}

// 写回页面通用方法
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// 注册路由
func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	// 路由
	r.Methods("GET").Path("/ping").Handler(httptransport.NewServer(
		endpoint,
		decodePingReq,
		encodeResponse,
	))

	return r
}
