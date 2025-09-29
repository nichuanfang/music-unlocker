package watcher

// Watcher 监听器接口定义
type Watcher interface {
    Start() error
    Stop() error
}
