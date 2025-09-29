# music-unlocker

NAS 音乐助手

## Features

- **模块化**：各个功能模块职责单一，方便维护和扩展。
- **扩展性强**：方便后续增加新的音乐平台监听、新的解锁方式、新的上传目标等。
- **符合设计模式**：采用接口抽象、依赖注入等设计模式，降低耦合。
- **易于开发**：清晰分层，方便多人协作开发。

---

## 目录结构

```bash
/music-unlocker
│
├── cmd/
│ └── music-unlocker/ # 主程序入口
│ └── main.go
│
├── config/ # 配置相关
│ └── config.go # 配置加载与管理
│
├── internal/ # 内部实现，禁止外部引用
│ ├── watcher/ # 监听模块
│ │ ├── watcher.go # 监听器接口及管理
│ │ ├── netease.go # 网易云监听实现
│ │ ├── qqmusic.go # QQ 音乐监听实现
│ │ └── metube.go # MeTube 监听实现
│ │
│ ├── unlocker/ # 解锁模块
│ │ ├── unlocker.go # 解锁器接口定义
│ │ ├── umunlock.go # UM 解锁实现
│ │ └── dummy.go # 测试用解锁实现（空操作）
│ │
│ ├── cleaner/ # 清理模块（删除加密文件）
│ │ └── cleaner.go
│ │
│ ├── uploader/ # 上传模块
│ │ ├── uploader.go # 上传器接口定义
│ │ ├── webdav.go # WebDAV 上传实现
│ │ └── dummy.go # 测试用上传实现（空操作）
│ │
│ ├── model/ # 业务模型定义（如文件信息结构体等）
│ │ └── file.go
│ │
│ ├── utils/ # 工具函数（日志、文件操作等）
│ │ └── utils.go
│ │
│ └── app.go # 应用核心逻辑，协调各模块
│
├── scripts/ # 辅助脚本（构建、部署、解锁工具包装等）
│
├── docs/ # 文档
│
├── go.mod
└── go.sum
```

---

## 设计说明

### 1. cmd/music-unlocker/main.go

- 程序入口，负责初始化配置、依赖注入、启动应用。

### 2. config/

- 统一管理配置文件加载（支持 JSON/YAML/ENV 等），方便不同环境配置。

### 3. internal/watcher/

- 监听模块定义`Watcher`接口，不同音乐平台实现各自监听逻辑。
- 方便后续新增其他音乐平台，只需实现接口并注册。

### 4. internal/unlocker/

- 解锁模块定义`Unlocker`接口，支持多种解锁方式。
- 目前实现 UM 解锁，未来可扩展其他解锁工具。

### 5. internal/cleaner/

- 负责删除加密音乐文件，解锁失败时调用。

### 6. internal/uploader/

- 上传模块定义`Uploader`接口，支持多种上传方式。
- 目前实现 WebDAV 上传，未来可扩展 FTP、SFTP 等。

### 7. internal/model/

- 定义业务相关结构体，如音乐文件信息、状态等。

### 8. internal/utils/

- 公共工具函数，如日志封装、文件操作辅助函数。

### 9. internal/app.go

- 应用核心，负责协调监听、解锁、清理、上传流程。
- 采用依赖注入，将各模块实例传入，解耦合。
