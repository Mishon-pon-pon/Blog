CREATE TABLE Articles(
  ArticleId    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  Title  TEXT,
  TextArticle TEXT
);
CREATE TABLE Users(
  UserID    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  Email  TEXT,
  Pass TEXT
);
