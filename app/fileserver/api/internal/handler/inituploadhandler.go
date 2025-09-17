package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 大文件上传初始化
func InitUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInitUploadLogic(r.Context(), svcCtx)
		resp, err := l.InitUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
