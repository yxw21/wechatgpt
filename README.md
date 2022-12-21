
# 依赖
- linux需要安装`xvfb`

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
- 登录需要谷歌验证码，引入了第三方破解，需要去网站`nopecha.com`购买key，价格很便宜

```
https://nopecha.com
```
- 浏览器

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
### WECHAT_KEY （必填）
破解谷歌验证码需要的key，需要去网站`nopecha.com`购买
### WECHAT_CHAT_GPT_TOKEN (可选)
chatgpt会话token, 大约一个月过期，获取步骤如下
1. 登录 https://chat.openai.com
2. 获取`cookie` `__Secure-next-auth.session-token`
<img width="945" alt="image" src="https://user-images.githubusercontent.com/16237562/206679314-7d412b03-98fc-422d-92bb-2d4a19f375b8.png">

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
