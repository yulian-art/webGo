# webGo 练习仓库

这是一个 Go 语言练习仓库，当前主要示例集中在 `agent/eino`，围绕 Cloudwego Eino 依次练习大模型、提示词模板、工具调用和 RAG（加载、切分、向量化、索引与检索）。仓库根目录下的其他目录是独立的学习笔记，不影响 Eino 示例的运行。

## 选择分支

| 分支 | 内容 | 用途 |
| --- | --- | --- |
| `main` | 完整仓库与修正后的文档 | 日常开发基线 |
| `practice` | 10 个示例目录，每个目录只有 `docs.md` | 只阅读文档或布置练习 |
| `answer` | 完整 `agent` 源码与文档，不含 `.env` 和 Milvus 数据卷 | 对照答案、运行示例 |

开始练习时建议使用 `practice`；需要运行代码或查看实现时切换到 `answer`：

```powershell
git clone https://github.com/yulian-art/webGo.git
cd webGo
git switch answer
```

## 环境准备

- Go `1.26.3` 或兼容版本
- 运行 `07indexer_example`、`08retriever_example` 时需要 Docker Desktop 和 Milvus
- 可访问所使用模型服务的网络环境

在 `answer` 或 `main` 分支配置模型环境变量。不要把真实密钥写入仓库：

```powershell
cd agent/eino
Copy-Item env.example .env
# 编辑 .env，填入真实的 KEY；URL 和 MODEL 按实际模型服务修改

$env:KEY = "your-api-key"
$env:URL = "https://ark.cn-beijing.volces.com/api/v3"
$env:MODEL = "doubao-seed-2-0-mini-260428"
```

`.env` 已被 Git 忽略，`env.example` 仅用于说明变量格式。

## 运行示例

每个示例都可以从自己的目录运行：

```powershell
cd agent/eino/00eino_init_example
go run .
```

推荐按以下顺序阅读和运行：

1. `00eino_init_example`：Eino 框架概览与最简单的大模型调用
2. `01chatModel_example`：使用 ChatModel 与模型交互
3. `02chatTemplate_example`：用模板生成消息上下文
4. `03tool_example`：定义并调用工具
5. `04loader_example`：加载知识库文档
6. `05transformer_example`：切分、过滤文档
7. `06embedding_example`：将文档向量化
8. `07indexer_example`：把文档向量写入 Milvus
9. `08retriever_example`：从 Milvus 检索相关内容
10. `09lambda_example`：用 Lambda 封装自定义业务逻辑

每个目录的详细说明见对应的 `docs.md`。依赖安装或更新可在示例目录执行：

```powershell
go mod download
```

## Milvus 示例

先进入 `07indexer_example`，确认 Docker Desktop 已启动，再运行 Milvus：

```powershell
cd agent/eino/07indexer_example
.\standalone.bat start
go run .
```

运行结束后可停止服务：

```powershell
.\standalone.bat stop
```

如果需要删除 Milvus 容器及其本地数据，再执行 `.\standalone.bat delete`。`08retriever_example` 依赖已写入 Milvus 的数据，通常应在索引示例之后运行。

Milvus 管理页面通常位于 `http://127.0.0.1:9091/webui/`。如使用 Attu，请先按其文档启动容器；Attu 前端在浏览器中运行，连接地址不要机械填写浏览器自己的 `localhost`，应根据 Docker 网络和主机地址配置。

## 常见问题

- 模型调用报认证或连接错误：检查当前 PowerShell 会话中的 `$env:KEY`、`$env:URL` 和 `$env:MODEL`。
- Embedding 调用失败：Embedding 服务的 URL 可能与 Chat 服务不同，按 `06embedding_example/docs.md` 配置。
- Milvus 连接失败：确认 Docker 容器已启动且端口 `19530` 可用；首次运行索引器前先启动 Milvus。
- `practice` 分支没有 Go 源码，这是刻意设计的文档练习版；请切换到 `answer` 运行代码。

## 安全提示

不要提交 `.env`、API 密钥或 Milvus `volumes` 数据。提交前可检查：

```powershell
git status
git diff --cached
```
