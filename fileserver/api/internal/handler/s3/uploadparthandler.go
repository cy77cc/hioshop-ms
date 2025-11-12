package s3

import (
	"net/http"

	"github.com/cy77cc/hioshop/fileserver/api/internal/logic/s3"
	"github.com/cy77cc/hioshop/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 新建分段上传任务
func UploadPartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadPartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := s3.NewUploadPartLogic(r.Context(), svcCtx)
		resp, err := l.UploadPart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
