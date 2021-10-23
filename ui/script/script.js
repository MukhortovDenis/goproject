document.addEventListener('DOMContentLoaded', function() {
  document.querySelector('#burger').addEventListener('click', function() {
    document.querySelector('#burger').classList.toggle('burger-open')
    document.querySelector('#user__menu').classList.toggle('user__menu_active')
  })
})