FROM golang:1.12.3-alpine3.9

WORKDIR /app

RUN apk --update add git nodejs nodejs-npm && rm -rf /var/cache/apk/* && \
    npm install -g bower grunt-cli && \
    echo '{ "allow_root": true }' > /root/.bowerrc

COPY . /app
RUN bower install
RUN npm install
RUN go get github.com/gorilla/mux
RUN go get github.com/golang/glog
RUN go get github.com/bobbydeveaux/ngis-status/
RUN go get github.com/bobbydeveaux/ngis-status/app/common
RUN go get github.com/bobbydeveaux/ngis-status/app/home
RUN go get github.com/bobbydeveaux/ngis-status/app/api
RUN go build main.go

CMD ./main

EXPOSE 8181