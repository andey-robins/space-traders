#!/bin/sh

atlas schema inspect \
  --url "sqlite://../st.db" \
  --format "{{ sql . }}" > schema.sql.tmp

cat schema.sql.tmp

# https://stackoverflow.com/questions/3231804/in-bash-how-to-add-are-you-sure-y-n-to-any-command-or-alias
read -r -p "Save migrated? [y/N] " response
case "$response" in
    [yY][eE][sS]|[yY]) 
        mv schema.sql.tmp schema.sql
        ;;
    *)
        rm schema.sql.tmp
        ;;
esac