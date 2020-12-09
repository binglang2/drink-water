drop table if exists dw_user;
create table dw_user(
                        id integer not null primary key autoincrement,
                        name text not null default '',
                        mobile text not null default '',
                        sex integer not null default 0,
                        dt_token text not null default '',
                        deleted boolean not null default false,
                        created timestamp default current_timestamp,
                        updated timestamp default current_timestamp
);

drop table if exists dw_joke;
create table dw_joke(
                        id integer not null primary key autoincrement,
                        content text not null default '',
                        status int not null default 0,
                        deleted boolean not null default false,
                        created timestamp default current_timestamp,
                        updated timestamp default current_timestamp
);
insert into dw_joke(content) values('喝水了！喝水了！举起旁边的杯子，没有水就赶紧满上，感情深，一口闷！');
insert into dw_joke(content) values('喝水了！喝水了！拿起旁边的杯子，Cheers！');
insert into dw_joke(content) values('美好的一天从喝水开始，你喝水了吗？');
insert into dw_joke(content) values('喝水！喝水！再忙也别忘了喝水哟！');
insert into dw_joke(content) values('You see see you, one day day de. 水也不喝，you 是要渴死 yourself 吗？喝水喝水。');
insert into dw_joke(content) values('别说话，喝水，多喝热水。');
insert into dw_joke(content) values("chinese: 喝水; 英语: drink water; japanese: 水を飲む; 法语: Bois de l'eau; Italian: Bere acqua; 德语: Wasser trinken");
