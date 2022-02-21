CREATE TABLE ads
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER NOT NULL,
    name        VARCHAR(200) NOT NULL,
    date        TIMESTAMP DEFAULT NOW(),
    price       INTEGER NOT NULL,
    photo       TEXT NOT NULL,
    description VARCHAR(1000)
);

CREATE TABLE photos
(
    id      SERIAL PRIMARY KEY,
    ad_id   INTEGER NOT NULL,
    link    TEXT NOT NULL,
    FOREIGN KEY (ad_id) REFERENCES ads(id) ON DELETE CASCADE
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
    FOREIGN KEY (ad_id) REFERENCES ads(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

INSERT INTO ads(user_id, name, date, price, photo, description) VALUES
    (1, 'Объявление 1', '2021-10-11', 1000, 'https://picsum.photos/id/101/200/200', 'Тестовое объявление 1'),
    (1, 'Объявление 2', '2021-10-12', 2000, 'https://picsum.photos/id/201/200/200', 'Тестовое объявление 2'),
    (1, 'Объявление 3', '2021-10-13', 3000, 'https://picsum.photos/id/301/200/200', 'Тестовое объявление 3'),
    (1, 'Объявление 4', '2021-10-14', 4000, 'https://picsum.photos/id/401/200/200', 'Тестовое объявление 4'),
    (1, 'Объявление 5', '2021-10-15', 5000, 'https://picsum.photos/id/501/200/200', 'Тестовое объявление 5'),
    (1, 'Объявление 6', '2021-10-16', 6000, 'https://picsum.photos/id/601/200/200', 'Тестовое объявление 6'),
    (1, 'Объявление 7', '2021-10-17', 7000, 'https://picsum.photos/id/701/200/200', 'Тестовое объявление 7'),
    (1, 'Объявление 8', '2021-10-18', 8000, 'https://picsum.photos/id/801/200/200', 'Тестовое объявление 8');

INSERT INTO photos(ad_id, link) VALUES
    (1, 'https://picsum.photos/id/102/200/200'),
    (1, 'https://picsum.photos/id/103/200/200'),
    (2, 'https://picsum.photos/id/202/200/200'),
    (2, 'https://picsum.photos/id/203/200/200'),
    (3, 'https://picsum.photos/id/302/200/200'),
    (3, 'https://picsum.photos/id/303/200/200'),
    (4, 'https://picsum.photos/id/402/200/200'),
    (4, 'https://picsum.photos/id/403/200/200'),
    (5, 'https://picsum.photos/id/502/200/200'),
    (5, 'https://picsum.photos/id/503/200/200'),
    (6, 'https://picsum.photos/id/602/200/200'),
    (6, 'https://picsum.photos/id/603/200/200'),
    (7, 'https://picsum.photos/id/702/200/200'),
    (7, 'https://picsum.photos/id/703/200/200'),
    (8, 'https://picsum.photos/id/802/200/200'),
    (8, 'https://picsum.photos/id/803/200/200');

INSERT INTO tags(name) VALUES
    ('ТЕСТ'),
    ('КРАСНЫЙ'),
    ('ЗЕЛЕНЫЙ'),
    ('СИНИЙ'),
    ('ВАЖНОЕ'),
    ('СПАМ');

INSERT INTO ads_tags(ad_id, tag_id) VALUES
    (1, 1),
    (2, 1),
    (3, 1),
    (4, 1),
    (5, 1),
    (6, 1),
    (7, 1),
    (8, 1),
    (1, 2),
    (1, 4),
    (1, 5),
    (2, 3),
    (2, 6),
    (3, 2),
    (3, 3),
    (3, 4),
    (3, 5),
    (4, 2),
    (4, 6),
    (5, 6),
    (6, 4),
    (7, 3),
    (7, 5),
    (8, 3),
    (8, 4),
    (8, 5);