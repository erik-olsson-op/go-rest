DROP TABLE IF EXISTS user;
CREATE TABLE user
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    email      TEXT NOT NULL UNIQUE,
    password   TEXT NOT NULL,
    registered DATETIME
);

DROP TABLE IF EXISTS event;
CREATE TABLE event
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    location    VARCHAR(255) NOT NULL,
    date_time   DATETIME     NOT NULL,
    user_id     INTEGER,
    FOREIGN KEY (user_id) REFERENCES user (id)
);