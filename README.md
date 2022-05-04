# seeder

The only tool for seeding volumes and storages (including databases as well).

## Docs

https://pkorobeinikov.github.io/seeder/

## Tutorials

- `postgres`: https://pkorobeinikov.github.io/seeder/tutorial/postgres/
- `vault`: https://pkorobeinikov.github.io/seeder/tutorial/vault/
- `s3`: https://pkorobeinikov.github.io/seeder/tutorial/s3/

## Known seeders

- s3
- vault
- postgres
- kafka

## Install

```shell
$ cd ./seeder
$ go install ./cmd/...
```

## Flags

- `-c` allows to specify seeder configuration file.
  > Note: dir from config file will be used as working dir for seed files.
- `-seeder-helper $name` shows help for specified seeder.
- `-known` shows known seeders

## Example config: `seeder.yaml`

```yaml
seeder:
  state:
    # seeding s3 data
    - name: s3 plain text file seed
      type: s3
      config:
        - bucket: "bucket"
          object-name: "seeded/file/seed.txt"
          option:
            content-type: text/plain
            content-encoding: utf8
          file: seed.txt

    - name: s3 json file seed
      type: s3
      config:
        - bucket: "bucket"
          object-name: "seeded/file/seed.json"
          option:
            content-type: application/json
            content-encoding: utf8
          file: seed.json

    # seeding vault secrets
    - name: vault file seed
      type: vault
      config:
        - key: "secret/data/seed/file/json"
          file: seed.json
        - key: "secret/data/seed/file/yaml"
          file: seed.yaml
        - key: "secret/data/seed/file/yml"
          file: seed.yml

    # seeding postgres data
    - name: postgres file seed
      type: postgres
      config:
        - file: seed.sql

```
