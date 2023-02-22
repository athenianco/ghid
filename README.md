# GHID: GitHub GraphQL Node ID tools

This project provides decoding and encoding utilities for GitHub GraphQL Node IDs.

## Why?

Ideally, ID in any kind of system must be unique. However, as the system scales, you may need to change primary keys
in your DB, and thus change the unique identifier used in your API as well. The same happened with GitHub GraphQL API.

To achieve a notion of Node IDs (a unique identifier, regardless of the object/node type), you likely don't want to
use generic IDs like UUID - it's more efficient to encode a type of the node and some kind of primary key in the ID.
This is exactly how GitHub Node IDs work: they encode some type-related information, so that the backend can quickly
figure out which table/collection to use for a lookup, and an identifier specific for that node type.

First version of GitHub Node IDs (aka IDv1) were text-based (base64-encoded text) which turned out to be inefficient.
GitHub migrated to IDv2 format, which is binary-based (base64 over msgpack encoding). However, the migration was not
handled well (in our humble opinion).

This library solves a problem: if you already have a huge dataset of cached info based on GitHub's GraphQL data,
and you want to migrate it to the new ID format in offline mode (without hitting the actual API) - we've got you covered.
This project provides a CLI and a Go library to handle these cases. As a bonus, it allows introspection of ID contents.
When your client receives an ID, you can instantly tell which type it contains, what is the commit SHA,
which repo it belongs to, etc. It can be useful as an optimization for some use cases.

*WARNING:* This library is written with some assumptions inferred from reverse-engineering a bunch of GitHub Node IDs,
which we've seen in the wild. It is NOT affiliated with GitHub. Generated output MAY be incorrect and break your system.
Always design a fallback code path that actually hits GitHub API and verify output for your specific use case.
