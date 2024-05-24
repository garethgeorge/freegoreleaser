# FreeGoReleaser is a FOSS fork of GoReleaser 

FreeGoReleaser is a fully open source fork of GoReleaser aiming to re-introduce some paywalled (but simple) features that are:

 * Low effort to maintain
 * But enable high value applications

It's certainly the case that not all (pro) functionality of the upstream goreleaser project will ever be in scope for FreeGoReleaser. 

This project will focus on re-introducing scriptability and extensibility hooks e.g. running commands at various points in the release process such that a GoReleaser user can customize their build arbitrarily.

## Features

 * **Scriptability**: Archive hooks

## Development

### Pulling in upstream changes

After pulling in upstream changes, any references to goreleaser's upstream package must be renamed e.g. run 

```
git remote add upstream github.com/goreleaser/goreleaser
git remote add origin github.com/garethgeorge/freegoreleaser
git pull upstream main
find . -type f -exec sed -i 's|github.com/goreleaser/goreleaser|github.com/freegoreleaser/goreleaser|g' {} \;
```
