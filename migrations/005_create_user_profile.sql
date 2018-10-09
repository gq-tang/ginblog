-- +migrate Up
drop table if exists user_profile;
create table user_profile(
    id int(10) not null auto_increment,
    realname varchar(15) default null,
    sex tinyint(1) default 1 comment '1 boy,0 girl',
    birth varchar(20) not null default '' comment '生日',
    email varchar(20) default null,
    phone varchar(11) default null,
    address varchar(255) not null default '' comment '地址',
    hobby varchar(255) not null default '' comment '爱好',
    intro text comment '介绍',
    primary key(id)
) engine=InnoDB auto_increment=2 default charset=utf8 comment='用户详情';

insert into `user_profile`(id,realname,sex,birth,email,phone,address,hobby,intro)
values ('1', '强哥', '1', '1986-02-03', 'tgq2004@163.com', '15888888888', '中国 上海', '看没有看过的电影', '加油，给自己');
-- +migrate Down
drop table user_profile;