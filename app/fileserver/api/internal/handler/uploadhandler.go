package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/common/resp"
	"github.com/cy77cc/hioshop_ms/common/xcode"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件上传
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r.Context(), svcCtx)
		data, err := l.Upload(r)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, resp.Error(xcode.FileUploadFail, err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp.Success(data))
		}
	}
}
