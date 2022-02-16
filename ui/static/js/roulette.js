const chests = document.querySelectorAll('.chest');
const chestsContent = document.querySelector('.chests__content');
const chestName = document.querySelector('.chest-content__name');
const chestContentList = document.querySelector('.chest-content__list');
const chestOverlay = document.querySelector('.chest-content__overlay');
const closeButton = document.querySelector('.close-popup__button');

const startButton = document.querySelector('.chest__buy-btn');
const stoneList = document.querySelector('.roulette__list');
const prize = document.querySelector('.prize');
const prizeBox = document.querySelector('.prize__box');
const prizeWindow = document.querySelector('.roulette__prize');
const acceptButton = document.querySelector('.accept__button');

const modal = document.querySelector('.modal');
const modalMessage = document.querySelector('.modal__message');
const confirmModalButton = document.querySelector('.modal__button');
const closeModalButton = document.querySelector('.close-modal__button')

const itemID = 50;
const itemWidth = 160;
const itemMargin = 10;

let chestID;
let chestContent;

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomNumber(min, max) {
  return Math.random() * (max - min + 1) + min;
}

function getChestId(chest) {
  let chestId;
  let chestInfo;
    
  chestInfo = chest.textContent.replace(/[\n\r]+|[\s]{2,}/g, ' ').trim();

  for (let i = 0; i < chestInfo.length; i++) {
    if (chestInfo[i] == +chestInfo[i]) {
      chestId = +chestInfo[i];
    }
  }

  return chestId;
}

function getChestItems(url) {
  return fetch(url).
    then(response => {
      return response.json()
  });
}

function showChestContent(items) {
  chestOverlay.classList.remove('display-none');
  chestContentList.classList.remove('display-none');
  closeButton.classList.remove('display-none');
  chestsContent.classList.remove('display-none');

  let stones;

  // <h2 class="chest-content__name poop"></h2>
  chestName.classList.add(`${items.chestRare}`);
  chestName.innerHTML = `${items.chestName}`;

  for (let item in items) {
    if (typeof items[item] == 'object') {
      stones = items[item];

      for (let stone in stones) {
        chestContentList.innerHTML += `
          <li class="chest-content__item item">
            <div class="item__wrapper">
              <img class="item__img" src="${stones[stone].stoneURL}" alt="${stones[stone].stoneName}">
              <div class="item__name ${stones[stone].stoneRare}">${stones[stone].stoneName}</div>
            </div>
          </li>
        `;
      }
    }
  }


  let chestPrice = items.chestPrice.toLocaleString();
  startButton.innerHTML = `Купить за ${chestPrice.replace(/\s/g, '.')} C`
} 

function showUserError(errorMessage) {
  modal.classList.remove('display-none');
  modalMessage.innerHTML = `${errorMessage}`;
}

function clearChestContent() {
  if (chestContentList.innerHTML != '') {
    chestContentList.innerHTML = '';
  }
}

function closePopup() {
  closeButton.addEventListener('click', function() {
    chestOverlay.classList.add('display-none');
    chestContentList.classList.add('display-none');
    closeButton.classList.add('display-none');
    chestsContent.classList.add('display-none');
    chestName.className = 'chest-content__name';

    clearChestContent();
  })
}

function closeModal() {
  closeModalButton.addEventListener('click', function() {
    modal.classList.add('display-none');
  })
}

chests.forEach(chest => {
  chest.addEventListener('click', function() {
    chestID = getChestId(chest);
    getChestItems(`/chest?id=${chestID}`)
      .then(function(data) {
        showChestContent(data);

        chestContent = data.chestContent; 
      })     
      .catch(err => console.log(err))

    closePopup();
    closeModal();
  })
});

function createRandomList(chestContent) {
  let chances = [100];
  let numberList = [];
  let randomNumber;

  for (let i = 0; i < chestContent.length; i++) {
    chances.push([
      chestContent[i].stoneName,
      chestContent[i].stoneChance,
      chestContent[i].stoneURL,
      chestContent[i].stoneRare,
    ]);
  }

  for (let i = 0; numberList.length != 100; i++) {
    let res = 100;
    
    randomNumber = getRandomNumber(0, 100);

    outer:
    for (let j = 1; j < chances.length; j++) {
      if (randomNumber >= (res - chances[j][1]) && randomNumber < res) {
        numberList.push(chances[j]);
        break;
      } else {
          res -= chances[j][1];
          continue outer;
      }
    }
  }  

  return numberList;
}

function pasteElements(list) {
  for (let i = 0; i < list.length; i++) { 
    stoneList.innerHTML += `<li class="roulette__item ${list[i][3]}"><img class="stone__img" src="${list[i][2]}" alt="${list[i][0]}"></li>` 
  }
}

function getPrize(url) {
  return fetch(url).
    then(response => {
      return response.json()
  });
}

function rotateTo(width, margin, id) {
  return (width * id + margin * id) - (width * 3 + margin * 3) - (width / 2) - (margin) + getRandomInt(1, width - 1);
}

function animate( {timing, draw, duration} ) {
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

function showPrize(list, id) {
  let stoneName = list[id][0];
  let formatWinStoneName = stoneName.charAt(0).toUpperCase() + stoneName.slice(1);

  prizeWindow.innerHTML = `
    <div class="prize__top">
      <img class="prize__img" src="${list[id][2]}" alt="${formatWinStoneName}">
    </div>
    <div class="prize__name ${list[id][3]}">${formatWinStoneName}</div>
    <div class="prize__description">${list[id][4]}</div>
  `;

  prize.classList.remove('display-none');

  animate({
    duration: 400,
    timing: function easyOut(timeFraction) {
      return 1 - Math.pow(1 - timeFraction, 3)
    },
    draw: function(progress) {
      prizeWindow.style.boxShadow = `0px 0px ${progress * 130}px 30px var(--${list[id][3]})`;
    }
  });
}

let audio = {};

function getAudio() {
  let number = getRandomInt(1, 5);
  let audioSrc = `static/audio/roulette${number}.mp3`;
  audio = new Audio(audioSrc);

  return audioSrc;
}

function startAudio() {
  audio.play();
}

function stopAudio() {
  audio.pause();
}


startButton.addEventListener('click', function() {
  let rotate = rotateTo(itemWidth, itemMargin, itemID);
  let createdList;

  createdList = createRandomList( chestContent );

  getPrize(`/open-chest?id=${chestID}`)
    .then(data => { 
        if (data.msg != 'Не авторизированный пользователь') {
          startButton.setAttribute('disabled', 'disabled');
          closeButton.classList.add('display-none');

          createdList[50] = [data.stoneName, '', data.stoneURL, data.stoneRare, data.stoneDescription ]

          pasteElements(createdList);

          startAudio( getAudio() );

          animate({
            duration: 8000,
            timing: function easeOut(timeFraction) {
              return 1 - Math.pow(1 - timeFraction, 3)
            },
            draw: function(progress) {
              stoneList.style.right = progress * rotate + 'px';
            }
          });
        
          setTimeout(function() {
            showPrize(createdList, itemID)
          }, 8200)
        } else {
          showUserError(data.msg);
        }
      }
    )
    .catch(err => console.log(err))

  // console.log( createdList );
})

acceptButton.addEventListener('click', function() {
  prize.classList.add('display-none');
  
  clearStoneList();

  stopAudio();

  startButton.removeAttribute('disabled', 'disabled');
  closeButton.classList.remove('display-none');
})

confirmModalButton.addEventListener('click', function() {
  modal.classList.add('display-none');
  modalMessage.innerHTML = '';
})