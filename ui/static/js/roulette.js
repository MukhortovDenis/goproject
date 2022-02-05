window.addEventListener('DOMContentLoaded', function() {
  const startButton = document.querySelector('.chest__buy-btn');
  const stoneList = document.querySelector('.roulette__list');
  const prizeBox = document.querySelector('.prize__box');
  const prizeWindow = document.querySelector('.roulette__prize');
  const acceptButton = document.querySelector('.accept__button');
  const closeButton = document.querySelector('.close-popup__button');

  const itemID = 50;
  const itemWidth = 160;
  const itemMargin = 10;

  function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  function getRandomNumber(min, max) {
    return Math.random() * (max - min + 1) + min;
  }

  function getChestContent() {
    let data;

    let stones = {
      "stones": [
        {
          "stoneName": "hatat",
          "stoneRarity": "common",
          "stoneWinChance": 58,
          "stoneImgURL": "../static/images/common/hatat.png"
        },
        {
          "stoneName": "amber",
          "stoneRarity": "uncommon",
          "stoneWinChance": 30,
          "stoneImgURL": "../static/images/common/amber.png"
        },
        {
          "stoneName": "heliotrope",
          "stoneRarity": "rare",
          "stoneWinChance": 7,
          "stoneImgURL": "../static/images/common/heliotrope.png"
        },
        {
          "stoneName": "lapis-lazuli",
          "stoneRarity": "mythical",
          "stoneWinChance": 3,
          "stoneImgURL": "../static/images/common/lapis-lazuli.png"
        },
        {
          "stoneName": "eremeevit",
          "stoneRarity": "legendary",
          "stoneWinChance": 1.44,
          "stoneImgURL": "../static/images/common/eremeevit.png"
        },
        {
          "stoneName": "aquamarin",
          "stoneRarity": "immortal",
          "stoneWinChance": 0.5,
          "stoneImgURL": "../static/images/common/aquamarin.png"
        },
        {
          "stoneName": "avatar",
          "stoneRarity": "arcana",
          "stoneWinChance": 0.05,
          "stoneImgURL": "../static/images/common/avatar.jpg"
        },
        {
          "stoneName": "emerald",
          "stoneRarity": "ancient",
          "stoneWinChance": 0.01,
          "stoneImgURL": "../static/images/common/emerald.png"
        },
      ]
    };

    for (let stone in stones) {
      data = stones[stone];
    }

    return data;
  }

  function createRandomList(chestContent) {
    let chances = [100];
    let numberList = [];
    let randomNumber;

    for (let i = 0; i < chestContent.length; i++) {
      chances.push([chestContent[i].stoneWinChance, 
                    chestContent[i].stoneName,
                    chestContent[i].stoneRarity,
                    chestContent[i].stoneImgURL
      ]);
    }
  
    for (let i = 0; numberList.length != 100; i++) {
      let res = 100;
      
      randomNumber = getRandomNumber(0, 100);

      outer:
      for (let j = 1; j < chances.length; j++) {
        if (randomNumber >= (res - chances[j][0]) && randomNumber < res) {
          numberList.push(chances[j]);
          break;
        } else {
            res -= chances[j][0];
            continue outer;
        }
      }

    }  

    console.log( numberList );

    return numberList;
  }

  function pasteElements(list) {
    for (let i = 0; i < list.length; i++) { 
      stoneList.innerHTML += `<li class="roulette__item ${list[i][2]}"><img class="stone__img" src="${list[i][3]}" alt=""></li>` 
    }
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
    let stoneName = list[id][1];
    let formatWinStoneName = stoneName.charAt(0).toUpperCase() + stoneName.slice(1);

    prizeWindow.innerHTML = `
      <img class="prize__img" src="${list[id][3]}" alt="">
      <div class="prize__name ${list[id][2]}">${formatWinStoneName}</div>
    `;

    prizeBox.classList.remove('display-none');

    animate({
      duration: 400,
      timing: function easyOut(timeFraction) {
        return 1 - Math.pow(1 - timeFraction, 3)
      },
      draw: function(progress) {
        prizeWindow.style.boxShadow = `0px 0px ${progress * 90}px 10px var(--${list[id][2]})`;
      }
    });
  }

  startButton.addEventListener('click', function() {
    let rotate = rotateTo(itemWidth, itemMargin, itemID);
    let chestContent;
    let createdList;

    chestContent = getChestContent();
    createdList = createRandomList( chestContent );

    pasteElements(createdList);

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