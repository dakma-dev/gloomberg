# Listings

As listings are usually not on-chain, we need to get them from the marketplace APIs.

## implemented sources for listings

- OpenSea (StreamAPI/WS)

### Todo

- get Blur listings (API?!)
- get X2Y2 listings (API & key available already)

## distribution

*see [pubsub docs](pubsub.md) for details*

listings can be pushed to a [redis pubsub](https://redis.io/docs/manual/pubsub/) channel, so that possible other instances can subscribe to it and get the listings without having to query the APIs themselves.

## storage

Currently, we do not store listings in the database, but we might want to do so in the future.
