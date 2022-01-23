window.addEventListener('DOMContentLoaded', () => {

  const closeButton = document.querySelector('#advertisigCloseButton');
  const advertisigElement = document.querySelector('.advertisig');
  const advertisigContent = document.querySelector('.advertisig__bottom');
  const advertisigLink = document.querySelector('.advertisig__link');


  let clickCounter = 0;

  closeButton.onclick = function() {
    
    clickCounter++;

    console.log( clickCounter );

    switch(clickCounter) {
      case 0:
        break;
      case 1:
        alert('Ненене, жопа же отвалится!!! Давай нажимай');
        break;
      case 2:
        alert('Зря пытаешься');
        break;
      case 3:
        alert('Да ты шо лох??? Заебал нажми на ссылку');
        break;
      case 4:
        alert('Ещё раз и эта реклама станет ещё больше');
        break;
      case 5:
        advertisigElement.classList.add('advertisig_big');
        advertisigContent.classList.add('advertisig__bottom_big');
        alert('Я предупреждал... Лучше перейди по ссылке!!! Иначе увидишь, что будет дальше!!!!!!')
        break;
      case 6:
        alert('Ну ладно ладно, что ж ты такой настойчивый');
        advertisigElement.classList.add('display-none');
        advertisigContent.classList.add('display-none');
        break;
    }
    
  }

  advertisigLink.onclick = function() {
    advertisigElement.classList.add('display-none');
    advertisigContent.classList.add('display-none');
  }

})