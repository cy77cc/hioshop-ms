package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件上传
func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
