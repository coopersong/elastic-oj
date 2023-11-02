-- +migrate up
INSERT INTO `problems` (`problem_id`, `title`, `description`, `es_index`, `standard_query`, `version`)
VALUES('trump-in-the-news-i', 'Trump in the News I', 'Please write an Elasticsearch query to get all news whose headline contains "trump"', 'news_headlines', '{"query": {"match": {"headline": "trump"}}}', 0);

INSERT INTO `problems` (`problem_id`, `title`, `description`, `es_index`, `standard_query`, `version`)
VALUES('trump-in-the-news-ii', 'Trump in the News II', 'Please write an Elasticsearch query to get all news whose any of fields contains "trump"', 'news_headlines', '{"query": {"multi_match": {"query": "trump", "fields": ["headline", "short_description", "category", "authors"]}}}', 0);

INSERT INTO `problems` (`problem_id`, `title`, `description`, `es_index`, `standard_query`, `version`)
VALUES('count-different-news-by-category', 'Count Different News by Category', 'Please write an Elasticsearch query to list all categories and the count of the category', 'news_headlines', '{"query": {"match": {"headline": "trump"}}}', 0);

-- +migrate down
DELETE FROM `problems` WHERE `problem_id` = 'trump-in-the-news-i';
DELETE FROM `problems` WHERE `problem_id` = 'trump-in-the-news-ii';
DELETE FROM `problems` WHERE `problem_id` = 'count-different-news-by-category';