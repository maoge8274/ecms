FROM golang:latest

#Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee


#Expose the application on port 8080
EXPOSE 8080

CMD["bee","run"]
