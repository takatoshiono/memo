# Interacting with etcd

[etcd/interacting\_v3\.md at master · etcd\-io/etcd](https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/interacting_v3.md)

## put

キーに値を書き込む(`put`)。キーは全てのetcd clusterに複製される

## get

キーから値を読む(`get`)。キーはrange指定、prefix指定も可能

## revision

etcd clusterのkey-value storeに変更があるとrevisionがインクリメントされる。過去のリビジョンを指定してキーを読むことが可能

## watch

キーを監視(`watch`)して変更があったときに知ることができる。watchするキーはrange指定、prefix指定も可能。Interactive mode(`watch -i`)もある。最新の変更だけでなく変更の履歴も欲しい時は`--rev`を指定することで受け取ることができる。Interactive modeの時に`progress`コマンドを使うと現在のリビジョンがわかる

## compact

`compact`することでリビジョンの履歴を消すことができる

## lease

etcd clusterから有効期限(time-to-live, TTL)つきのleaseをもらえる(`grant`)。leaseに複数のキーを紐づけることができて、TTLが過ぎると全部消える。TTLが来る前に消す(`revoke`)ことも可能。またleaseを生き永らえさせる(`keep-alive`)こともできる。おもしろい


