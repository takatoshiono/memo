# Set up a local cluster

クラスタ構成で動かしてみたい。このドキュメントを見ながらやる

[etcd/local\_cluster\.md at master · etcd\-io/etcd](https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/local_cluster.md)

## Build

[前回](../running-etcd/README.md)はDockerで動かしたけどProcfileでmulti-member clusterを起動するのにMacの方が便利なので、まずはetcdを使えるようにしたくて、ビルドすることにした。やり方はこのドキュメントに書いてある

[etcd/dl\_build\.md at master · etcd\-io/etcd](https://github.com/etcd-io/etcd/blob/master/Documentation/dl_build.md#build-the-latest-version)

と言ってもgit cloneして`./build`するだけでOK。`etcdctl`は`go get go.etcd.io/etcd/etcdctl`したものを使う

## Run

`goreman`というforemanのgo版みたいなやつを使って3つetcdを起動して、一つ止めて値をputして、止めたetcdを再起動した後にそのetcdに対して値をgetしたときにちゃんと値を取れる、というのを試す感じ。ドキュメント通りにやるだけ

## etcdctlのオプション

### endpoints

`--endpoints`を指定できる。

```
      --endpoints=[127.0.0.1:2379]		gRPC endpoints
```

e.g.
```
etcdctl --endpoints=localhost:22379 get key
```

### write-out

出力フォーマットを変更できる。`protobuf`があるのがおもしろい

```
    -w, --write-out="simple"			set the output format (fields, json, protobuf, simple, table)
```

e.g:

#### protobuf
バイナリだ

```
etcd (master) $ etcdctl --write-out protobuf member list

�ϊ�ܳ���伲��� �伲���infra1http://127.0.0.1:12380"http://127.0.0.1:2379Ƃ��ޑinfra2http://127.0.0.1:22380"http://127.0.0.1:22379Ȝ���infra3http://127.0.0.1:32380"http://127.0.0.1:32379
```

Unmarshalするとこうなる

```
$ cd dump-proto
$ go run main.go ../member-list-format-protobuf
{Header:cluster_id:17237436991929493444 member_id:9372538179322589801 raft_term:4  Members:[ID:9372538179322589801 name:"infra1" peerURLs:"http://127.0.0.1:12380" clientURLs:"http://127.0.0.1:2379"  ID:10501334649042878790 name:"infra2" peerURLs:"http://127.0.0.1:22380" clientURLs:"http://127.0.0.1:22379"  ID:18249187646912138824 name:"infra3" peerURLs:"http://127.0.0.1:32380" clientURLs:"http://127.0.0.1:32379" ]}
```

#### simple
```
etcd (master) $ etcdctl member list
8211f1d0f64f3269, started, infra1, http://127.0.0.1:12380, http://127.0.0.1:2379
91bc3c398fb3c146, started, infra2, http://127.0.0.1:22380, http://127.0.0.1:22379
fd422379fda50e48, started, infra3, http://127.0.0.1:32380, http://127.0.0.1:32379
```

#### fields
```
etcd (master) $ etcdctl --write-out fields  member list
"ClusterID" : 17237436991929493444
"MemberID" : 9372538179322589801
"Revision" : 0
"RaftTerm" : 4
"ID" : 9372538179322589801
"Name" : "infra1"
"PeerURL" : "http://127.0.0.1:12380"
"ClientURL" : "http://127.0.0.1:2379"

"ID" : 10501334649042878790
"Name" : "infra2"
"PeerURL" : "http://127.0.0.1:22380"
"ClientURL" : "http://127.0.0.1:22379"

"ID" : 18249187646912138824
"Name" : "infra3"
"PeerURL" : "http://127.0.0.1:32380"
"ClientURL" : "http://127.0.0.1:32379"
```

#### json
```
etcd (master) $ etcdctl --write-out json  member list
{"header":{"cluster_id":17237436991929493444,"member_id":9372538179322589801,"raft_term":4},"members":[{"ID":9372538179322589801,"name":"infra1","peerURLs":["http://127.0.0.1:12380"],"clientURLs":["http://127.0.0.1:2379"]},{"ID":10501334649042878790,"name":"infra2","peerURLs":["http://127.0.0.1:22380"],"clientURLs":["http://127.0.0.1:22379"]},{"ID":18249187646912138824,"name":"infra3","peerURLs":["http://127.0.0.1:32380"],"clientURLs":["http://127.0.0.1:32379"]}]}
```

#### table
```
etcd (master) $ etcdctl --write-out table  member list
+------------------+---------+--------+------------------------+------------------------+
|        ID        | STATUS  |  NAME  |       PEER ADDRS       |      CLIENT ADDRS      |
+------------------+---------+--------+------------------------+------------------------+
| 8211f1d0f64f3269 | started | infra1 | http://127.0.0.1:12380 |  http://127.0.0.1:2379 |
| 91bc3c398fb3c146 | started | infra2 | http://127.0.0.1:22380 | http://127.0.0.1:22379 |
| fd422379fda50e48 | started | infra3 | http://127.0.0.1:32380 | http://127.0.0.1:32379 |
+------------------+---------+--------+------------------------+------------------------+
```
