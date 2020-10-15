# go-blog
a blog service written by golfing

P91

1. 新增标签
```
curl -X POST http://localhost:8000/api/v1/tags -F 'name=Go' -F created_by=Henry
```

2. 获取标签列表
```
curl -X GET 'http://localhost:8000/api/v1/tags?page=1&page_size=2'
```

3.  修改标签
```
curl -X PUT http://localhost:8000/api/v1/tags/1 -F state=0 -F modified_by=Lee
```

4. 删除标签
```
curl -X DELETE http://localhost:8000/api/v1/tags/1
```