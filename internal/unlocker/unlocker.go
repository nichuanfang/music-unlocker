package unlocker

// Unlocker 解锁器接口定义
type Unlocker interface {
    Unlock(file string) error
}
