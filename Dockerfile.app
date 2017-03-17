FROM alpine

COPY app /usr/bin/app

EXPOSE 12345

CMD ["app"]
