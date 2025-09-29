package unlocker

// DummyUnlocker 测试用解锁实现（空操作）
type DummyUnlocker struct{}

func (d *DummyUnlocker) Unlock(file string) error {
    // 空操作，直接返回 nil
    return nil
}
