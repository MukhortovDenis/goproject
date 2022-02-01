const chests = document.querySelectorAll('.chest');
const chestsContent = document.querySelector('.chests__content');
const chestName = document.querySelector('.chest-content__name');
const stoneList = document.querySelector('.chest-content__list');
const mainWindow = document.querySelector('.main__window');
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

function getChestItems(url) {
  return fetch(url).then(response => {
    return response.json()
  });

  // let data = {
  //   "treasureName": "The Chest of the Legendary Quarry",
  //   "treasureContent": [
  //     {
  //       "stoneName": "Хуета",
  //       "stoneID": 1,
  //       "stoneURL": "static/images/common/chest.png",
  //       "stoneRare": "poop"
  //     }, 
  //     {
  //       "stoneName": "Залупня",
  //       "stoneID": 2,
  //       "stoneURL": "static/images/common/chest.png",
  //       "stoneRare": "poop"
  //     }
  //   ]
  // }
}

function showChestContent(items) {
  mainWindow.classList.remove('display-none');
  stoneList.classList.remove('display-none');
  closeButton.classList.remove('display-none');
  chestsContent.classList.remove('display-none');

  let stones;

  chestName.innerHTML = `${items.treasureName}`;

  for (let item in items) {
    if (typeof items[item] == 'object') {
      stones = items[item];

      for (let stone in stones) {
        stoneList.innerHTML += `
          <li class="chest-content__item item">
            <img class="item__img" src="${stones[stone].stoneURL}" alt="${stones[stone].stoneRare}">
            <div class="item__name ${stones[stone].stoneRare}">${stones[stone].stoneName}</div>
          </li>
        `;
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
    chestsContent.classList.add('display-none');

    clearChestContent();
  })
}

chests.forEach(chest => {
  chest.addEventListener('click', function() {
    getChestItems(`/chest?id=${getChestId(chest)}`)
      .then(data => showChestContent(data))
      .catch(err => console.log(err))
    closePopup();
  })
});