#!/bin/bash

work_path='../config/'
cd ${work_path} || exit

openssl genrsa 2048 > private.pem
openssl rsa -in private.pem -pubout > public.pem
# to pkcs8 format
# openssl pkcs8 -topk8 -inform PEM -in private.pem -outform PEM -nocrypt -out private_pkcs8.pem
