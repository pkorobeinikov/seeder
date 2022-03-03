# Showcase

Current folder contains working seeder usage examples.

## How to use

### Kafka seeder

1. Run:
   ```shell
   $ docker-compose up
   ```
2. Wait until all containers are up and running.
3. In other shell run seeder with ENV variables and config. Example:
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

4. Run your app to consume seeded data.

### Vault seeder

Seeder run example:

```shell
$ SEEDER_VAULT_ADDRESS=http://0.0.0.0:8200 SEEDER_VAULT_TOKEN=secret seeder -c ./201_vault_file/seeder.yaml
```
