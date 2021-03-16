# Backend API Server

## Installation

sudo su

### Golang

wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

### Google Cloud Platform SDK

curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-331.0.0-linux-x86_64.tar.gz
tar xvzf google-cloud-sdk-331.0.0-linux-x86_64.tar.gz
./google-cloud-sdk/install.sh

gcloud components install kubectl
gcloud components install beta

### Docker
```
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

verify with
`sudo docker run hello-world`

### Database Setup: MongoDB

sudo docker run -dp 27017:27017 -v $(pwd)/tmp/db:/data/db mongo:4.2-bionic


## Run

for development:
`export FYPENV="dev"`

for production:
`export FYPENV="prod"`

### ENV
```
export PROJECTNAME=your_project_name \
export CLUSTERNAME=your_cluster_name \
export ZONENAME=your_zone_name \
export CLUSTER="projects/$PROJECTNAME/locations/$ZONENAME/clusters/$CLUSTERNAME" \
export CRED=path_to_your_gcp_cred_json_file \
```

for example:
```
export CLUSTERNAME="demo" \
export ZONENAME="us-central1-a" \
export PROJECTNAME="fyp-demo-306511" \
export CLUSTER="projects/$PROJECTNAME/locations/$ZONENAME/clusters/$CLUSTERNAME"
export CRED="/home/justbadcodes/fyp/backend/fyp-demo-sa.json"
```

