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
│ ├── app/ # 应用核心逻辑，协调各模块
│ │ └── app.go
│ │
├── scripts/ # 辅助脚本（构建、部署、解锁工具包装等）
│
├── docs/ # 文档
│
├── go.mod
└── go.sum
```

---

## Docker 支持

### Dockerfile

项目提供 Dockerfile，方便构建容器镜像，快速部署。

### docker-compose.yml

提供 docker-compose 配置示例，支持一键启动服务及依赖。

### 使用示例

```bash
# 构建镜像
docker build -t music-unlocker .

# 运行容器
docker run -d --name music-unlocker music-unlocker

# 使用 docker-compose 启动
docker-compose up -d
```

---
