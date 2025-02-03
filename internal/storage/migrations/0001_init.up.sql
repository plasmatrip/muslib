BEGIN;

CREATE TABLE IF NOT EXISTS music_library (
	id serial NOT NULL UNIQUE,
    group_name varchar(255) NOT NULL,
    song_name varchar(255) NOT NULL,
    release_date timestamp NOT NULL,
    lyrics text,
    link varchar(255),
	PRIMARY KEY (id),
    UNIQUE (group_name, song_name)
);

CREATE INDEX IF NOT EXISTS idx_music_library_id ON music_library (id);
CREATE INDEX IF NOT EXISTS idx_music_library_group_name ON music_library (group_name);
CREATE INDEX IF NOT EXISTS idx_music_library_song_name ON music_library (song_name);

COMMIT;