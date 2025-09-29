# 🎵 music-unlocker

> NAS 音乐助手 — 让你的音乐解锁与管理更智能、高效

---

## ✨ 核心特性

- 🔧 **模块化设计**：职责单一，易于维护与扩展  
- 🚀 **高扩展性**：轻松支持新增音乐平台监听、解锁方式及上传目标  
- 🧩 **设计模式**：接口抽象与依赖注入，降低耦合度  
- 🤝 **协作友好**：清晰分层，支持多人协同开发  

---

## 📁 项目结构

```bash
/music-unlocker
│
├── cmd/
│   └── music-unlocker/       # 主程序入口
│       └── main.go
│
├── config/                   # 配置相关
│   └── config.go             # 配置加载与管理
│
├── internal/                 # 内部实现，禁止外部引用
│   ├── watcher/              # 监听模块
│   │   ├── watcher.go        # 监听器接口及管理
│   │   ├── netease.go        # 网易云监听实现
│   │   ├── qqmusic.go        # QQ 音乐监听实现
│   │   └── metube.go         # MeTube 监听实现
│   │
│   ├── unlocker/             # 解锁模块
│   │   ├── unlocker.go       # 解锁器接口定义
│   │   ├── umunlock.go       # UM 解锁实现
│   │   └── dummy.go          # 测试用解锁实现（空操作）
│   │
│   ├── cleaner/              # 清理模块（删除加密文件）
│   │   └── cleaner.go
│   │
│   ├── uploader/             # 上传模块
│   │   ├── uploader.go       # 上传器接口定义
│   │   ├── webdav.go         # WebDAV 上传实现
│   │   └── dummy.go          # 测试用上传实现（空操作）
│   │
│   ├── model/                # 业务模型定义（如文件信息结构体等）
│   │   └── file.go
│   │
│   ├── utils/                # 工具函数（日志、文件操作等）
│   │   └── utils.go
│   │
│   ├── app/                  # 应用核心逻辑，协调各模块
│   │   └── app.go
│
├── scripts/                  # 辅助脚本（构建、部署、解锁工具包装等）
│
├── docs/                     # 文档
│
├── go.mod
└── go.sum
```

---

## 🐳 Docker 支持

### Dockerfile

项目内置 Dockerfile，轻松构建容器镜像，快速部署。

### docker-compose.yml

提供 docker-compose 配置示例，一键启动服务及依赖。

### 快速开始

```bash
# 构建镜像
docker build -t music-unlocker .

# 运行容器
docker run -d --name music-unlocker music-unlocker

# 使用 docker-compose 启动
docker-compose up -d
```

---

## 🚧 开发计划与任务拆分

为高效推进项目开发，任务拆分如下：

1. 🎧 **监听模块开发**  
   - 网易云监听（internal/watcher/netease.go）  
   - QQ 音乐监听（internal/watcher/qqmusic.go）  
   - MeTube 监听（internal/watcher/metube.go）  

2. 🔓 **解锁模块开发**  
   - UM 解锁（internal/unlocker/umunlock.go）  
   - 测试用解锁（internal/unlocker/dummy.go）  

3. 📤 **上传模块开发**  
   - WebDAV 上传（internal/uploader/webdav.go）  
   - 测试用上传（internal/uploader/dummy.go）  

4. 🧹 **清理模块开发**  
   - 删除加密文件（internal/cleaner/cleaner.go）  

5. ⚙️ **核心应用逻辑**  
   - 应用协调与管理（internal/app.go）  

6. ⚙️ **配置与工具**  
   - 配置加载与管理（config/config.go）  
   - 公共工具函数（internal/utils/utils.go）  

7. 🚀 **主程序入口**  
   - 程序初始化与启动（cmd/music-unlocker/main.go）  

---

## 📅 里程碑计划

| 阶段   | 任务内容               | 预计完成时间 |
|--------|------------------------|--------------|
| 阶段一 | 监听模块基础实现       | 2 周         |
| 阶段二 | 解锁与上传模块开发     | 3 周         |
| 阶段三 | 清理模块与核心逻辑完善 | 1 周         |
| 阶段四 | 配置、工具及主程序完善 | 1 周         |
| 阶段五 | 测试、文档及发布       | 1 周         |

---

## 🤝 贡献指南

欢迎提交 PR 和 Issue，详细说明功能需求或问题描述。

---
