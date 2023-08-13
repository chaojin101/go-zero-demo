#!/usr/bin/env bash

# https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/deploy/script/mysql/genModel.sh

host=172.30.112.1
port=3306
username=root
passwd=password
dbname=bookstore

table=$1
outDir=./model

echo "开始创建库：$dbname 的表：$table"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${table}"  -dir="${outDir}" --style=goZero