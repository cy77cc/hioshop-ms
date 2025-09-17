package s3

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/logic/s3"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 新建分段上传任务
func CreateUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := s3.NewCreateUploadLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
