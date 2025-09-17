package upload

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/logic/upload"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/cy77cc/hioshop_ms/common/response"
	"github.com/cy77cc/hioshop_ms/common/xcode"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 初始化上传，生成 uploadId
func FileUploadInitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upload.NewFileUploadInitLogic(r.Context(), svcCtx)
		data, err := l.FileUploadInit(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.Error(xcode.FileUploadFail, err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(data))
		}
	}
}
