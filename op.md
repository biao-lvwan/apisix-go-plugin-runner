# 使用的例子

https://blog.csdn.net/weixin_42873928/article/details/123279381

docker build -t  192.168.8.79:30290/apisix-go-plugin-runner:v1.0.1 -f Dockerfile .

docker-compose -f docker-compose.yaml up -d 

编译
make build

无法获取到请求的路由ID等信息


curl "http://127.0.0.1:9180/apisix/admin/routes/my-response-route" -H "X-API-KEY: <您的-admin-api-key>" -X PUT -d '
{
"uri": "/my-service",
"methods": ["GET"],
"plugins": {
"ext-plugin-post-req": {
"plugins": [
{
"name": "response-metrics",
"conf": {
"field_name": "data.status",
"push_gateway": "http://localhost:9091",
"job_name": "my_service_metrics",
"instance_name": "apisix-gateway-instance-1"
}
}
]
}
},
"upstream": {
"type": "roundrobin",
"nodes": {
"127.0.0.1:8080": 1  # 您的后端服务
}
}
}'