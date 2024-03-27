# rosetta-cli

## Install on macOS with Brew
```bash
brew tap rosetta/homebrew-rosetta
brew install rosetta
```

## Run using docker
### Setup tenant
```
export ROSETTA_TENANT=<my-tenant>
```
### Create Custom Configuration YAML
To tailor the behavior of the filters, create a custom configuration YAML file using the following command:
```bash
docker run --rm -v $PWD:/app -w /app -e ROSETTA_TENANT=$ROSETTA_TENANT effoeffi/rosetta:main config
```
### Count tokens
Use the following command to count the number of tokens in the specified language (e.g. go, java, js, python):
```bash
docker run --rm -v $PWD:/app -w /app -e ROSETTA_TENANT=$ROSETTA_TENANT effoeffi/rosetta:main count /rest -l go
```

Generate OpenAPI Specification
```bash
docker run --rm -v $PWD:/app -w /app -e ROSETTA_TENANT=$ROSETTA_TENANT effoeffi/rosetta:main gen /rest -l go
```

### Debug on localhost
Use `-e ROSETTA_HOME=localhost:8080 --network=host`
