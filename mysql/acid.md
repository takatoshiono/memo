# ACID

ACID とは信頼性のあるデータベースシステムが持つべき性質のこと。

1970年代後半、[ジム・グレイ](https://ja.wikipedia.org/wiki/%E3%82%B8%E3%83%A0%E3%83%BB%E3%82%B0%E3%83%AC%E3%82%A4) という人が提唱した(すごい人だ...)。

Wikipediaの「よく知られている業績」に書いてある

* 信頼できるトランザクション処理の要求仕様（ACIDテスト）を明確にし、それをソフトウェアに実装した。
* (describing the requirements for reliable transaction processing (memorably called the ACID test) and implementing them in software.)

これのことかな。

## Jim Gray の著書

### [Transaction Processing: Concepts and Techniques (The Morgan Kaufmann Series in Data Management Systems)](http://www.amazon.co.jp/Transaction-Processing-Concepts-Techniques-Management/dp/1558601902/ref=sr_1_3?s=english-books&ie=UTF8&qid=1435875006&sr=1-3&keywords=jim+gray)

![image](https://cloud.githubusercontent.com/assets/10000/8489074/dc600aea-2152-11e5-92ef-f89333320aaa.png)

### [トランザクション処理 上](http://www.amazon.co.jp/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E5%87%A6%E7%90%86-%E4%B8%8A-%E3%82%B0%E3%83%AC%E3%82%A4-%E3%82%B8%E3%83%A0/dp/4822281027/ref=sr_1_1?s=books&ie=UTF8&qid=1435875165&sr=1-1&keywords=%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E5%87%A6%E7%90%86)

![image](https://cloud.githubusercontent.com/assets/10000/8489106/26c7ae4e-2153-11e5-8142-cff81ca6a7e1.png)

### [トランザクション処理 下](http://www.amazon.co.jp/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E5%87%A6%E7%90%86-%E4%B8%8B-%E3%82%B0%E3%83%AC%E3%82%A4-%E3%82%B8%E3%83%A0/dp/4822281035/ref=asap_bc?ie=UTF8)

![image](https://cloud.githubusercontent.com/assets/10000/8489111/37c9f3a0-2153-11e5-8ab9-a9e3838a1c18.png)

## ACID in MySQL

MySQL では主にInnoDBがACIDに準拠している。

### 原子性(Atomicity)

トランザクションが原子的な作業単位で、それはコミットされるかロールバックされるかのどちらかである。言い換えると変更が全て適用されるか、全て元に戻るかのどちらかであるということ。

* 自動コミット設定
* COMMIT ステートメント
* ROLLBACK ステートメント
* INFORMATION_SCHEMA テーブルの運用データ(これは何のことだ?)

### 一貫性、整合性(Consistency)

常に一貫した状態を保つという性質。トランザクション中に複数のテーブルを操作する場合、クエリは古い値か新しい値のどちらかを見る。その混合にはならない。

* InnoDB [二重書き込みバッファー](http://dev.mysql.com/doc/refman/566/ja/glossary.html#glos_doublewrite_buffer)
* InnoDB [クラッシュリカバリ](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_crash_recovery)

### 分離性、独立性、隔離性(Isolation)

トランザクションはお互いから保護(分離)される。お互い干渉できず、コミットされていない変更を見ることはできない。この性質はロックメカニズムによって実現される。

形式的にはトランザクション履歴が直列化されていることと言える。この性質と性能はトレードオフの関係にあるため一部を緩和して実装されることが多い(Wikipediaより)。

* 自動コミット設定
* SET ISOLATION LEVEL ステートメント
* InnoDB [ロック](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_locking)の低レベルの詳細。これらの詳細は、パフォーマンスチューニング時に INFORMATION_SCHEMA テーブルから参照します(どういうことだ?)

### 持続性、継続性(Durability)

トランザクションの結果(保存した内容)が失われないということ。これにはハードウェアが関係してきてなかなか難しい。多くの実装ではトランザクション操作を永続記憶装置にログとして記録して、なにか異常が発生したらそれを使って元に戻すことで持続性を実現している(Wikipediaより)。InnoDBではdoublewrite bufferが持続性をサポートしている。

以下を見るといろんなレイヤーで持続性を担保しているんだな...と思った。

* innodb_doublewrite 構成オプションでオンとオフが切り替えられる InnoDB の[二重書き込みバッファー](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_doublewrite_buffer)
* innodb_flush_log_at_trx_commit 構成オプション
* sync_binlog 構成オプション
* innodb_file_per_table 構成オプション
* ストレージデバイス内の書き込みバッファー (ディスクドライブ、SSD、RAID アレイなど)
* ストレージデバイス内のバッテリーでバックアップされるキャッシュ
* MySQL を実行する際に使用されるオペレーティングシステム (特に、fsync() システムコールでのサポート)
* MySQL サーバーを実行し、MySQL データを格納するすべてのコンピュータサーバーおよびストレージデバイスへの電力を保護する無停電電源装置 (UPS)
* バックアップ方針 (頻度、バックアップのタイプ、バックアップの保存期間など)
* 分散型またはホスト型のデータアプリケーションの場合、MySQL サーバー用のハードウェアが配置されているデータセンター、およびデータセンター間のネットワーク接続の特定の特性

## 参考文献

* [MySQL :: MySQL 5.6 リファレンスマニュアル :: 14.2.1 MySQL および ACID モデル](http://dev.mysql.com/doc/refman/5.6/ja/mysql-acid.html)
* [MySQL :: MySQL 5.6 リファレンスマニュアル :: MySQL 用語集(ACID)](http://dev.mysql.com/doc/refman/5.6/ja/glossary.html#glos_acid)

