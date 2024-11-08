# Ditto file processor

## What?

This repository contains the source code for Go application. It does the following:
* takes a directory with JSON files imported from Ditto Words
* for each file creates a new one without the project key. In other words, the imported file has `project_key: {object}` structure, the outputted one has `{object}` structure

This program is intended to be used in CI as a Docker image.

## Requirements

* `input` directory contains JSON files
* JSON files have names in format of `components__{project-name}__{locale}.json`
    * `{project-name}` is the name of a Ditto project
    * `{locale}` is a locale name, for example `base` for English, `es-mx` for Mexican Spanish


## Docker image

Docker image available on Docker Hub: https://hub.docker.com/repository/docker/paveltovchigrechko/ditto-file-processor/general

Use example:
```
docker container run -v ./input:/app/input -v ./output:/app/output paveltovchigrechko/ditto-file-processor:main

docker cp container-name:app/output /path/to/output
```