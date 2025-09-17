package logic

import (
	"context"
	"io"
	"net/http"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"

	filepb "github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文件上传
func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(r *http.Request, req *types.UploadReq) (resp *types.UploadResp, err error) {
	// 限制内存用于解析表单（还会生成临时文件，视需求调整）
	const maxMemory = 32 << 20 // 32MB
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return nil, err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 建立到 RPC 文件服务的 gRPC 客户端连接（此处用通用 gRPC 客户端示例）
	// conn, _ := grpc.Dial(fileServiceAddr, grpc.WithInsecure()) // 生产请用 TLS / zrpc config
	// client := filepb.NewFileServiceClient(conn)

	ctx := r.Context()
	stream, err := l.svcCtx.FileRpc.Upload(ctx)
	if err != nil {
		return nil, err
	}

	// 读取并分块发送（建议块大小例如 4MB）
	buf := make([]byte, 4<<20)
	first := true
	for {
		n, readErr := file.Read(buf)
		if n > 0 {
			chunk := &filepb.UploadChunk{
				Data: buf[:n],
			}
			if first {
				chunk.Filename = header.Filename
				chunk.ContentType = header.Header.Get("Content-Type")
				first = false
			}
			if err := stream.Send(chunk); err != nil {
				stream.CloseSend()
				return nil, err
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			stream.CloseSend()
			return nil, readErr
		}
	}

	// 发送完毕，关闭并接收服务端返回
	rpcResp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &types.UploadResp{
		Id:   rpcResp.Id,
		Name: rpcResp.Filename,
		Size: rpcResp.Size,
		Url:  rpcResp.Url,
	}, nil
}
