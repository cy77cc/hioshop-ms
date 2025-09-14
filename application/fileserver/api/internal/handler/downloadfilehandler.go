package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件下载
func downloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDownloadFileLogic(r.Context(), svcCtx)
		resp, err := l.DownloadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
