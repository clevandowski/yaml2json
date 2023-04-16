# yaml2json

Transform yaml input stream into json.

Support yaml with multi-documents input stream.

## Build

Requires go >= v1.20

```bash
go build
```

## Install

```bash
sudo ln -s ${PWD}/yaml2json /usr/local/bin/yaml2json
```

## Usage

```bash
cat myfile.yaml | yaml2json
```

```bash
echo 'test: 3' | yaml2json
```
