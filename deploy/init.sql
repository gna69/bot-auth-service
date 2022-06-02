CREATE TABLE IF NOT EXISTS users
(
    id        INTEGER PRIMARY KEY,
    firstName VARCHAR(64),
    lastName  VARCHAR(64),
    userName  VARCHAR(64),
    langCode  VARCHAR(3),
    isBot     BOOLEAN,
    chatId    INTEGER
);

CREATE TABLE IF NOT EXISTS groups
(
    id      SERIAL primary key,
    ownerId INTEGER REFERENCES users (id) NOT NULL,
    "name"    VARCHAR(64),
    members INTEGER ARRAY,
    unique (id, "name")
);