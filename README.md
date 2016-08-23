# Memo App

## dots.女子部ハンズオンで作成しました

http://eventdots.jp/event/595084

### Vagrant設定

    vagrant up

    sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
    sudo apt-get update -y
    sudo apt-get install -y golang mysql-server

### Mysql設定

    mysql -u root

    create database memo;

    use memo;

    CREATE TABLE `memos` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `body` varchar(100) NOT NULL,
    `created_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

    GRANT select,update,insert,delete ON memo.* TO 'testuser'@'localhost' IDENTIFIED BY 'testpassword';

### 起動コマンド

    cd /vagrant

    revel run memo

### アクセス

http://localhost:9000/
