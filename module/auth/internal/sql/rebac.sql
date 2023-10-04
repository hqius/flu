create table domain
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(32) not null comment '领域名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '领域表';

create table relation
(
    id          int primary key auto_increment comment '主键ID',
    did         int         not null comment '领域ID',
    name        varchar(32) not null comment '关系名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '领域的关系表';


create table domain_relation
(
    id          int primary key auto_increment comment '主键ID',
    object      varchar(32) not null comment '目标',
    rel         varchar(32) not null comment '关系',
    user        varchar(32) not null comment '用户',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '领域与领域的关系表';

create table instance_relation
(
    id              int primary key auto_increment comment '主键ID',
    drid            int         not null comment '领域与领域关系表ID',
    target_instance varchar(64) not null comment '目标实例',
    relation        varchar(32) not null comment '关系',
    source_instance varchar(64) not null comment '源实例',
    update_time     timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time     timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '领域与领域的关系表';

create table user
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(32) not null comment '用户名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户表';

