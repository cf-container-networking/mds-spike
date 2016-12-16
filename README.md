# mds spike

## Setup
From your local machine, build the go programs and SCP them to a diego cell:

```bash
GOOS=linux go build -o bin/mds-proxy proxy/main.go
GOOS=linux go build -o bin/mds-server server/main.go

bosh-cli scp bin/mds-proxy diego-cell/32ad6c94-4828-4bcc-9711-ecda940b7c93:/usr/local/bin/
bosh-cli scp bin/mds-server diego-cell/32ad6c94-4828-4bcc-9711-ecda940b7c93:/usr/local/bin/
bosh-cli scp netns-exec.c diego-cell/32ad6c94-4828-4bcc-9711-ecda940b7c93:/tmp
```

All subsequent commands are from inside a Diego cell:
```bash
bosh-cli ssh diego-cell/32ad6c94-4828-4bcc-9711-ecda940b7c93
```

On the cell, build the `netns-exec` binary so we can run a process in the network namespace:
```bash
gcc -o /usr/local/bin/netns-exec /tmp/netns-exec.c
```

## Find a container to target
```bash
ls /var/vcap/data/garden/depot/
## be979867-6f90-4cb6-74c0-86de31ad5865

export CONTAINER_ID=be979867-6f90-4cb6-74c0-86de31ad5865
```

## Running
In one window, start the server, listening on a local socket specific for that container:
```bash
mds-server /var/vcap/data/mds/mds-$CONTAINER_ID.sock
```

In another bosh SSH session, start the proxy inside the container's network namespace:
```bash
export CONTAINER_ID=be979867-6f90-4cb6-74c0-86de31ad5865

# discover the pid of a process in the container
export NETNS_PID=$(cat /var/vcap/data/garden/depot/$CONTAINER_ID/pidfile)

# start the proxy listening in the container netns
netns-exec /proc/$NETNS_PID/ns/net mds-proxy /var/vcap/data/mds/mds-$CONTAINER_ID.sock
```

Finally, we can connect to the `mds` from within the container:

```bash
cf ssh myapp
nc localhost 5000
## hello
```
