# gloomberg

## commands

* `live` - Create a new project.

### commands after split up (WIP)

-> *feel free to propose better names*

* `bridge` - bridge events (sales, listings, ...) from the OpenSea API & Ethereum nodes to the redis database.
  * receives events from the [OpenSea API](https://docs.opensea.io) and pushes them to the [redis](data/redis.md) database.
  * receives logs from the websockets API of ethereum nodes ([geth](https://geth.ethereum.org/) for example) and pushes them to the [redis](data/redis.md) database.
* `watch` - like `live` or the "reverse"/"other side" of the `bridge` command.
  * watches the redis database for new events and logs and prints them to the terminal

## Project layout

    cmd/          # Commands for gloomberg
        live/     # ...
        os2rdb/   # Receives events from the OpenSea API and pushes them to the redis database.
        node2rdb/ # Receives logs from the websockets API of ethereum nodes and pushes them to the redis database.
    ddt/          # Newer implementation of most stuff
    docs/
        index.md  # The documentation homepage.
        ...       # Other markdown pages, images and other files.
