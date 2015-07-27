# MVCC

「マルチバージョン並列性制御 (MultiVersion Concurrency Control)」の略語。

変更された行の古い情報を保持するための仕組み。情報はテーブルスペース内の[ロールバックセグメント](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_rollback_segment)に格納される。この情報は以下のようなときに使用される。

* トランザクションのロールバック
* 一貫性読み取りのために行の初期バージョンを構築する際

[MVCC](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_mvcc)があると、他のトランザクションが更新している行を参照して、更新が行われる前の値を確認できる。これは他のトランザクションが保持しているロックを待機することなくクエリーを進行できるので、並列性を高める強力な手法となる。

InnoDB は内部的に3つのフィールドを各行に追加する。

![mvcc](https://cloud.githubusercontent.com/assets/10000/8896030/f22494d8-3429-11e5-9a54-24fba2df01eb.png)

Undo ログは2つに分割できる。

* 挿入 Undo ログ
  * トランザクションロールバックのみで必要
  * トランザクションのコミット直後に破棄できる
* 更新 Undo ログ
  * スナップショットが割り当てられたトランザクションが存在しなくなったあとで破棄できる

InnoDB では SQL で行を削除してもすぐにデータベースから物理的に削除されるわけではない。削除用に書き込まれた Undo ログが破棄されたときにのみ、対応する行とインデックスレコードを物理的に削除する。このような削除操作は[パージ](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_purge)と呼ばれる。非常に高速。通常は SQL と同じ時系列順で削除される。

## 参考文献

* [MySQL :: MySQL 5.6 リファレンスマニュアル :: 14.2.12 InnoDB マルチバージョン](http://dev.mysql.com/doc/refman/5.6/ja/innodb-multi-versioning.html)

