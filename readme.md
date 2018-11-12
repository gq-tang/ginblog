# **gin Blog**
---
## 概述
- 用gin框架搭建的简单博客系统
- UI部分FORK [beegoblog](https://github.com/lock-upme/beegoblog)
- 项目结构MVC
- 所有静态资源用assets打包
- 数据库采用mysql，支持自动迁移
- 支持make编辑
## linux下运行说明
- go get github.com/gq-tang/ginblog
- cd $GOPATH/src/github.com/gq-tang/ginblog
- make dev-requirements
- make build 
- ./build/blog