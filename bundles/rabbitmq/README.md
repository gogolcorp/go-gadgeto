# RabbitMQ bundle

This bundle is used to communicate with rabbitMQ in your project
It creates the needed files to init a connection, produce messages and consume them

## Setup

Once the files have been copied to your project, you just have to init connection, create your producer and your consumer and that's it !

You have to create the `rabbitMQ service` inside your docker-compose and provide the `AMQP_URL` env variable.

docker-compose service :

```yml
rabbitmq:
  image: rabbitmq:3-management
  ports:
    - "5672:5672"
    - "15672:15672"
  environment:
    RABBITMQ_DEFAULT_USER: user
    RABBITMQ_DEFAULT_PASS: bitnami
  networks:
    - backend
  healthcheck:
    test: [ "CMD", "curl", "-f", "http://localhost:15672" ]
    interval: 30s
    timeout: 10s
    retries: 5
```
