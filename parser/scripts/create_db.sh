#!/bin/bash

psql --command "CREATE USER testing WITH SUPERUSER PASSWORD 'testing';"
createdb -O testing testing
psql testing -f ./parser/sql/schema.sql
