#!/usr/bin/bash

set -ex

## Variables
PROTO_DIR="${PROTO_DIR:-$(pwd)/proto}"
PREFIX="${PREFIX:-usr}"
COMPONENT_SUMMARY="Runs golang scripts on received webhooks"
COMPONENT_DESCRIPTION="Runs golang scripts on received webhooks"
COMPONENT_NAME="webhooked"
COMPONENT_VERSION="${COMPONENT_VERSION:-1.0}"

## Prototype area preparation
rm -rf "$PROTO_DIR"
mkdir -p "$PROTO_DIR/i386/$PREFIX/bin"
mkdir -p "$PROTO_DIR/i386/opt"

## Build
go build -o "$PROTO_DIR/i386/$PREFIX/bin/webhooked" cmd/*.go
cp -r usr "$PROTO_DIR/i386/opt/webhooked"
install -d "$PROTO_DIR/i386/lib/svc/manifest"
install -c "$PROTO_DIR/i386/lib/svc/manifest" -m 0644 -u root -g bin svc/manifest/webhooked.xml

## Packaging commands
pkgsend generate "$PROTO_DIR/i386" | pkgfmt > "$PROTO_DIR/webhooked.p5m.1"

pkgmogrify -DARCH="$(uname -p)" -DCOMPONENT_NAME="$COMPONENT_NAME" -DCOMPONENT_VERSION="$COMPONENT_VERSION" -DCOMPONENT_SUMMARY="$COMPONENT_SUMMARY" -DCOMPONENT_DESCRIPTION="$COMPONENT_DESCRIPTION" "$PROTO_DIR/webhooked.p5m.1" packaging/webhooked.mog | pkgfmt > "$PROTO_DIR/webhooked.p5m.2"

pkgdepend generate -md "$PROTO_DIR/i386" "$PROTO_DIR/webhooked.p5m.2" | pkgfmt > "$PROTO_DIR/webhooked.p5m.3"

pkgdepend resolve -m "$PROTO_DIR/webhooked.p5m.3"

pkglint "$PROTO_DIR/webhooked.p5m.3.res"

pkgrepo create "$COMPONENT_NAME-repo"

pkgrepo -s "$COMPONENT_NAME-repo" set publisher/prefix=openflowlabs

pkgsend -s "$COMPONENT_NAME-repo" publish -d "$PROTO_DIR/i386" "$PROTO_DIR/webhooked.p5m.3.res"

pkgrecv -s "$COMPONENT_NAME-repo" -a -d "$COMPONENT_NAME.p5p" "$COMPONENT_NAME"