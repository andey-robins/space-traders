#!/bin/bash

auth="'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiU1RBUl9TQ1JJUFQiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMDMtMTAiLCJpYXQiOjE3MTAzNjY4NjYsInN1YiI6ImFnZW50LXRva2VuIn0.gEgwtS7_qtphyb7qit_dkrSJh5fRIEeX8GTGvSkP5Z9Zkwn0Y_XxlcTqkxnYMbdGIT6OiBvnXyrY6r58Hy0ucNebcgL4iM4AGrUR1xOzaRhCPFUkXQsGz-heTqRCjKs0JPPOweExm0wSTb3EEYrLwnV5Zp5J6nqAJ6Klawy8L6mhwqECPGoooFoFg9zTA6mVmbg1UsXWzrCGWlutw7urIleSj-fH-m-YoBBVnfhs1kpyYhifD0RsDOBu8RcAONxr1Z-gc4FhxTblINBrY1HKRR542toYBsP5Ox81FNAvIvrY9XAO83bDz1VODpWhKtQcEq8wtbZfrBPlJf3Y9BuCUw'"

# register an agent
# -- this is what gave us the user token
curl --request POST \
    --url 'https://api.spacetraders.io/v2/register' \
    --header 'Content-Type: application/json' \
    --data '{
        "symbol": "STAR_SCRIPT",
        "faction": "VOID"
    }'

curl 'https://api.spacetraders.io/v2/my/agent' \
    --header $auth