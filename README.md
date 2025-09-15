# GI

A simple HTTP service that generates .gitignore files.

## Usage

```bash
# get default .gitignore
curl gi.caml.cc > .gitignore

# list available templates  
curl gi.caml.cc/list

# get specific template
curl gi.caml.cc/go > .gitignore

# combine templates
curl gi.caml.cc/go,macos > .gitignore
```

## Development

```bash
make full
```
