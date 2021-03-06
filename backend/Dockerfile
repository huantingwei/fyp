FROM golang:1.16

WORKDIR /app

COPY . .

RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-330.0.0-linux-x86_64.tar.gz
RUN mkdir -p /usr/local/gcloud && tar -C /usr/local/gcloud -xvf google-cloud-sdk-330.0.0-linux-x86_64.tar.gz && /usr/local/gcloud/google-cloud-sdk/install.sh
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin
RUN gcloud components install kubectl

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]


