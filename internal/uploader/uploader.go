package uploader

// Uploader 上传器接口定义
type Uploader interface {
    Upload(file string) error
}
