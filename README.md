# quake3-importer

[![build](https://github.com/adobromilskiy/quake3-importer/actions/workflows/ci.yml/badge.svg)](https://github.com/adobromilskiy/quake3-importer/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/adobromilskiy/quake3-importer)](https://goreportcard.com/report/github.com/adobromilskiy/quake3-importer) [![Coverage Status](https://coveralls.io/repos/github/adobromilskiy/quake3-importer/badge.svg?branch=main&kill_cache=1)](https://coveralls.io/github/adobromilskiy/quake3-importer?branch=main)

Application to import quake3 stats files into mongodb.

**Note:** application parse only FFA stats file. You are wellcome to contribute.

- Repositories:
  - GitHub: [github.com/adobromilskiy/quake3-importer](https://github.com/adobromilskiy/quake3-importer)
  - DockerHub: [adobromilskiy/quake3-importer](https://hub.docker.com/r/adobromilskiy/quake3-importer)


- Dockerfile: [https://github.com/adobromilskiy/quake3-importer/blob/main/Dockerfile](https://github.com/adobromilskiy/quake3-importer/blob/main/Dockerfile)


- Maintained by: [Alexander Dobromilskiy](https://twst.dev)

## Quick start

Clone repository and run `make build` command. Go to the **app** directory and run application with next parameters:

- **dbconn** - connection string to mongodb
- **dbname** - name of the database
- **path** - path to the directory with stats files

Example:

```console
git clone git@github.com:adobromilskiy/quake3-importer.git
cd quake3-importer
make build
cd app
./app --dbconn=mongodb://localhost:27017 --dbname=quake3 --path=/path/to/stats/files
```

or via docker:

```console
docker run -v /path/to/stats/files:/stats --network=mongo_network adobromilskiy/quake3-importer:latest /app --dbconn=mongodb://mongohost:27017 --dbname=quake3 --path=/stats
```