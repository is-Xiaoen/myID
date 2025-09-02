# ID Card Recognition System (身份证识别系统)

基于Go + GoCV + PaddleOCR的身份证识别系统，用于学习和演示防伪技术。

## 技术栈

- **后端**: Go 1.22+ 
- **图像处理**: GoCV 0.35.0 (OpenCV 4.9.0)
- **OCR引擎**: PaddleOCR 2.7.0
- **Web框架**: Gin
- **容器化**: Docker & Docker Compose

## 快速开始

### 前置要求

- Docker Desktop已安装并运行
- 已配置Docker中国镜像加速（可选，提高下载速度）

### 构建和运行

```bash
# 构建所有服务
docker-compose build

# 启动服务
docker-compose up -d

# 查看运行状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

### 测试服务

```bash
# 健康检查
curl http://localhost:8080/api/v1/health

# 上传身份证图片进行识别
curl -X POST -F "image=@test_id.jpg" http://localhost:8080/api/v1/idcard/recognize
```

## API端点

- `GET /api/v1/health` - 健康检查
- `POST /api/v1/idcard/recognize` - 身份证识别

## 项目结构

```
myID/
├── main.go              # 主程序入口
├── handlers/            # HTTP处理器
├── services/            # 业务逻辑层
├── models/              # 数据模型
├── utils/               # 工具函数
├── config/              # 配置文件
├── tests/               # 测试文件
├── Dockerfile.app-gocv  # Go应用Docker配置
├── Dockerfile.paddleocr-cn # PaddleOCR Docker配置
└── docker-compose.yml   # Docker编排文件
```

## 停止服务

```bash
# 停止服务
docker-compose stop

# 停止并删除容器
docker-compose down

# 完全清理（包括镜像）
docker-compose down --rmi all
```

## 注意事项

- 本项目仅用于学习和研究目的
- 专注于防伪检测和安全验证技术
- 不得用于非法用途