create table `t_like_count`
(
    `id`            bigint(20) unsigned not null auto_increment comment '主键ID',
    `video_id`      bigint(20) unsigned not null default 0 comment '视频id',
    `like_count`    bigint(20)          not null default 0 comment '点赞计数',
    `dislike_count` bigint(20)          not null default 0 comment '点踩计数',
    `update_time`   datetime            not null default current_timestamp on update current_timestamp comment '最后修改时间',
    `create_time`   datetime            not null default current_timestamp comment '创建时间',
    primary key (`id`)
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '计数表';

create table `t_like_status`
(
    `id`          bigint(20) unsigned not null auto_increment comment '主键ID',
    `video_id`    bigint(20) unsigned not null default 0 comment '视频id',
    `user_id`     bigint(20) unsigned not null default 0 comment '用户id',
    `state`       tinyint(4) unsigned          default 0 not null comment '点赞状态 1-喜欢；2-不喜欢',
    `update_time` datetime            not null default current_timestamp on update current_timestamp comment '最后修改时间',
    `create_time` datetime            not null default current_timestamp comment '创建时间',
    primary key (`id`)
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '行为表';