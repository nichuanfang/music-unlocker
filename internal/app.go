package internal

import (
    "log"
    "sync"

    "github.com/nichuanfang/music-unlocker/internal/cleaner"
    "github.com/nichuanfang/music-unlocker/internal/uploader"
    "github.com/nichuanfang/music-unlocker/internal/unlocker"
    "github.com/nichuanfang/music-unlocker/internal/watcher"
)

// App 应用核心结构体，协调各模块
type App struct {
    watchers []watcher.Watcher
    unlocker unlocker.Unlocker
    cleaner  *cleaner.Cleaner
    uploader uploader.Uploader

    wg sync.WaitGroup
}

// NewApp 创建应用实例
func NewApp(
    watchers []watcher.Watcher,
    unlocker unlocker.Unlocker,
    cleaner *cleaner.Cleaner,
    uploader uploader.Uploader,
) *App {
    return &App{
        watchers: watchers,
        unlocker: unlocker,
        cleaner:  cleaner,
        uploader: uploader,
    }
}

// Start 启动所有监听器
func (a *App) Start() error {
    for _, w := range a.watchers {
        a.wg.Add(1)
        go func(w watcher.Watcher) {
            defer a.wg.Done()
            if err := w.Start(); err != nil {
                log.Printf("Watcher start error: %v", err)
            }
        }(w)
    }
    return nil
}

// Stop 停止所有监听器
func (a *App) Stop() error {
    for _, w := range a.watchers {
        if err := w.Stop(); err != nil {
            log.Printf("Watcher stop error: %v", err)
        }
    }
    a.wg.Wait()
    return nil
}

// ProcessFile 处理文件：解锁、上传、清理
func (a *App) ProcessFile(file string) error {
    if err := a.unlocker.Unlock(file); err != nil {
        log.Printf("Unlock failed for %s: %v", file, err)
        return err
    }
    if err := a.uploader.Upload(file); err != nil {
        log.Printf("Upload failed for %s: %v", file, err)
        return err
    }
    if err := a.cleaner.Clean(file); err != nil {
        log.Printf("Clean failed for %s: %v", file, err)
        return err
    }
    return nil
}
