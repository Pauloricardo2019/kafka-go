# KAFKA WITH GOLANG

### In this repository, I save some concepts and examples of Kafka with Golang.
### Functions as Producer, Consumer.

#### To see the examples, you need to have Docker installed in your machine.
#### And just run docker-compose up -d, for up the container

#### CLI - COMMANDS

```
*CRIAR UM TÓPICO= 

kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

*VER DETALHER DO TÓPICO=

kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe

*PRODUZIR EM UM TÓPICO=

kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

*CONSUMIR DE UM TÓPICO= 

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

*LER TODAS AS MENSAGEM JÁ ENVIADAS NO TÓPICO=

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

*CRIAR CONSUMIDORES DO MESMO GRUPO=

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --group=x

*VER DETALHES DOS CONSUMERS DE UM MESMO GRUPO=

kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe
