# 环境变量
###### CHAT_GPT_SESSION_TOKEN (chatgpt会话token, 大约一个月过期)
1. 登录 https://chat.openai.com
2. 获取`cookie` `__Secure-next-auth.session-token`
###### CHAT_GPT_WECHAT_POLICY (好友添加策略)
同意好友添加请求
```
CHAT_GPT_WECHAT_POLICY = agree
```
不处理好友添加请求
```
CHAT_GPT_WECHAT_POLICY = ignore
```
当验证消息与给定的正则表达式匹配时才会同意添加
```
CHAT_GPT_WECHAT_POLICY = agree,123456
```
远程验证（GET请求http://example.com/验证信息），响应状态码为201同意添加
```
CHAT_GPT_WECHAT_POLICY = agree,https://example.com
```
# 使用(可直接下载二进制文件)
```
./wechatgpt_linux_amd64
```
在浏览器打开url扫码登录即可