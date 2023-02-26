
goctl api go -api api/ops.api -dir .

goctl rpc protoc ops-rpc.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

docker build -t opsrpc:v1 -f rpc/Dockerfile .
docker run --name test-rpc -p 9001:9001 opsrpc:v1 

docker build -t ops:v1 -f app/Dockerfile .
docker run --name test-api -p 8888:8888 ops:v1

生成k8s部署文件 rpc
goctl kube deploy -name ops-rpc -namespace my-cicd -image jiangbo2018/ops-rpc:v1.0.1 -o ops-rpc.yaml -port 9001

goctl kube deploy -name ops-api -namespace my-cicd -image jiangbo2018/ops-api:v1.0.1 -o ops-api.yaml -port 8888
kubectl apply -f ops-api.yaml

环境总结：
docker容器
dokcer k8s
docker gitlab启动
gitlab-runner start 通过brew安装的   容器安装后总是找不到docker和make命令
kuboard加入k8s

以上验证OK
下一步：
1 gitlab-runner直接执行k8s apply应用新镜像。此镜像从本地获取
2 搭建harbor
3 把镜像push到harbor
4 gitlab-runner直接执行k8s apply应用新镜像。此镜像从harbor获取
配置问题！ 前端问题！ 网关问题 没用用到ingress/网关 etcd和配置中心 数据库连接等敏感信息（variable应该可以解决）
