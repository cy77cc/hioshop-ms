好的，我帮你整理一份完整的 **前端 → API → RPC** 的请求和响应数据结构表格，按照你目前设计的 **分片上传**
（小文件一个分片，大文件多个分片）来写。

---

## 1️⃣ 前端调用 API 层

### 1.1 初始化上传（InitUpload）

| 方法   | URL          | 请求体                                                                                             | 响应体                                                      |
|------|--------------|-------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| POST | `/file/init` | `json { "hash": "xxxx", "fileName": "abc.txt", "fileType": "text/plain", "fileSize": 1048576 }` | `json { "uploadId": "123456", "bucket": "file-bucket" }` |

* **说明**：

    * 前端计算文件 hash，并传给后端
    * 后端生成 `uploadId`，用于标识整个文件上传过程
    * 数据库记录文件信息（fileName, fileType, hash, status=上传中）

---

### 1.2 上传分片（UploadPart）

| 方法   | URL            | 请求体                                                                      | 响应体                        |
|------|----------------|--------------------------------------------------------------------------|----------------------------|
| POST | `/file/upload` | `json { "uploadId": "123456", "partNumber": 1, "partData": "<binary>" }` | `json { "etag": "etag1" }` |

* **说明**：

    * `partNumber`：第几个分片
    * `partData`：分片内容，可以用 multipart/form-data 或 base64 编码
    * API 层调用 RPC 层对应方法上传分片到 MinIO，返回 ETag

---

### 1.3 完成上传（CompleteUpload）

| 方法   | URL              | 请求体                                             | 响应体                                                               |
|------|------------------|-------------------------------------------------|-------------------------------------------------------------------|
| POST | `/file/complete` | `json { "uploadId": "123456", "hash": "xxxx" }` | `json { "fileId": 1001, "fileUrl": "http://xxx/bucket/xx/xxxx" }` |

* **说明**：

    * 合并所有分片（调用 MinIO `CompleteMultipartUpload`）
    * 更新数据库状态为已完成
    * 返回文件 ID 和访问 URL

---

## 2️⃣ API 层调用 RPC 层

### 2.1 InitUpload

```go
type InitUploadReq struct {
    Hash     string
    FileName string
    FileType string
    FileSize int64
}

type InitUploadResp struct {
    UploadId string
    Bucket   string
}
```

---

### 2.2 UploadPart

```go
type UploadPartReq struct {
    UploadId   string
    PartNumber int
    PartData   []byte
}

type UploadPartResp struct {
    ETag string
}
```

* RPC 层直接调用 MinIO `PutObjectPart(uploadId, partNumber, PartData)`

---

### 2.3 CompleteUpload

```go
type CompleteUploadReq struct {
    UploadId string
    Hash     string
}

type CompleteUploadResp struct {
    FileId  int64
    FileUrl string
}
```

* RPC 层调用 MinIO `CompleteMultipartUpload`
* 更新数据库状态，返回 fileId 和 fileUrl

---

这样整理好后，你就可以直接按照 **前端 → API → RPC → MinIO** 的流程实现，每个接口的数据格式都清楚了。

如果你愿意，我可以帮你画一个 **完整的调用流程图**，前端上传一个大文件的整个过程，含每个请求和响应数据。这样文档化后开发更直观。

你想让我画吗？
