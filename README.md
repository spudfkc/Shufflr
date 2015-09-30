# Shufflr
Better Spotify Radio

Spotify's radio sucks. Their Discovery Weekly playlists are pretty cool though.


In it's simplest form, there will be a continuely-modified Spotify playlist.

1. Given some seeds (artists/songs/genres) a playlist [of an arbitrary size] is generated.
2. When a song is done playing, it is removed from the playlist.
3. A new song (or batch of songs) is added to the playlist as-needed.

The idea is to also have a web interface for displaying the playlist and acting on it, such as adjusting the seeds, manually adding/removing songs, etc.

The problem with this is the way Spotify works...and either the Spotify Desktop app or the Spotify Web client would also need to be open, which prevents this solution from being an isolated application.


TODO: figure out go dependencies...
