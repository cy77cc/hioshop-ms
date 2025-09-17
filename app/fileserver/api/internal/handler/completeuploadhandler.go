package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件上传完成
func CompleteUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCompleteUploadLogic(r.Context(), svcCtx)
		resp, err := l.CompleteUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
