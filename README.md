# invite-code-go
一个Golang实现的邀请码生成器，根据ID生成对应邀请码，连续的ID生成的邀请码不连续，支持通过邀请码反推ID

- CHARSET为移除易混淆的字母数字后的57个字符
```text
23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
```
- 生成邀请码长度为6时，ID支持范围是 `0-34296447248`
```text
ID:0 code:999999
ID:1 code:7pMEF7
ID:2 code:FqLQMF
ID:3 code:EycjQE
ID:4 code:MrGcLM
ID:5 code:pbryqp
```
- 生成邀请码长度为7时，ID支持范围是 `0-1954897493192`
```text
ID:0 code:9999999
ID:1 code:7FMQ7Ep
ID:2 code:FMLcFQq
ID:3 code:EQc5Ejy
ID:4 code:MLG4Mcr
ID:5 code:pqrkpyb
```
- 生成邀请码长度为8时，ID支持范围是 `0-111429157112000`
```text
ID:0 code:99999999
ID:1 code:7EQ7MdFp
ID:2 code:FQcFL3Mq
ID:3 code:Ej5EcHQy
ID:4 code:Mc4MGnLr
ID:5 code:pykprgqb
```