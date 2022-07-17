# gloomberg

## docker
- get your wallet address or ens
  ***yourDegenWallet.eth*** or ***0x9e7DC5307940fa170F9093Ca548bDa0EDB602767***
- get an account at [Infura](https://infura.io)/[Alchemy](https://www.alchemy.com)/*whatever* to get a websockets endpoint to an ethereum node
  ***wss://mainnet.infura.io/ws/v3/32e98f6ffb81456df24087ab5b***

```bash
docker run --rm -it ghcr.io/benleb/gloomberg:9c1abc0-amd64 live --endpoints="wss://mainnet.infura.io/ws/v3/32e98f6ffb81456df24087ab5b" --wallets="yourDegenWallet.eth"
```
