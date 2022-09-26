# gloomberg

*no support, sorry! you have to know how to use this code yourself. issues closed, PRs open (ping me if you want to contribute)*

## random screenshots

<p><img title="screenshot 22/09/26" width="480" alt="image" src="https://user-images.githubusercontent.com/512997/192209465-edb95aeb-be2f-419f-b4d1-84da3f35e97e.png">
<img title="screenshot 22/09/26" align="right" width="470" alt="image" src="https://user-images.githubusercontent.com/512997/192209940-538105c5-8552-42ed-bf16-508da83611d0.png"></p>

---

## requirements

- get your wallet address or ens
  ***yourDegenWallet.eth*** or ***0x9e7DC5307940fa170F9093Ca548bDa0EDB602767***
- get an account at [Infura](https://infura.io)/[Alchemy](https://www.alchemy.com)/*whatever* to get a websockets endpoint to an ethereum node
  ***wss://mainnet.infura.io/ws/v3/32e98f6ffb81456df24087ab5b***

## recommended ☝️

‼️ use [redis](https://redis.io) as cache for collection & ENS names! it saves a lot of (limited) calls to a node and is also faster than calls to a node. you can use the [redis docker image](https://hub.docker.com/_/redis) for this or [install redis](https://redis.io/docs/getting-started/installation/) on a local machine.

most simple configuration example with `redis` running on `10.0.0.2:6379` (default port, database and no password):
  
  ```yaml
  # redis cache
  redis:
    # use redis as name & sale cache
    enabled: true
    # redis host
    host: 10.0.0.2
  ```

there is also an ultra-simple built-in cache but without any persistence and therefore empty on every new start of `gloomberg`.

## lfg! or *getting started*

```bash
# get link to latest linux amd64 binary
GBL=$(curl -L -s -H 'Accept: application/json' https://github.com/benleb/gloomberg/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
# download binary and extract it to /usr/local/bin
wget -qO- https://github.com/benleb/gloomberg/releases/download/${GBL}/gloomberg_${GBL/v/}_linux_amd64.tar.gz | sudo tar -C /usr/local/bin -vzx gloomberg

# run
gloomberg live -e "wss://mainnet.infura.io/ws/v3/32e9..." -w "yourDegenWallet.eth"
```

### docker

```shell
docker run --rm -it \
  --env "GLOOMBERG_ENDPOINTS=wss://mainnet.infura.io/ws/v3/32e9..."
  --env "GLOOMBERG_WALLETS=yourDegenWallet.eth"
  ghcr.io/benleb/gloomberg:latest live
```

## gloomberg‽

The name is a homage to the famous [gloomberg professional](https://opensea.io/assets/ethereum/0x495f947276749ce646f68ac8c248420045cb7b5e/99817193321473223322497783689261477808362186321335987444674465937111627333639) artwork created by [OSF](https://osf.art)

Thanks also to all the other testers and contributors not shown in the git history! 💰 ❌ 💤
