#!/bin/bash

work_path='../config/'
cd ${work_path}
openssl genrsa 2048 > private.dev.pem
openssl rsa -in private.dev.pem -pubout > public.dev.pem
# to pkcs8 format
# openssl pkcs8 -topk8 -inform PEM -in private.dev.pem -outform PEM -nocrypt -out private_pkcs8.dev.pem

openssl genrsa 2048 > private.prod.pem
openssl rsa -in private.prod.pem -pubout > public.prod.pem
# to pkcs8 format
# openssl pkcs8 -topk8 -inform PEM -in private.prod.pem -outform PEM -nocrypt -out private_pkcs8.prod.pem
