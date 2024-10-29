#!/bin/bash

podman run -d --rm --name imageset \
  -e POSTGRESQL_USER=imagesetuser \
  -e POSTGRESQL_PASSWORD=imagesetpassword \
  -e POSTGRESQL_ADMIN_PASSWORD=adminpassword \
  -e POSTGRESQL_DATABASE=imageset \
  -p 5432:5432 \
  registry.redhat.io/rhel9/postgresql-16:1-29.1726696141
