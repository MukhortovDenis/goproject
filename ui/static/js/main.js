document.addEventListener('DOMContentLoaded', function() {
  document.querySelector('#burger').addEventListener('click', function() {
    document.querySelector('.user__icon').classList.toggle('rotate')
    document.querySelector('#user__menu').classList.toggle('user__menu_active')
  })
});