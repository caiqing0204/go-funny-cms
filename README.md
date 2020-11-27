## 项目

### 项目目录整体划分

- app 核心目录 ，控制器、model 等都在这里

- bootstrap:  辅助函数

- config:  配置加载

- pkg:  配置加载service

- public:  静态文件

- resources: 资源文件/前端项目

- .air.toml:  监听

- .env: 配置

- .env.example: 辅助配置参考

### 前端后台

    # 参考至
    git clone https://github.com/vueComponent/ant-design-vue-pro.git
    
目录地址: `resource/front`
    
编译:
    
    # 项目根目录
    
    npm install
    npm run server
    
### 前端客户端

目录地址: `resource/views`
    
    
### 第一次操作

    cp .env.example .env
    go build main.go && ./main generate-jwt
    # 或者
    go run main.go generate-jwt
    :按提示操作
    
    # 开发环境配置
    # 安装air
    # 根目录运行
    air
    
### 进度

**11-21**: 

    jwt、响应、日志、db等操作基本完成
    
**11-24**

    配置jwt密钥生成
    
**11-25**

    # 创建账户
    > go run main create-admin-user -h
    > go run create-admin-user --account [你的账户名称]
    
    # 登陆
    > air
    curl --location --request POST '127.0.0.1:8082/api/admin/login' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'account=chenf' \
    --data-urlencode 'password=12345678'

**11-27**

    # 新增中间件
    
    # 新增配置
    # 默认分钟
    # pkg/auth/jwt.go:25
    JWT_EXPIRE_AT=10
    
    # 前端登陆注销对接完成
