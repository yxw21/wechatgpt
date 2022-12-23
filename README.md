
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
### WECHAT_KEY （必填）
破解谷歌验证码需要的key，需要去网站`nopecha.com`购买
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
# 微信登录流程
目前只支持扫描终端二维码登录。

# CHATGPT登录流程
如果提供了chatgpt用户名和密码会使用账号密码登录获取token。

账号密码和token必须提供一项
# 支持的功能
1. 私聊
2. 群里@
3. 好友请求
# 使用(可直接下载二进制文件)
```
./wechatgpt
```
# Docker
https://hub.docker.com/r/yxw21/wechatgpt
