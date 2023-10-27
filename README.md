# Generate ED25519 Keypairs for Tendermint

- Generate a random private key and derive a public key
- Generate a public key (Tendermint node ID) from a base64 encoded private key input

Build:

```sh
go build -o keygen
```

## Public Key (Node ID)

Generate public key from Base 64 encoded private key (e.g. `priv_key` value stored in `config/node_key.json`:

```bash
./keygen --secret=<base64 encoded private key>
```

Key the raw node address (assumes `jq` is installed:

```bash
./keygen --secret=<base 64 encoded secret> | jq -r '.address'
```

## Generate Keypair

```bash
./keygen
```

## Checks

Check what you are generating by passing the `-v` flag for verbosity.
