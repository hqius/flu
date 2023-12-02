# 文章评论表
# 使用文章ID作为键，评论的时间戳作为分值，评论ID作为成员存储在Sorted Set中
zadd article_comment:345 1624422000 1

# 评论回复表
# 使用文章ID作为键，评论的时间戳作为分值，评论ID作为成员存储在Sorted Set中
zadd article_comment_reply:1 1624422000 22

# 评论内容表
# 使用文章ID作为键，评论的时间戳作为分值，评论ID作为成员存储在Sorted Set中
hset article_comment_content:345 1 '{"content": "Good post!"}'
