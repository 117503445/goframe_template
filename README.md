# goframe_template

基于 goframe 的 后端模板项目

## dev

1. 复制 ./config/config.toml.example 为 ./config/config.toml
2. 根目录下 `docker-compose up -d`

## CRUD 生成

1. install gf cli tool <https://goframe.org/pages/viewpage.action?pageId=1115782>
2. ./document/create.sql 写建表 SQL
3. go main.go
 gf gen dao
1. ./script/crud_gen/cfg 下新建配置
2. `python crud_gen.py`
3. 