# BUID (Brijesh's Unique Identifier) Specification v1.0

<!--toc:start-->

- [Introduction](#introduction)
- [Who Should Use This](#who-should-use-this)
- [Structure](#structure)
  - [Components](#components)
- [Example](#example)
- [Comparison with UUID](#comparison-with-uuid)
- [Shortcomings](#shortcomings)
- [Version](#version)
<!--toc:end-->

## Introduction

I created BUID for fun and to address the need for a compact, human-readable, and sortable unique identifier.

## Who Should Use This

BUID is ideal for systems that require:

- Compact unique identifiers (shorter than UUID)
- Rough chronological sorting capability
- URL-safe and human-readable format

## Structure

- Total length: 18 characters
- Encoding: Base32 (A-Z and 2-7)

### Components

1. **Timestamp** (first 7 characters):
2. **Random component** (next 11 characters):

## Example

`BTMSEM5RXDBPOQLHXC`

## Comparison with UUID

| Feature                 | BUID              | UUID (v4)              |
| ----------------------- | ----------------- | ---------------------- |
| Length                  | 18 characters     | 36 characters          |
| Character set           | Base32 (A-Z, 2-7) | Hexadecimal (0-9, A-F) |
| Sortability             | Sortable\*        | Not sortable           |
| Randomness              | 2^55 / second     | 2^128                  |
| Double-click selectable | Yes               | No                     |

## Shortcomings

- The timestamp is generated using the local system time, converted to UTC. Users should be aware that strict chronological ordering of BUIDs across different systems is not guaranteed.
- As BUID is only 18 characters long, it's randomness will fall short if usage reaches over hundred thousand IDs generated per second.

## Version

Specification Version: 1.0
