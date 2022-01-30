const chests = document.querySelectorAll('.chest');
const mainWindow = document.querySelector('.main__window');
const stoneList = document.querySelector('.chest-content__list');
const closeButton = document.querySelector('.close-popup__button');

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

function getChestItems(url, chestId) {
  // const response = await fetch(url, {
  //   method: 'GET',
  //   body: chestId,
  //   headers: {
  //     'Content-Type': 'application/json;charset=utf-8'
  //   },
  // });

  // data = await response.json();
  
  let data = {
    "treasureName": "The Chest of the Legendary Quarry",
    "treasureId": "1",
    "treasureContent": {
      "1": "Rubby",
      "2": "Emerald",
      "3": "Diamond",
      "4": "Amber",
      "5": "Hatat",
      "6": "Lapis Lazuli"
    }
  }

  return data;
  /* {
        "treasureName": "The Chest of the Legendary Quarry",
        "treasureId": "1",
        "treasureContent": {
          "1": "Rubby",
          "2": "Emerald",
          "3": "Diamond",
          "4": "Amber",
          "5": "Hatat",
          "6": "Lapis Lazuli"
        }
      } */
}

function showChestContent() {
  mainWindow.classList.remove('display-none');
  stoneList.classList.remove('display-none');
  closeButton.classList.remove('display-none');

  let stones;
  let items;

  items = getChestItems();

  for (let item in items) {
    console.log( items[item] );

    if (typeof items[item] == 'object') {
      stones = items[item];

      for (let stone in stones) {
        stoneList.innerHTML += `<li class="chest-content__item">${stones[stone]}</li>`;

        console.log( stones[stone] );
      }
    }
  }
} 

function clearChestContent() {
  if (stoneList.innerHTML != '') {
    stoneList.innerHTML = '';
  }
}

function closePopup() {
  closeButton.addEventListener('click', function() {
    mainWindow.classList.add('display-none');
    stoneList.classList.add('display-none');
    closeButton.classList.add('display-none');

    clearChestContent();
  })
}

chests.forEach(chest => {
  chest.addEventListener('click', function() {
    // getChestItems( '/loh', getChestId(chest) );
    console.log( getChestId(chest) );

    showChestContent();
    // console.log( getChestId(chest) );

    closePopup();
  })
})