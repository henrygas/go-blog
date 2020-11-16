# go-blog
a blog service written by golfing

## 进度
https://github.com/go-programming-tour-book/blog-service

P163

## 1. 标签
### 1.1 新增标签
```
curl -X POST http://localhost:8000/api/v1/tags -F 'name=Go' -F created_by=Henry
```

### 1.2 获取标签列表
```
curl -X GET 'http://localhost:8000/api/v1/tags?page=1&page_size=2'
```

### 1.3  修改标签
```
curl -X PUT http://localhost:8000/api/v1/tags/1 -F state=0 -F modified_by=Lee
```

### 1.4 删除标签
```
curl -X DELETE http://localhost:8000/api/v1/tags/1
```

## 2. 文章

### 2.1 新增文章
```
curl -X POST 'http://localhost:8000/api/v1/articles' \
-F tag_id\=1 \
-F 'title=新增文章01-标题' \
-F 'desc=新增文章01-简述' \
-F cover_image_url\=https://www.eddycjy.com \
-F 'content=新增文章01-内容' \
-F created_by\=henry \
-F state\=1
```

### 2.2 修改文章
```
curl -X PUT 'http://localhost:8000/api/v1/articles/1' \
-F tag_id\=1 \
-F 'title=测试文章-标题-更新' \
-F 'desc=测试文章-简述-更新' \
-F cover_image_url=https://ww.eddycjy.com \
-F 'content=测试文章-内容-更新' \
-F 'modified_by=henry-更新' \
-F state\=1

```

### 2.3 删除文章
```
curl -X DELETE 'http://localhost:8000/api/v1/articles/1'
```

### 2.4 获取指定文章
```
curl -X GET 'http://localhost:8000/api/v1/articles/1'
```

### 2.5 获取文章列表
```
curl -X GET 'http://localhost:8000/api/v1/articles?tag_id=1&page=1&page_size=10'
```

## 3. 文件上传

### 3.1 上传文件
```
curl -X POST 'http://localhost:8000/upload/file' \
-F 'file=@/Users/zhengguang.li/code/go_play/go-blog/tony.jpeg' \
-F 'type=1'
```

## 4. 鉴权
### 4.1 获取token
```
curl -X GET 'http://localhost:8000/auth?app_key=henry-key&app_secret=henry-secret'
```

### 4.2 没有获取token
```
curl -X GET 'http://localhost:8000/api/v1/tags'
```

### 4.3 token错误
```
curl -X GET 'http://localhost:8000/api/v1/tags' -H 'token:abcdefghijklmn'
```

### 4.4 token超时
```
curl -X GET 'http://localhost:8000/api/v1/tags' -H 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiYWIyODVkYTc5NTgzOTQ5NGE2OTJhN2Y1NmU3NWUzZGQiLCJhcHBfc2VjcmV0IjoiODBlMjc2YzFmMjBlZDdmOGJjMmMzZWIzNWUxYTAxZmUiLCJleHAiOjE1MDMzNDU5NTQsImlzcyI6ImJsb2ctc2VydmljZSJ9.eiDvkMTwQJKUyOmgNlJ0DER9hTjiTVYRMl0wkiSvlhc'
```

## 5. 使用流程说明
+ 通过下列链接获得服务端生成的token
```
curl -X GET 'http://localhost:8000/auth?app_key=henry-key&app_secret=henry-secret'
```
+ 带token去访问API
```
curl -X GET 'http://localhost:8001/api/v1/tags?page=1&page_size=2&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiYWIyODVkYTc5NTgzOTQ5NGE2OTJhN2Y1NmU3NWUzZGQiLCJhcHBfc2VjcmV0IjoiODBlMjc2YzFmMjBlZDdmOGJjMmMzZWIzNWUxYTAxZmUiLCJleHAiOjE2MDU1MDQzMTAsImlzcyI6ImJsb2ctc2VydmljZSJ9.MLoT7dpLiKBwgH10yqCjfb3fO0goH-67c8M-FkkYdsw'
```