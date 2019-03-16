# Grqphql API

this is an API for BotNet and the cli

## create keys

```bash
mkdir keys
cd keys
```

```bash
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
```

```bash
openssl rsa -pubout -in private_key.pem -out public_key.pem
```
