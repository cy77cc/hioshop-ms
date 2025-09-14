package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取文件元数据
func getFileInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFileInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetFileInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
