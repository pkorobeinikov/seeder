# Showcase

Current folder contains working seeder usage examples.


## How to use

1. Start containers:

    ```shell
    $ docker-compose up
    ```

2. Wait until all containers are up and running.
3. Run any seeder (in new console).
4. Check seeded data by opening in browser http://localhost:9000.
5. Stop containers by pressing Ctrl-c.

### Kafka seeder

Seeder run example:

> Note: port must be same as exposed in `docker-compose.yml`.

```shell
$ SEEDER_KAFKA_PEER=127.0.0.1:9092 seeder -c ./401_kafka/seeder.yaml
```

Output:

```
working dir: 401_kafka
seeding json file: 401_kafka/seed.json
seeded items: 2
seeding yaml file: 401_kafka/seed.yml
seeded items: 2
```

> Note: output may differ.

### Vault seeder

Seeder run example:

> Note: port must be same as exposed in `docker-compose.yml`.

```shell
$ SEEDER_VAULT_ADDRESS=http://0.0.0.0:8200 SEEDER_VAULT_TOKEN=secret seeder -c ./201_vault_file/seeder.yaml
```

### S3

TBD

### Postgres

TBD
