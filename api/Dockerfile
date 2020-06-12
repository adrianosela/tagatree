FROM alpine:latest AS api
RUN apk add --update bash curl && rm -rf /var/cache/apk/*
COPY ./tagatree /bin/tagatree
EXPOSE 443
EXPOSE 80
CMD ["/bin/tagatree"]
