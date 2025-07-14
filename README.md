# gogogadget

# Local Development

## Running goleaser snapshot builds
You can use goreleaser to build the snapshot binaries and docker images so you can run the server locally

```bash
goreleaser release --snapshot --clean
```

> [!TIP]
> If you're using podman with goreleaser, you'll need to create symlink to docker in your `$PATH` because the OSS version of goreleaser doesn't support podman
> Please see

Here's an example output
```bash
goreleaser release --snapshot --clean
  • skipping announce, publish, and validate...
  • cleaning distribution directory
  • loading environment variables
  • getting and validating git state
    • ignoring errors because this is a snapshot     error=git doesn't contain any tags - either add a tag or use --snapshot
    • git state                                      commit=653e9bceff0f465810844b951ab33e90d9d05865 branch=goreleaser-fl current_tag=v0.0.0 previous_tag=<unknown> dirty=false
    • pipe skipped or partially skipped              reason=disabled during snapshot mode
  • parsing tag
  • setting defaults
  • snapshotting
    • building snapshot...                           version=0.0.0-SNAPSHOT-653e9bc
  • running before hooks
    • running                                        hook=go mod tidy
  • ensuring distribution directory
  • setting up metadata
  • writing release metadata
  • loading go mod information
  • build prerequisites
  • building binaries
    • building                                       binary=dist/gogogadgetrepo_darwin_arm64_v8.0/gogogadget
    • building                                       binary=dist/gogogadgetrepo_linux_amd64_v1/gogogadget
    • building                                       binary=dist/gogogadgetrepo_darwin_amd64_v1/gogogadget
    • building                                       binary=dist/gogogadgetrepo_linux_arm64_v8.0/gogogadget
  • archives
    • archiving                                      name=dist/gogogadgetrepo_0.0.0-SNAPSHOT-653e9bc_darwin_arm64.tar.gz
    • archiving                                      name=dist/gogogadgetrepo_0.0.0-SNAPSHOT-653e9bc_linux_arm64.tar.gz
    • archiving                                      name=dist/gogogadgetrepo_0.0.0-SNAPSHOT-653e9bc_linux_amd64.tar.gz
    • archiving                                      name=dist/gogogadgetrepo_0.0.0-SNAPSHOT-653e9bc_darwin_amd64.tar.gz
  • calculating checksums
  • docker images
    • building docker image                          image=ghcr.io/filthystinkingcasual/gogogadgetrepo:0.0.0-SNAPSHOT-653e9bc-arm64
    • building docker image                          image=ghcr.io/filthystinkingcasual/gogogadgetrepo:0.0.0-SNAPSHOT-653e9bc-amd64
  • writing artifacts metadata
  • release succeeded after 6s
  • thanks for using GoReleaser!
```

Notice the binaries are all built and saved under the `dist` directory

You should see the images on you local machine now via docker/podman

```bash
docker images | grep gogo
ghcr.io/filthystinkingcasual/gogogadgetrepo                                          0.0.0-SNAPSHOT-653e9bc-amd64  7e389feae445  2 minutes ago  283 MB
ghcr.io/filthystinkingcasual/gogogadgetrepo                                          0.0.0-SNAPSHOT-653e9bc-arm64  28ce6b10ded3  2 minutes ago  279 MB
```

## Using podman with goreleaser
According to the [goreleaser docs](https://goreleaser.com/customization/docker/#customization), Podman is a GoReleaser Pro feature and is only available on Linux.

As a workaround, I simply created a symlink from the podman binary to docker that resides in my `$PATH`

First determine where podman was installed

```bash
which podman
/opt/podman/bin/podman
```

Next we'll create the symlink to docker within that same path to make goreleaser happy
```bash
ln -s /opt/podman/bin/podman /opt/podman/bin/docker
```

It should look similar to this
```bash
ls -l /opt/podman/bin
total 108780
lrwxr-xr-x 1 root wheel       22 Jul  8 11:10 docker -> /opt/podman/bin/podman
-rwxr-xr-x 1 root wheel 22730512 Feb 11 13:02 gvproxy
-r-xr-xr-x 1 root wheel  1397632 Feb 11 13:02 krunkit
-rwxr-xr-x 1 root wheel 45489072 Feb 11 13:02 podman
-rwxr-xr-x 1 root wheel  5866048 Feb 11 13:02 podman-mac-helper
-rwxr-xr-x 1 root wheel 35897344 Feb 11 13:02 vfkit
```