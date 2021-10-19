CREATE TABLE ads
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    date        TIMESTAMP NOT NULL,
    price       INTEGER NOT NULL,
    description VARCHAR(1000)
);

CREATE TABLE photos
(
    id      SERIAL PRIMARY KEY,
    link    TEXT NOT NULL
);

CREATE TABLE ads_photos
(
    ad_id       INTEGER NOT NULL,
    photo_id    INTEGER NOT NULL,
    is_main     BOOLEAN NOT NULL,
    PRIMARY KEY (ad_id, photo_id),
    FOREIGN KEY (ad_id) REFERENCES ads(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);

CREATE TABLE tags
(
    id      SERIAL PRIMARY KEY,
    name    TEXT NOT NULL
);

CREATE TABLE ads_tags
(
    ad_id       INTEGER NOT NULL,
    tag_id      INTEGER NOT NULL,
    PRIMARY KEY (ad_id, tag_id),
    FOREIGN KEY (ad_id) REFERENCES ads(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

