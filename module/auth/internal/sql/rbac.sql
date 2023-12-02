create table role
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(50) not null comment '角色英文名',
    cname       varchar(50) not null comment '角色中文名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '角色';

create table role_permission
(
    id            int primary key auto_increment comment '主键ID',
    role_id       int       not null comment '角色',
    permission_id int       not null comment '权限',
    update_time   timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time   timestamp not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '角色-权限';

create table permission
(
    id          int primary key auto_increment comment '主键ID',
    type        tinyint     not null default 0 comment '类型',
    name        varchar(64) not null default '' comment '权限名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '权限表';



