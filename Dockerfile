FROM iron/go

RUN mkdir /app
WORKDIR /app
ADD voucher-datastore-service /app/
ADD .env.keys /app/

ENTRYPOINT ["./voucher-datastore-service"]
