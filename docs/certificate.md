### x509 子命令完成证书查看，签发证书的功能
####  查看证书详情
```shell
openssl x509 -in xx.crt -text
```


#### 生成自签发证书 req子命令，可以生成csr 查看csr
##### 1. 生成csr，同时生成秘钥对 -nodes 表示不加密秘钥
```shell
openssl req -newkey rsa:2048 -nodes  -keyout ca.key -out ca.csr -subj "/C=CN/ST=SH/L=SH/O=xiaohongshu/OU=serverless/CN=rsi.xiaohongshu.com"
```

##### 2. 自签发证书用自己的key签发
```shell
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

```shell
Signature ok
subject=/C=CN/ST=SH/L=SH/O=xiaohongshu/OU=serverless/CN=rsi.xiaohongshu.com
Getting Private key
```

#### 使用自签发的证书签发新的Server证书
###### 1. 生成csr，同时生成秘钥对 -nodes 表示不加密秘钥，
```shell
openssl req -newkey rsa:2048 -nodes -keyout server.key -out server.csr -config ../test/data/ssl/openssl.cnf -subj "/C=CN/ST=SH/L=SH/O=xiaohongshu/OU=serverless/CN=*.rsi-apiserver.xiaohongshu.com"
```

###### 2. 生成新的Server证书
```shell
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
```

###### 3.验证证书是否是ca签发的
```shell
openssl verify -CAfile ca.crt server.crt
```

#### 使用自签发的证书签发新的Server证书
###### 1. 生成csr，同时生成秘钥对 -nodes 表示不加密秘钥，
```shell
openssl req -newkey rsa:2048 -nodes -keyout server.key -out server.csr -config ../test/data/ssl/openssl.cnf -subj "/C=CN/ST=SH/L=SH/O=xiaohongshu/OU=serverless/CN=local.rsi-apiserver.xiaohongshu.com"
```

###### 2. 生成新的Server证书
```shell
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extfile ../test/data/ssl/openssl.cnf -extensions v3_req
```

###### 3.验证证书是否是ca签发的
```shell
openssl verify -CAfile ca.crt server.crt
```

#### 使用自签发的证书签发新的Client证书
###### 1. 生成csr，同时生成秘钥对 -nodes 表示不加密秘钥，
```shell
openssl req -newkey rsa:2048 -nodes -keyout client.key -out client.csr -config ../test/data/ssl/openssl.cnf -subj "/C=CN/ST=SH/L=SH/O=system:masters/CN=dev"
```

###### 2. 生成新的Client证书
```shell
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt
```

###### 3.验证证书是否是ca签发的
```shell
openssl verify -CAfile ca.crt client.crt
```