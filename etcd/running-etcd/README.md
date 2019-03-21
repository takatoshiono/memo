# Running etcd

とりあえずetcdを動かしたい

https://github.com/etcd-io/etcd のREADMEにはpre-built release binaryを使うのが手っ取り早くて、その方法は[Releases · etcd\-io/etcd](https://github.com/etcd-io/etcd/releases)に書いてあるということなので、ここに書いてあるDockerを使って動かすことにする

## Run

`--rm`オプションを付け加えた。理由は2回目の起動時にコンテナイメージの削除に失敗するため

```
$ ./run.sh
+ rm -rf /tmp/etcd-data.tmp
+ mkdir -p /tmp/etcd-data.tmp
+ docker rmi gcr.io/etcd-development/etcd:v3.3.12
Untagged: gcr.io/etcd-development/etcd:v3.3.12
Untagged: gcr.io/etcd-development/etcd@sha256:21012df620d0ed6b4bf63ab5d1e98f81e2e6e4f83ffff620d4cd4dae31cd857b
Deleted: sha256:28c771d7cfbf436cc2471523350d58a75a4c28a7e8684b1dd54b7e8ba321f84b
Deleted: sha256:ec4d742bbe65cbef1d74fd1e4af527005dd1efb1d3d9de6f21a5a629c065529d
Deleted: sha256:1471a73dedaa2e0ac192aabaa736928dbdf19845e6b9eca1346917769f053667
Deleted: sha256:63538dd40a36110915bd1ea63df21cecce20c73cf2a13a77c3cc3274a9713f0e
Deleted: sha256:8a4def463b2ce6ec19dbc912590e5167e926f19b00ce7146960cef21a3aced7d
Deleted: sha256:d4d0872de6dd842ebc822354c3b22845af6eb7444d181636f1330f5da7e86e32
Deleted: sha256:503e53e365f34399c4d58d8f4e23c161106cfbce4400e3d0a0357967bad69390
+ docker run --rm -p 2379:2379 -p 2380:2380 --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data --name etcd-gcr-v3.3.12 gcr.io/etcd-development/etcd:v3.3.12 /usr/local/bin/etcd --name s1 --data-dir /etcd-data --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380 --initial-cluster s1=http://0.0.0.0:2380 --initial-cluster-token tkn --initial-cluster-state new
Unable to find image 'gcr.io/etcd-development/etcd:v3.3.12' locally
v3.3.12: Pulling from etcd-development/etcd
6c40cc604d8e: Pull complete
0c0f49f488a2: Pull complete
2a8faafca7e5: Pull complete
a29a1dc564f3: Pull complete
0772095a9172: Pull complete
b7439762be07: Pull complete
Digest: sha256:21012df620d0ed6b4bf63ab5d1e98f81e2e6e4f83ffff620d4cd4dae31cd857b
Status: Downloaded newer image for gcr.io/etcd-development/etcd:v3.3.12
2019-03-21 15:16:43.431430 I | etcdmain: etcd Version: 3.3.12
2019-03-21 15:16:43.431519 I | etcdmain: Git SHA: d57e8b8
2019-03-21 15:16:43.431531 I | etcdmain: Go Version: go1.10.8
2019-03-21 15:16:43.431543 I | etcdmain: Go OS/Arch: linux/amd64
2019-03-21 15:16:43.431554 I | etcdmain: setting maximum number of CPUs to 2, total number of available CPUs is 2
2019-03-21 15:16:43.435064 I | embed: listening for peers on http://0.0.0.0:2380
2019-03-21 15:16:43.435134 I | embed: listening for client requests on 0.0.0.0:2379
2019-03-21 15:16:43.448347 I | etcdserver: name = s1
2019-03-21 15:16:43.448382 I | etcdserver: data dir = /etcd-data
2019-03-21 15:16:43.448433 I | etcdserver: member dir = /etcd-data/member
2019-03-21 15:16:43.448452 I | etcdserver: heartbeat = 100ms
2019-03-21 15:16:43.448469 I | etcdserver: election = 1000ms
2019-03-21 15:16:43.448486 I | etcdserver: snapshot count = 100000
2019-03-21 15:16:43.448629 I | etcdserver: advertise client URLs = http://0.0.0.0:2379
2019-03-21 15:16:43.448683 I | etcdserver: initial advertise peer URLs = http://0.0.0.0:2380
2019-03-21 15:16:43.448713 I | etcdserver: initial cluster = s1=http://0.0.0.0:2380
2019-03-21 15:16:43.461137 I | etcdserver: starting member 59a9c584ea2c3f35 in cluster f9f44c4ba0e96dd8
2019-03-21 15:16:43.461185 I | raft: 59a9c584ea2c3f35 became follower at term 0
2019-03-21 15:16:43.461207 I | raft: newRaft 59a9c584ea2c3f35 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
2019-03-21 15:16:43.461220 I | raft: 59a9c584ea2c3f35 became follower at term 1
2019-03-21 15:16:43.469235 W | auth: simple token is not cryptographically signed
2019-03-21 15:16:43.472572 I | etcdserver: starting server... [version: 3.3.12, cluster version: to_be_decided]
2019-03-21 15:16:43.476228 I | etcdserver: 59a9c584ea2c3f35 as single-node; fast-forwarding 9 ticks (election ticks 10)
2019-03-21 15:16:43.476545 I | etcdserver/membership: added member 59a9c584ea2c3f35 [http://0.0.0.0:2380] to cluster f9f44c4ba0e96dd8
2019-03-21 15:16:43.669442 I | raft: 59a9c584ea2c3f35 is starting a new election at term 1
2019-03-21 15:16:43.669556 I | raft: 59a9c584ea2c3f35 became candidate at term 2
2019-03-21 15:16:43.669622 I | raft: 59a9c584ea2c3f35 received MsgVoteResp from 59a9c584ea2c3f35 at term 2
2019-03-21 15:16:43.669664 I | raft: 59a9c584ea2c3f35 became leader at term 2
2019-03-21 15:16:43.669700 I | raft: raft.node: 59a9c584ea2c3f35 elected leader 59a9c584ea2c3f35 at term 2
2019-03-21 15:16:43.670417 I | etcdserver: setting up the initial cluster version to 3.3
2019-03-21 15:16:43.672800 N | etcdserver/membership: set the initial cluster version to 3.3
2019-03-21 15:16:43.672943 I | etcdserver/api: enabled capabilities for version 3.3
2019-03-21 15:16:43.673027 I | etcdserver: published {Name:s1 ClientURLs:[http://0.0.0.0:2379]} to cluster f9f44c4ba0e96dd8
2019-03-21 15:16:43.673515 I | embed: ready to serve client requests
2019-03-21 15:16:43.678319 N | embed: serving insecure client requests on [::]:2379, this is strongly discouraged!
```

## Try some commands

```
$ ./test.sh
+ docker exec etcd-gcr-v3.3.12 /bin/sh -c '/usr/local/bin/etcd --version'
etcd Version: 3.3.12
Git SHA: d57e8b8
Go Version: go1.10.8
Go OS/Arch: linux/amd64
+ docker exec etcd-gcr-v3.3.12 /bin/sh -c 'ETCDCTL_API=3 /usr/local/bin/etcdctl version'
etcdctl version: 3.3.12
API version: 3.3
+ docker exec etcd-gcr-v3.3.12 /bin/sh -c 'ETCDCTL_API=3 /usr/local/bin/etcdctl endpoint health'
127.0.0.1:2379 is healthy: successfully committed proposal: took = 1.2741ms
+ docker exec etcd-gcr-v3.3.12 /bin/sh -c 'ETCDCTL_API=3 /usr/local/bin/etcdctl put foo bar'
OK
+ docker exec etcd-gcr-v3.3.12 /bin/sh -c 'ETCDCTL_API=3 /usr/local/bin/etcdctl get foo'
foo
bar
```

## Stop

```
$ docker stop etcd-gcr-v3.3.12
```
