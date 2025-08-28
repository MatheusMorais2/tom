ALTER TABLE post_keywords RENAME TO article_keywords;
ALTER TABLE article_keywords RENAME COLUMN post_id TO article_id;
