-- +migrate up
INSERT INTO `cases` (`case_id`, `problem_id`, `docs`, `version`)
VALUES('trump-in-the-news-i-1', 'trump-in-the-news-i', '["{\\"date\\": \\"2022-09-23\\", \\"short_description\\": \\"trump xxx\\", \\"@timestamp\\": \\"2022-09-23T00:00:00.000+08:00\\", \\"link\\": \\"https://www.huffpost.com/entry/covid-boosters-uptake-us_n_632d719ee4b087fae6feaac9\\", \\"category\\": \\"U.S. NEWS\\", \\"headline\\": \\"trump xxx\\", \\"authors\\": \\"Carla K. Johnson, AP\\"}"]', 0);

INSERT INTO `cases` (`case_id`, `problem_id`, `docs`, `version`)
VALUES('trump-in-the-news-ii-1', 'trump-in-the-news-ii', '["{\\"date\\": \\"2022-09-23\\", \\"short_description\\": \\"trump xxx\\", \\"@timestamp\\": \\"2022-09-23T00:00:00.000+08:00\\", \\"link\\": \\"https://www.huffpost.com/entry/covid-boosters-uptake-us_n_632d719ee4b087fae6feaac9\\", \\"category\\": \\"U.S. NEWS\\", \\"headline\\": \\"trump xxx\\", \\"authors\\": \\"Carla K. Johnson, AP\\"}"]', 0);

INSERT INTO `cases` (`case_id`, `problem_id`, `docs`, `version`)
VALUES('count-different-news-by-category-1', 'count-different-news-by-category', '["{\\"date\\": \\"2022-09-23\\", \\"short_description\\": \\"trump xxx\\", \\"@timestamp\\": \\"2022-09-23T00:00:00.000+08:00\\", \\"link\\": \\"https://www.huffpost.com/entry/covid-boosters-uptake-us_n_632d719ee4b087fae6feaac9\\", \\"category\\": \\"U.S. NEWS\\", \\"headline\\": \\"trump xxx\\", \\"authors\\": \\"Carla K. Johnson, AP\\"}"]', 0);

-- +migrate down
DELETE FROM `cases` WHERE `case_id` = 'trump-in-the-news-i-1';
DELETE FROM `cases` WHERE `case_id` = 'trump-in-the-news-ii-1';
DELETE FROM `cases` WHERE `case_id` = 'count-different-news-by-category-1';