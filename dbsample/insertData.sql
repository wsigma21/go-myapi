# 記事データ２つ
insert into articles
(title, contents, username, nice, created_at)
values
('firstPost', 'This is my first blog', 'abe', 2, now());

insert into articles
(title, contents, username, nice)
values
('2nd', 'Second blog post', 'abe', 4);

# コメントデータ２つ
insert into comments
(article_id, message, created_at)
values
(1, '1st comment yeah', now());

insert into comments
(article_id, message)
values
(1, 'welcome');