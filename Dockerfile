# 使用 alpine 镜像作为基础镜像
FROM alpine AS prod

# 拷贝执行文件
COPY main /main
# 拷贝配置文件
COPY ./config.yml ./config.yml
COPY ./json/ ./json/
COPY log.txt ./log.txt

# 添加执行权限
RUN chmod +x /main
RUN chmod 777 /log.txt

RUN echo -e 'https://mirrors.aliyun.com/alpine/v3.6/main/\nhttps://mirrors.aliyun.com/alpine/v3.6/community/' > /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 定义容器启动时执行的命令
CMD ["./main"]
