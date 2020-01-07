# 镜像来源
FROM alpine
# 镜像制作者
LABEL maintainer="taoshumin_vendor@sensetime.com"
# 宿主机拷贝
COPY ./music-web-go /tmp/music-web-go
# 进入容器目录
WORKDIR /tmp/
# RUN
RUN chmod +x music-web-go
# 启动容器运行的第一条命令
ENTRYPOINT ["/tmp/music-web-go"]
# 对外端口
EXPOSE 9092