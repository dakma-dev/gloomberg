# Publish/Subscribe

we use a [redis pubsub](https://redis.io/docs/manual/pubsub/) channel to distribute data across instances.

## channels

- `rawListings` - contains all [listings](listings.md) from all sources without any modifications (as JSON strings)

### ideas for future channels

- `listings` - contains all [listings](listings.md) from all sources as `ListingEvent` objects.
  This channel could be used by the frontends to display listings.
- `rawSales` - contains all [sales](sales.md) from all sources without any modifications (as JSON strings)
- `sales` - contains all [sales](sales.md) from all sources as `SaleEvent` objects.
  This channel could be used by the frontends to display listings.
