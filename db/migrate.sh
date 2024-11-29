#!/bin/sh

cp schema.sql newschema.tmp.sql

atlas schema apply \
  --url "sqlite://../st.db" \
  --to "file://newschema.tmp.sql" \
  --dev-url "sqlite://file?mode=memory"

rm newschema.tmp.sql

./export.sh
