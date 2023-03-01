## Description

This poc has been created to study and understand how implement a kafka-producer and consumer.

## Run

Run the commands

```shell
make tidy
```
To download the libraries necessary

```shell
make env-up
```
It will create the docker images necessary to run.

Then run:

```shell
make env-setup
```

It will set up the topic to the app to start


Finally:

```shell
make run
```
To start producer and consumer

More useful commands, see Makefile.