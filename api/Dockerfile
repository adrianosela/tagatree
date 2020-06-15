FROM ubuntu:latest
RUN apt-get install -y bash && rm -rf /var/cache/apk/*
COPY ./tagatree /bin/tagatree
EXPOSE 443
EXPOSE 80
CMD ["/bin/tagatree"]
