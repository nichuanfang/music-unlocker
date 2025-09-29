package uploader

// DummyUploader 测试用上传实现（空操作）
type DummyUploader struct{}

func (d *DummyUploader) Upload(file string) error {
    // 空操作，直接返回 nil
    return nil
}
