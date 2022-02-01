window.addEventListener('DOMContentLoaded', () => {
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
})