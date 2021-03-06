FROM golang:1.16

WORKDIR /app

COPY . .

RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-330.0.0-linux-x86_64.tar.gz
RUN mkdir -p /usr/local/gcloud && tar -C /usr/local/gcloud -xvf google-cloud-sdk-330.0.0-linux-x86_64.tar.gz && /usr/local/gcloud/google-cloud-sdk/install.sh
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin

RUN gcloud components install kubectl && gcloud components install beta

RUN go build -o app .

EXPOSE 8989

CMD ["./app"]


still encounter this error...

Connected to MongoDB...
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

panic: stat /root/.kube/config: no such file or directory

goroutine 1 [running]:
github.com/huantingwei/fyp/util.GetKubeClientSet(0xc000690d20)
	/app/util/kubeAuth.go:25 +0x267
github.com/huantingwei/fyp/overview.NewService(0xc00007f680, 0xc0002ae380, 0xc000690d20)
	/app/overview/api.go:30 +0x1ba
main.main()
	/app/main.go:28 +0xec
