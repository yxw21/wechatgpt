# 支持的功能
1. 私聊回复
2. 群里@回复
3. 好友请求处理
# 使用
### 直接下载二进制文件运行
需要安装[依赖](https://github.com/yxw21/wechatgpt#%E4%BE%9D%E8%B5%96)
并且提供一些[环境变量](https://github.com/yxw21/wechatgpt#%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F)
然后运行
```
./wechatgpt
```

### 使用docker

参考 https://hub.docker.com/r/yxw21/wechatgpt

# 微信登录流程
目前只支持扫描终端二维码登录。

# CHATGPT登录流程
如果提供了`WECHAT_CHAT_GPT_USERNAME`和`WECHAT_CHAT_GPT_PASSWORD`会自动登录获取`AccessToken`。

`WECHAT_CHAT_GPT_USERNAME`、`WECHAT_CHAT_GPT_PASSWORD`和`WECHAT_CHAT_GPT_ACCESS_TOKEN`必须提供一项

# 依赖
### Xvfb （只有linux环境需要安装)

Ubuntu or Debian
```
apt update
apt install xvfb
```
CentOS
```
yum update
yum install xorg-x11-server-Xvfb
```
Alpine
```
apk update
apk add xvfb
```
### Key
登录需要谷歌验证码，引入了第三方破解，需要去网站`nopecha.com`购买key，价格很便宜

```
https://nopecha.com
```
### Chrome

Ubuntu or Debian
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
apt install ./google-chrome-stable_current_amd64.deb
```
CentOS
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
yum localinstall -y google-chrome-stable_current_x86_64.rpm
```
Alpine
```
apk add chromium
```


# 环境变量

### WECHAT_CHAT_GPT_USERNAME (可选)
openai用户名
### WECHAT_CHAT_GPT_PASSWORD (可选)
openai密码
### WECHAT_MSG_RETRY（可选）
chatgpt请求失败重试的次数，次数越多回复消息就越慢(默认3)
### WECHAT_KEY （可选）
破解谷歌验证码需要的key，需要去网站`nopecha.com`购买

如果提供了`WECHAT_CHAT_GPT_USERNAME`和`WECHAT_CHAT_GPT_PASSWORD`则必须提供`WECHAT_KEY`

提供的是`WECHAT_CHAT_GPT_ACCESS_TOKEN`就可以忽略
### WECHAT_CHAT_GPT_ACCESS_TOKEN (可选)
大概7天过期
1. 登录 https://chat.openai.com
2. 访问 https://chat.openai.com/api/auth/session
### WECHAT_CHAT_GPT_POLICY (可选)
好友添加策略

同意好友添加请求
```
WECHAT_CHAT_GPT_POLICY = agree
```
不处理好友添加请求
```
WECHAT_CHAT_GPT_POLICY = ignore
```
当验证消息与给定的正则表达式匹配时才会同意添加
```
WECHAT_CHAT_GPT_POLICY = agree,123456
```
远程验证（GET请求`http://example.com/验证信息`），响应状态码为201同意添加
```
WECHAT_CHAT_GPT_POLICY = agree,https://example.com
```
