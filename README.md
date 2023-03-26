Create connection to rabbitmq local, through to command:

`docker run -d -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=user rabbitmq:3.11.11-management`