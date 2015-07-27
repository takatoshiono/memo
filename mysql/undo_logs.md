# Undo ログ

Undo ログというのは、トランザクションが変更しようとしている行の変更前のデータを保持しておくデータ領域のこと。これによって非ロック一貫性読み取りが可能となる。そのあたりは[ここ](https://github.com/takatoshiono/memo/blob/master/mysql/mvcc.md)に書いた。

通常、Undo ログはシステムテーブルスペース内に保存されるが、MySQL 5.6.3 からは undo テーブルスペースという場所を独立して作ることができるらしい。マニュアルには以下のように書いてあった。

* InnoDB は 128 の undo ログをサポートしている
* それぞれの undo ログは 1023 のトランザクションをサポートできる(リードオンリーは含まない)
* つまり合計で 128,000 のトランザクションをサポートできる

詳しくは以下を参照。

* [MySQL :: MySQL 5.6 リファレンスマニュアル :: 14.5.6 個別のテーブルスペースへの InnoDB Undo ログの格納](http://dev.mysql.com/doc/refman/5.6/ja/innodb-undo-tablespace.html) 

## 参考文献

* [MySQL :: MySQL 5.6 Reference Manual :: 14.2.5 InnoDB Undo Logs](http://dev.mysql.com/doc/refman/5.6/en/innodb-undo-logs.html)

