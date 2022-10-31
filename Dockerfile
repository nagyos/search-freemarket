FROM node:16.17.1-alpine

WORKDIR /myapp
# WORKDIR /next-app

COPY ./next-app/package.json ./
COPY ./next-app/yarn.lock ./
# COPY package.json ./
# COPY yarn.lock ./

# RUN yarn install --frozen-lockfile
RUN yarn install

# install golang
RUN apk add --update --no-cache vim git make musl-dev go curl

# install chrome-driver
RUN apk add --update \
    wget \
    udev \
    ttf-freefont \
    chromium \
    chromium-chromedriver \

# FROM node:16.17.1-alpine

# WORKDIR /myapp
# # WORKDIR /next-app

# COPY ./next-app/package.json ./
# COPY ./next-app/yarn.lock ./
# # COPY package.json ./
# # COPY yarn.lock ./

# # RUN yarn install --frozen-lockfile
# RUN yarn install

# # install golang
# RUN apk add --update --no-cache vim git make musl-dev go curl && \
#     export GOPATH= && \
#     export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH && \
#     export GOBIN=$GOROOT/bin && \
#     mkdir -p ${GOPATH}/src ${GOPATH}/bin && \
#     export GO111MODULE=off

# # install chrome-driver
# RUN apk add --update \
#     wget \
#     udev \
#     ttf-freefont \
#     chromium \
#     chromium-chromedriver \