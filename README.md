# `nostr-json`

Generate a `nostr.json` used for verification of a Nostr user

> **Note**: Should be placed at `/.well-known/nostr.json` on the server you are verifying the name for.

Based on code found in the [nip19](https://github.com/nbd-wtf/go-nostr/tree/master/nip19) package in [github.com/nbd-wtf/go-nostr](https://github.com/nbd-wtf/go-nostr).

## Installation

```sh
go install github.com/peterhellberg/nostr-json@latest
```

## Usage

Call `nostr-json` with the **name** and **npub** you want to generate a JSON document for.

```sh
nostr-json -name peter -npub npub1lzkvkjqs8amrd9m0dmhf4l3wdf7tjawyslrcm96zgttrxzfx9kcsn0dz7s
```

```json
{
  "names": {
    "peter": "f8accb48103f7636976f6eee9afe2e6a7cb975c487c78d974242d63309262db1"
  }
}
```

> **Note**: You can also run the tool without first installing it:
> `go run github.com/peterhellberg/nostr-json@v1.0.0 -name peter -npub npub1lzkvkjqs8amrd9m0dmhf4l3wdf7tjawyslrcm96zgttrxzfx9kcsn0dz7s`

## Links

 - [NIP-19](https://github.com/nostr-protocol/nips/blob/master/19.md)

## License (MIT)

Copyright (c) 2023 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
