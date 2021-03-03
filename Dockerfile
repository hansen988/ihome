

FROM golang:latest
 
#创建工作目录
RUN mkdir -p /go/src/go_test/static/
 
#进入工作目录
WORKDIR /go/src/go_test
 
#将DockerFile文件所在目录下的所有文件复制到指定位置
COPY . /go/src/go_test
 
#下载beego、bee以及mysql驱动
#RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
 
#端口
EXPOSE 8080
 
#CMD ["bee","run"]
#修改docker_test文件的权限
RUN chmod 777 ./ihome
#执行docker_test文件
ENTRYPOINT ["./ihome"]