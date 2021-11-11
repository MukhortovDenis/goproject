window.addEventListener('DOMContentLoaded', () => {

  function getSong(song) {
    songName.innerHTML = song;
    audio.src = `static/audio/${song}.mp3`;
  }

  function playSong() {
    audio.play();
  }

  function pauseSong() {
    audio.pause();
  }

  function nextSong() {
    songIndex++;

    if (songIndex > songs.length - 1) {
      songIndex = 0;
    }

    getSong(songs[songIndex]);

    playSong();

    playBtn.classList.add('control__hidden');
  }

  function prevSong() {
    songIndex--;

    if (songIndex < 0 ) {
      songIndex = songs.length - 1;
    }

    getSong(songs[songIndex]);

    playSong();

    playBtn.classList.add('control__hidden');
  }

  const playBtn = document.querySelector('.player__control_play');
  const pauseBtn = document.querySelector('.player__control_pause');
  const prevBtn = document.querySelector('.player__control_prev');
  const nextBtn = document.querySelector('.player__control_next');
  const audio = document.querySelector('.player__audio');
  const songName = document.querySelector('.player__song-name');

  const songs = [`fem.love_Fotografiruyu_zakat`, `BUNIN_Sosedi`, 
                 `MORGENSHTERN_RATATATA`, `MetaGame_International`, 
                 `Skrillex_Bangarang`, `MORGENSHTERN_DULO`, 
                 `Baauer_Harlem_Shake`, `Skrillex_First_Of_The_Year_Equinox`, 
                 `AZAZLO_Igrok_s_nikom`, `Skrillex_Kill_EVERYBODY`, 
                 `shadowraze_shadowfiend`, `MetaGame_Luchshijj_farmer_v_dote`]

  let songIndex = 0;


  getSong(songs[songIndex]);

  playBtn.addEventListener('click', () => {
    playBtn.classList.add('control__hidden');
    pauseBtn.classList.remove('control__hidden');
    playSong();
  });

  pauseBtn.addEventListener('click', () => {
    pauseBtn.classList.add('control__hidden');
    playBtn.classList.remove('control__hidden')
    pauseSong();
  });

  nextBtn.addEventListener('click', nextSong);
  
  prevBtn.addEventListener('click', prevSong);

  audio.addEventListener('ended', nextSong);

});