FROM alpine:3.20.3
ARG WEB_SERVER_PORT
ARG APP_SERVER_PORT

# web server
EXPOSE ${WEB_SERVER_PORT}

# app server
EXPOSE ${APP_SERVER_PORT}

# install devenvs
RUN apk add --update --no-cache go npm
