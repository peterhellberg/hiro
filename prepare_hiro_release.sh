#!/usr/bin/env bash

if [ "$#" -ne 1 ]
then
  echo "Usage: ./prepare_hiro_release.sh <version number>"
  exit 1
fi

function createRelease {
  release_name="$RELEASE_BASE-$1-$2"
  release_dir="$RELEASES_DIR/$release_name"

  # Create the release dir
  mkdir -p $release_dir

  # Cross compile for the given OS/Arch
  GOOS=$1 GOARCH=$2 go build -o $release_dir/$3

  # Compress the binary (junk paths, quiet)
  zip -j -q -r $release_dir.zip $release_dir

  # Remove the binary
  rm $release_dir/$3

  # Remove the release dir
  rmdir $release_dir
}

RELEASES_DIR="/tmp/hiro/releases/$1"
RELEASE_BASE="hiro-$1"

# Architectures to create releases for

WINDOWS_386=(
  windows
  386
  hiro.exe
)

WINDOWS_amd64=(
  windows
  amd64
  hiro.exe
)

DARWIN_amd64=(
  darwin
  amd64
  hiro
)

LINUX_amd64=(
  linux
  amd64
  hiro
)

# Create the releases
createRelease ${WINDOWS_386[@]}
createRelease ${WINDOWS_amd64[@]}
createRelease ${DARWIN_amd64[@]}
createRelease ${LINUX_amd64[@]}

# Open the releases dir
open $RELEASES_DIR
