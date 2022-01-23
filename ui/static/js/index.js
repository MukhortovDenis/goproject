window.addEventListener('DOMContentLoaded', function() {

  document.querySelector('.arrow').addEventListener('click', function() {
    document.querySelector('.arrow__icon').classList.toggle('rotate');
    document.querySelector('.user__menu').classList.toggle('user__menu_active');
  });

  const scrollTo = element => {
    window.scroll({
      left: 0,
      top: element.offsetTop,
      behavior: 'smooth'  
    });
  };

  const link = document.querySelector('.scroll__btn');
  const cards = document.querySelector('#cards');

  link.addEventListener('click', () => {
    scrollTo(cards);
  });

});


