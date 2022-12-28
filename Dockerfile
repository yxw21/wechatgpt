FROM alpine
WORKDIR /home/wechatgpt
RUN apk add ca-certificates chromium xvfb && \
    addgroup -S wechatgpt && adduser -S wechatgpt -G wechatgpt -s /bin/ash && \
    version=$(wget -qO- -t1 -T2 "https://api.github.com/repos/yxw21/wechatgpt/releases/latest" | grep "tag_name" | head -n 1 | awk -F ":" '{print $2}' | sed 's/\"//g;s/,//g;s/ //g') && \
    wget https://github.com/yxw21/wechatgpt/releases/download/$version/wechatgpt_linux_amd64.tar.gz && \
    tar -zxvf wechatgpt_linux_amd64.tar.gz && \
    rm wechatgpt_linux_amd64.tar.gz README.md && \
    chown -R wechatgpt:wechatgpt wechatgpt && \
    chmod 777 wechatgpt
USER wechatgpt
ENTRYPOINT [ "/home/wechatgpt/wechatgpt" ]