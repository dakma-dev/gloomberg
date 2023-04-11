# gloomberg

⚠️ this is a wip poc (also the docs!) so I cant provide any support, sorry! you have to know how to use this code yourself for now ⚠️



## random screenshots

<p><img title="screenshot 22/09/26" width="48%" align="center" alt="image" src="https://user-images.githubusercontent.com/512997/192209465-edb95aeb-be2f-419f-b4d1-84da3f35e97e.png">
<img title="screenshot 22/09/26" align="center" width="48%" alt="image" src="https://user-images.githubusercontent.com/512997/192209940-538105c5-8552-42ed-bf16-508da83611d0.png"></p>

---

## requirements

- get your wallet address or ens
    **_yourDegenWallet.eth_** or **_0x9e7DC5307940fa170F9093Ca548bDa0EDB602767_**
- get an account at [Infura](https://infura.io)/[Alchemy](https://www.alchemy.com)/_whatever_ to get a websockets endpoint to an ethereum node
    **_wss://mainnet.infura.io/ws/v3/32e98f6ffb81456df24087ab5b_**

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

## lfg! or _getting started_

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

## development (wip)

_see the [(development) docs](<[docs/development.md](https://benleb.github.io/gloomberg/)>) for further information_

issues closed, PRs open (ping me if you want to contribute)

### pre-commit

we use [pre-commit](https://pre-commit.com) to run some checks before committing. install like described in the [docs](https://pre-commit.com/#install)

see the [pre-commit config](.pre-commit-config.yaml) for the checks. you can also run them manually on your machine with `pre-commit run --all-files`. for more information see the [docs](https://pre-commit.com/#usage)

### GitHub actions

we use similar checks to the pre-commit ones in the [golangci-lint](.github/workflows/golangci-lint.yml) workflow to also check this on the "server side" and do it for pull requests. you can also run them manually on your machine with `golangci-lint run`. for more information see the [docs](https://golangci-lint.run/usage/install/#local-installation)

## docs

we use [mkdocs material (insiders)](https://squidfunk.github.io/mkdocs-material/) (based on [mkdocs.org](https://www.mkdocs.org)) for documentation. install like described in the [mkdocs material docs](https://squidfunk.github.io/mkdocs-material/getting-started/)

the docs will be automatically be built and deployed to [benleb.github.io/gloomberg/](https://benleb.github.io/gloomberg/) via the [docs workflow](.github/workflows/mkdocs.yml) (on every push/merge to the `main` branch)

### contribute

#### local preview

start a local server to preview the docs at [http://localhost:8000/](http://localhost:8000/)

```shell
mkdocs serve
```

#### build

build the docs to the `site` folder

```shell
mkdocs build
```
