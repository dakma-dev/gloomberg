# gloomberg

⚠️ this is a wip poc (also the docs!) so I cant provide any support, sorry! you have to know how to use this code
yourself for now ⚠️

## random screenshots

<p><img title="screenshot 22/09/26" width="48%" align="center" alt="image" src="https://user-images.githubusercontent.com/512997/192209465-edb95aeb-be2f-419f-b4d1-84da3f35e97e.png">
<img title="screenshot 22/09/26" align="center" width="48%" alt="image" src="https://user-images.githubusercontent.com/512997/192209940-538105c5-8552-42ed-bf16-508da83611d0.png"></p>

---

## requirements

-   websocket endpoint of an own ethereum node (e.g. [geth](https://geth.ethereum.org), ...) or at least one to that you
    can make as many calls as you want. Node providers (
    e.g. [Infura](https://infura.io)/[Alchemy](https://www.alchemy.com)/_whatever_) also work, but as we make a lot of
    calls to the node, you will hit the limits very fast or you will have to pay a lot of money for it.
-   a [redis(-stack)](https://redis.io) instance. you can use the [redis docker image](https://hub.docker.com/_/redis) for
    this or [install redis](https://redis.io/docs/getting-started/installation/) on a local machine. it is used as a cache
    for collection & ENS names. it saves a lot of (limited) calls to a node and is also faster than calls to a node.

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

The name is a homage to the
famous [gloomberg professional](https://opensea.io/assets/ethereum/0x495f947276749ce646f68ac8c248420045cb7b5e/99817193321473223322497783689261477808362186321335987444674465937111627333639)
artwork created by [OSF](https://osf.art)

Thanks also to all the other testers and contributors not shown in the git history! 💰 ❌ 💤

## development (wip)

_see the [(development) docs](<[docs/development.md](https://benleb.github.io/gloomberg/)>) for further information_

issues closed, PRs open (ping me if you want to contribute)

### pre-commit

we use [pre-commit](https://pre-commit.com) to run some checks before committing. install like described in
the [docs](https://pre-commit.com/#install)

see the [pre-commit config](.pre-commit-config.yaml) for the checks. you can also run them manually on your machine
with `pre-commit run --all-files`. for more information see the [docs](https://pre-commit.com/#usage)

### GitHub actions

we use similar checks to the pre-commit ones in the [golangci-lint](.github/workflows/golangci-lint.yml) workflow to
also check this on the "server side" and do it for pull requests. you can also run them manually on your machine
with `golangci-lint run`. for more information see
the [docs](https://golangci-lint.run/usage/install/#local-installation)

## docs

todo

### contribute

todo
