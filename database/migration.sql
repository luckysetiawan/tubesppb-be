CREATE TABLE users (
    uid  SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    profile_picture TEXT,
    friend_mode BOOLEAN DEFAULT FALSE
);

CREATE TABLE masimelrowoo (
    id SERIAL PRIMARY KEY,
    masimelrowoo TEXT NOT NULL,
    user_from INT NOT NULL,
    user_to INT,
    is_read BOOLEAN DEFAULT FALSE,
    is_liked BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_from)
        REFERENCES users (uid),
    FOREIGN KEY (user_to)
        REFERENCES users (uid)
);

CREATE TABLE friend_list (
    friend_a INT NOT NULL,
    friend_b INT NOT NULL,
    FOREIGN KEY (friend_a)
        REFERENCES users (uid),
    FOREIGN KEY (friend_b)
        REFERENCES users (uid),
    PRIMARY KEY (friend_a, friend_b)
);