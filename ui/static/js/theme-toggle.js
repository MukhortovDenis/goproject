window.addEventListener('DOMContentLoaded', function() {
  const body = document.querySelector('body');
  const option1 = document.querySelector('#lightThemeOption');
  const option2 = document.querySelector('#darkThemeOption');

  if (localStorage.getItem('theme') === 'light') {
    body.classList.remove('page__body_theme_dark');
    body.classList.add('page__body_theme_light');
    option2.removeAttribute('checked');
    option1.setAttribute('checked', 'checked');
  } else if (localStorage.getItem('theme') === 'dark') {
    body.classList.remove('page__body_theme_light');
    body.classList.add('page__body_theme_dark');
    option1.removeAttribute('checked');
    option2.setAttribute('checked', 'checked');
  }

  option1.addEventListener('click', function() {
    if (body.classList.contains('page__body_theme_dark')) {
      body.classList.remove('page__body_theme_dark');
    }
    body.classList.add('page__body_theme_light');
    localStorage.removeItem('theme', 'dark');
    localStorage.setItem('theme', 'light');
  });

  option2.addEventListener('click', function() {
    if (body.classList.contains('page__body_theme_light')) {
      body.classList.remove('page__body_theme_light');
    }
    body.classList.add('page__body_theme_dark');
    localStorage.removeItem('theme', 'light');
    localStorage.setItem('theme', 'dark');
  });

  console.log( 'hi' );
});