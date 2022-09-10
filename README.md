# quake3-importer

[![build](https://github.com/adobromilskiy/quake3-importer/actions/workflows/ci.yml/badge.svg)](https://github.com/adobromilskiy/quake3-importer/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/adobromilskiy/quake3-importer)](https://goreportcard.com/report/github.com/adobromilskiy/quake3-importer) [![Coverage Status](https://coveralls.io/repos/github/adobromilskiy/quake3-importer/badge.svg?branch=main&kill_cache=1)](https://coveralls.io/github/adobromilskiy/quake3-importer?branch=main)

Application to import quake3 stats files into mongodb.

**Note:** application parse only FFA stats file. You are wellcome to contribute.

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