# TuringBackend
图灵杯网站后端

## 运行项目
### 前提条件
1. docker
2. docker compose
### 运行
#### windows
```
git clone https://github.com/TuringCup/TuringBackend.git
cd TuringBackend
docker compose pull
docker compose build
docker compose up -d
```
#### Linux
```
git clone https://github.com/TuringCup/TuringBackend.git
cd TuringBackend
sudo docker-compose pull
sudo docker-compose build
sudo docker-compose up -d
```

等待一段时间后，项目成功运行，在浏览器访问http://localhost:5001/api/ping

如果看到`success`则表示项目运行成功
### 停止运行
#### windows
在项目根目录下
```
docker compose down
```
#### Linux
```
sudo docker-compose down
```
## 如何贡献
1. 从main分支创建一个新的功能分支 例如: feature_xxx
2. 在feature_xxx分支下进行开发，然后commit push
3. 最后提交一个pull request请求主分支合并
4. 通知其他人对代码进行审阅

## 项目结构
> 参考 https://github.com/CocaineCong/gin-mall
- routes 定义了各个路径对应的handler函数
- api 定义各个handler
- service 实现各个handler的具体实现细节
- types 各个请求和响应的结构体定义 例如: UserRegisterReq UserRegisterRsp
- config 服务运行所需的各种配置
- repository/db/dao 对数据库的一些操作
- repository/db/model 数据库中的各种数据模型
- repository/cache redis缓存的一些操作
- middleware 服务所需的一些中间件 比如: token验证之类的
- data 存放持久化数据
- pkg 额外定义的包，比如错误处理，发送邮件
- consts 服务运行时候需要的各种常量
