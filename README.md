# effective-bassoon

Physic-guide is a NSQ producer event with topic: `effective_bassoon`

## How to use

Start nsqd:

``` shell
nsqlookupd & 
nsqd --lookupd-tcp-address=127.0.0.1:4160 &
```

Run the producer with the following command:

```shell
go run producer.go
```

You will find a consummer available [here](https://github.com/MaxencePPierre/psychic-guide)

## Build and run in docker

To build the application:

```shell
docker build --tag effective-bassoon .
```

In order to run this newly created image, use docker run command:

```shell
docker run -p 4150:4150 -it effective-bassoon
```
