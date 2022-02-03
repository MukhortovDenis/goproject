window.addEventListener('DOMContentLoaded', function() {
  const stoneList = document.querySelector('.roulette__list');
  const startButton = document.querySelector('.chest__buy-btn');
  const prizeBox = document.querySelector('.prize__box');
  const prizeWindow = document.querySelector('.roulette__prize');
  const mainWindow = document.querySelector('.main__window');
  const acceptButton = document.querySelector('.accept__button');
  const closeButton = document.querySelector('.close-popup__button');

  function getRandom(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  function createRandomList() {
    let chances = [70, 24, 5, 1];
    let numberList = [];
    let randomNumber;

    for (let i = 0; i < 100; i++) {
      randomNumber = getRandom(0, 100);

      if (randomNumber >= 30 && randomNumber < 100) {
        numberList.push(chances[0]);
      } else if (randomNumber >= 6 && randomNumber < 30) {
          numberList.push(chances[1]);
      }  else if (randomNumber >= 1 && randomNumber < 6) {
          numberList.push(chances[2]);
      } else if (randomNumber < 1) {
          numberList.push(chances[3]);
      }
    }

    numberList[50] = 'Жопа муравья';

    return numberList;
  }

  function pasteElements(randomList) {
    for (let i = 0; i < randomList.length; i++) {

      switch(randomList[i]) {
        case 70:
          stoneList.innerHTML += `<li class="roulette__item uncommon"><img class="stone__img" src="static/images/common/hatat.png" alt=""></li>`;
          break;
        case 24:
          stoneList.innerHTML += `<li class="roulette__item rare"><img class="stone__img" src="static/images/common/amber.png" alt=""></li>`;
          break;
        case 5:
          stoneList.innerHTML += `<li class="roulette__item legendary"><img class="stone__img" src="static/images/common/ruby.png" alt=""></li>`;
          break;
        case 1:
          stoneList.innerHTML += `<li class="roulette__item arcana"><img class="stone__img" src="static/images/common/avatar.jpg" alt=""></li>`;
          break;
        case 'Жопа муравья':
          stoneList.innerHTML += `<li class="roulette__item arcana"><img class="stone__img" src="static/images/common/avatar.jpg" alt=""></li>`;
          break;
      }
    }
  }

  function getPx(width, margin, id) {
    return width * id + margin * id + (getRandom(21, 189));
  }

  function animate({timing, draw, duration}) {
    let start = performance.now();
  
    requestAnimationFrame(function animate(time) {
      let timeFraction = (time - start) / duration;

      if (timeFraction > 1) timeFraction = 1;
  
      let progress = timing(timeFraction);
  
      draw(progress);
  
      if (timeFraction < 1) {
        requestAnimationFrame(animate);
      }
    });
  }

  function clearStoneList() {
    stoneList.style.right = 0;
    stoneList.innerHTML = '';
  }

  function showPrize(randomList) {
    prizeWindow.innerHTML = `
      <img class="prize__img" src="static/images/common/avatar.jpg" alt="">
      <div class="prize__name arcana-text">${randomList[50]}</div>
    `;

    prizeBox.classList.remove('display-none');
    mainWindow.classList.remove('display-none');
  }

  startButton.addEventListener('click', function() {
    let px = getPx(150, 8, 50);
    
    pasteElements( createRandomList() );

    animate({
      duration: 8000,
      timing: function easeOut(timeFraction) {
        return 1 - Math.pow(1 - timeFraction, 3)
      },
      draw: function(progress) {
        stoneList.style.right = progress * px + 'px';
      }
    });

    setTimeout(function() {
      showPrize( createRandomList() )

      animate({
        duration: 300,
        timing: function easyOut(timeFraction) {
          return 1 - Math.pow(1 - timeFraction, 3)
        },
        draw: function(progress) {
          prizeWindow.style.boxShadow = `0px 0px ${progress * 90}px 10px #ade55c`;
        }
      });
    }, 8200)

    startButton.setAttribute('disabled', 'disabled');
    closeButton.classList.add('display-none');
  })

  acceptButton.addEventListener('click', function() {
    prizeBox.classList.add('display-none');
    
    clearStoneList();

    startButton.removeAttribute('disabled', 'disabled');
    closeButton.classList.remove('display-none');
  })
});