from golang

copy . /app
workdir /app

RUN GOPROXY=https://goproxy.io go build -o /rater

entrypoint /rater

expose 50051