package queries

const (
	AddSong = `
		INSERT INTO music_library (group_name, song_name, release_date, lyrics, link)
		VALUES (@group_name, @song_name, @release_date, @lyrics, @link);
	`
	DeleteSong = `
		DELETE FROM music_library
		WHERE group_name = @group_name AND song_name = @song_name;
	`

	UpdateSong = `
		UPDATE music_library
		SET group_name = @group_name,
			song_name = @song_name,
			release_date = COALESCE(@release_date, release_date),
			lyrics = CASE WHEN TRIM(@lyrics) != '' THEN @lyrics ELSE lyrics END,
			link = CASE WHEN TRIM(@link) != '' THEN @link ELSE link END
		WHERE group_name = @group_name AND song_name = @song_name;
	`
)
