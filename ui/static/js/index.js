window.addEventListener('DOMContentLoaded', function() {
  document.querySelector('.arrow').addEventListener('click', function() {
    document.querySelector('.arrow__icon').classList.toggle('rotate');
    document.querySelector('.user__menu').classList.toggle('user__menu_active');
  });
});