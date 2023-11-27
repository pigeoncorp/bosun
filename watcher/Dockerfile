FROM alpine:3.18

RUN apk upgrade && apk add \
    curl \
    git

# install kubectl
RUN curl -L https://dl.k8s.io/release/v1.28.4/bin/linux/amd64/kubectl > /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl

# install helm
RUN curl https://get.helm.sh/helm-v3.13.2-linux-amd64.tar.gz | tar -C /opt -xvz && ln -s /opt/linux-amd64/helm /usr/local/bin/helm
