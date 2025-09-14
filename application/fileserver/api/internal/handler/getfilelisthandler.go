package handler

import (
	"net/http"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/logic"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件列表（分页）
func getFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFileListLogic(r.Context(), svcCtx)
		resp, err := l.GetFileList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
