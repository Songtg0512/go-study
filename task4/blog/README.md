# 个人博客系统 API 文档

## 项目简介
这是一个使用 Go + Gin + GORM 开发的个人博客系统后端，支持用户认证、文章管理和评论功能。

## 技术栈
- **Go** 1.x
- **Gin** Web框架
- **GORM** ORM库
- **SQLite** 数据库
- **JWT** 用户认证
- **bcrypt** 密码加密
- **logrus** 日志记录

## 项目结构
```
blog/
├── config/          # 配置文件
├── controllers/     # 控制器
├── middlewares/     # 中间件
├── models/          # 数据模型
├── routes/          # 路由配置
├── utils/           # 工具函数
├── main.go          # 程序入口
└── blog.db          # SQLite 数据库文件
```

## 启动项目

### 1. 安装依赖
```bash
go mod tidy
```

### 2. 编译项目
```bash
go build -o blog-server
```

### 3. 运行服务
```bash
./blog-server
```

服务默认运行在 `http://localhost:8080`

## API 接口文档

### 基础信息
- **Base URL**: `http://localhost:8080/api`
- **认证方式**: JWT Token (Header: `Authorization: Bearer <token>`)

---

### 1. 健康检查

**接口**: `GET /api/health`

**说明**: 检查服务是否正常运行

**响应示例**:
```json
{
  "status": "ok"
}
```

---

### 2. 用户认证

#### 2.1 用户注册

**接口**: `POST /api/auth/register`

**请求体**:
```json
{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}
```

**响应示例**:
```json
{
  "code": 201,
  "message": "注册成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-11-18T11:18:33Z",
      "updated_at": "2025-11-18T11:18:33Z"
    }
  }
}
```

#### 2.2 用户登录

**接口**: `POST /api/auth/login`

**请求体**:
```json
{
  "username": "testuser",
  "password": "password123"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com"
    }
  }
}
```

#### 2.3 获取用户信息

**接口**: `GET /api/profile`

**需要认证**: 

**请求头**:
```
Authorization: Bearer <token>
```

**响应示例**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-11-18T11:18:33Z",
    "updated_at": "2025-11-18T11:18:33Z"
  }
}
```

---

### 3. 文章管理

#### 3.1 创建文章

**接口**: `POST /api/posts`

**需要认证**: 

**请求体**:
```json
{
  "title": "我的第一篇博客",
  "content": "这是我的第一篇博客文章的内容"
}
```

**响应示例**:
```json
{
  "code": 201,
  "message": "创建成功",
  "data": {
    "id": 1,
    "title": "我的第一篇博客",
    "content": "这是我的第一篇博客文章的内容",
    "user_id": 1,
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com"
    },
    "created_at": "2025-11-18T11:19:06Z",
    "updated_at": "2025-11-18T11:19:06Z"
  }
}
```

#### 3.2 获取文章列表

**接口**: `GET /api/posts?page=1&page_size=10`

**需要认证**: 

**查询参数**:
- `page`: 页码（默认: 1）
- `page_size`: 每页数量（默认: 10，最大: 100）

**响应示例**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "posts": [
      {
        "id": 1,
        "title": "我的第一篇博客",
        "content": "这是我的第一篇博客文章的内容",
        "user_id": 1,
        "user": {
          "id": 1,
          "username": "testuser"
        },
        "created_at": "2025-11-18T11:19:06Z",
        "updated_at": "2025-11-18T11:19:06Z"
      }
    ],
    "page": 1,
    "page_size": 10,
    "total": 1
  }
}
```

#### 3.3 获取文章详情

**接口**: `GET /api/posts/:id`

**需要认证**: 

**响应示例**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": 1,
    "title": "我的第一篇博客",
    "content": "这是我的第一篇博客文章的内容",
    "user_id": 1,
    "user": {
      "id": 1,
      "username": "testuser"
    },
    "comments": [],
    "created_at": "2025-11-18T11:19:06Z",
    "updated_at": "2025-11-18T11:19:06Z"
  }
}
```

#### 3.4 更新文章

**接口**: `PUT /api/posts/:id`

**需要认证**:  (仅作者)

**请求体**:
```json
{
  "title": "更新后的博客标题",
  "content": "更新后的博客内容"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "id": 1,
    "title": "更新后的博客标题",
    "content": "更新后的博客内容",
    "user_id": 1,
    "created_at": "2025-11-18T11:19:06Z",
    "updated_at": "2025-11-18T11:20:36Z"
  }
}
```

#### 3.5 删除文章

**接口**: `DELETE /api/posts/:id`

**需要认证**:  (仅作者)

**响应示例**:
```json
{
  "code": 200,
  "message": "删除成功"
}
```

---

### 4. 评论功能

#### 4.1 创建评论

**接口**: `POST /api/comments`

**需要认证**: 

**请求体**:
```json
{
  "content": "这是一条评论",
  "post_id": 1
}
```

**响应示例**:
```json
{
  "code": 201,
  "message": "评论成功",
  "data": {
    "id": 1,
    "content": "这是一条评论",
    "user_id": 1,
    "user": {
      "id": 1,
      "username": "testuser"
    },
    "post_id": 1,
    "created_at": "2025-11-18T11:19:57Z",
    "updated_at": "2025-11-18T11:19:57Z"
  }
}
```

#### 4.2 获取文章评论列表

**接口**: `GET /api/comments/post/:post_id`

**需要认证**: 

**响应示例**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "content": "这是一条评论",
      "user_id": 1,
      "user": {
        "id": 1,
        "username": "testuser"
      },
      "post_id": 1,
      "created_at": "2025-11-18T11:19:57Z",
      "updated_at": "2025-11-18T11:19:57Z"
    }
  ]
}
```

---

## 错误码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 201 | 创建成功 |
| 400 | 请求参数错误 |
| 401 | 未认证或认证失败 |
| 403 | 无权限访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 测试命令示例

### 1. 注册用户
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123","email":"test@example.com"}'
```

### 2. 登录
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### 3. 创建文章（需要先获取 token）
```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"title":"我的博客","content":"博客内容"}'
```

### 4. 获取文章列表
```bash
curl -X GET http://localhost:8080/api/posts
```

### 5. 创建评论
```bash
curl -X POST http://localhost:8080/api/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"content":"很好的文章","post_id":1}'
```

---

## 环境变量配置

可以通过环境变量配置以下参数：

- `JWT_SECRET`: JWT 密钥（默认: your-secret-key-change-in-production）
- `SERVER_PORT`: 服务端口（默认: 8080）

示例：
```bash
export JWT_SECRET="my-super-secret-key"
export SERVER_PORT="3000"
./blog-server
```

---

## 功能特性

✅ 用户注册和登录  
✅ JWT 认证和授权  
✅ 密码加密存储（bcrypt）  
✅ 文章 CRUD 操作  
✅ 文章作者权限控制  
✅ 评论功能  
✅ 分页查询  
✅ 统一错误处理  
✅ 请求日志记录  
✅ 软删除支持  

---

## 数据库表结构

### users 表
- `id`: 主键
- `username`: 用户名（唯一）
- `password`: 加密后的密码
- `email`: 邮箱（唯一）
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间（软删除）

### posts 表
- `id`: 主键
- `title`: 文章标题
- `content`: 文章内容
- `user_id`: 作者ID（外键）
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间（软删除）

### comments 表
- `id`: 主键
- `content`: 评论内容
- `user_id`: 评论者ID（外键）
- `post_id`: 文章ID（外键）
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间（软删除）

---