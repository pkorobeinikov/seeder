# seeder

The only tool for seeding storages (including databases as well).

## Example config

```yaml
seeder:
  state:
    # seeding vault secretes
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

## Similar projects

* https://github.com/go-testfixtures/testfixtures
