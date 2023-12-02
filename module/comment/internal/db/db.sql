create table t_article
(
    id          int primary key auto_increment comment '主键',
    title       varchar(255) not null comment '标题',
    content     text         not null comment '正文',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '文章表';

create table t_comment
(
    id            int primary key auto_increment comment '主键',
    article_id    int comment '文章id',
    user_id       int comment '用户id',
    content       varchar(1024) not null comment '正文(限制340字符)',
    like_count    int                    default 0 comment 'like count',
    dislike_count int                    default 0 comment 'dislike count',
    update_time   timestamp     not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time   timestamp     not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '评论表';

create table t_comment_like
(
    id          int primary key auto_increment comment '主键',
    user_id     int comment '用户id',
    like_status int                default 0 comment '1-like, 2-dislike',
    update_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '评论点赞表';


create table t_reply
(
    id          int primary key auto_increment comment '主键',
    comment_id  int comment '评论id',
    user_id     int comment '评论人id',
    content     varchar(256) not null comment '回复内容(限制80字符)',
    reply_time  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '回复表';

