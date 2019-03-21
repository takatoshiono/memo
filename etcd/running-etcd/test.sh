#!/bin/sh

set -x

docker exec etcd-gcr-v3.3.12 /bin/sh -c "/usr/local/bin/etcd --version"
docker exec etcd-gcr-v3.3.12 /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl version"
docker exec etcd-gcr-v3.3.12 /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl endpoint health"
docker exec etcd-gcr-v3.3.12 /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl put foo bar"
docker exec etcd-gcr-v3.3.12 /bin/sh -c "ETCDCTL_API=3 /usr/local/bin/etcdctl get foo"
