window.addEventListener('DOMContentLoaded', function() {

  function getRandom(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min)) + min;
  }

  const roulette = document.querySelector('.roulette');
  const rouletteBox = document.querySelector('.roulette__box')
  const buyButton = document.querySelector('.chest__buy-btn');
  const rouletteItem = document.querySelector('.roulette__item');
  const wonText = document.querySelector('.roulette__won-text');

  function anim(num) {
  
    if (num < numberList.length) {
      if (numberList[num] == 0) {
        rouletteItem.innerHTML = 'Рубин';
      } else if (numberList[num] == 1) {
        rouletteItem.innerHTML = 'Изумруд';
      } else if (numberList[num] == 2) {
        rouletteItem.innerHTML = 'Алмаз';
      } else if (numberList[num] == 3) {
        rouletteItem.innerHTML = '<img class="avatar__img" src="/static/images/common/avatar.jpg" alt="">';
      } else if (numberList[num] == 4) {
        rouletteItem.innerHTML = 'Лазурит';
      } else if (numberList[num] == 95) {
        rouletteItem.innerHTML = 'Говно';
        wonText.classList.remove('display-none');
      }
      // rouletteItem.innerHTML = numberList[num];
      setTimeout(anim, 65, ++num);
    }
  }
 
  let numberList = [];

  buyButton.addEventListener('click', () => {
    roulette.classList.remove('display-none');
    rouletteBox.classList.remove('display-none');

    for (let i = 0; i < 100; i++) {

      numberList.push( getRandom(0, 5) );
  
    }
  
    console.log( numberList );
  
    numberList.pop();
    numberList.push(95);
  
    console.log( numberList );
  
    setTimeout(anim, 65, 0);
    
  })

});