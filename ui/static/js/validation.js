window.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('form-register');
  const userName = document.getElementById('name');
  const userEmail = document.getElementById('email');
  const userPassword = document.getElementById('password');
  const userConfPassword = document.getElementById('confirn-password');

  form.addEventListener('submit', (e) => {
    e.preventDefault();

    checkInputs();
  });

  function checkInputs() {
    const userNameValue = userName.value.trim();
    const userEmailValue = userEmail.value.trim();
    const userPasswordValue = userPassword.value.trim();
    const userConfPasswordValue = userConfPassword.value.trim();

    if (userNameValue === '') {
      setErrorFor(userName, 'Это поле не должно быть пустым');
    } else if (!isName(userNameValue)) {
      setErrorFor(userName, 'Логин введён некорректно');
    }
    else {
      setSuccessFor(userName)
    }

    if (userEmailValue === '') {
      setErrorFor(userEmail, 'Это поле не должно быть пустым');
    } else if (!isEmail(userEmailValue)) {
      setErrorFor(userEmail, 'Email введён некорректно');
    } else {
      setSuccessFor(userEmail);
    }

    if (userPasswordValue === '') {
      setErrorFor(userPassword, 'Это поле не должно быть пустым');
    } else if (!isPassword(userPasswordValue)) {
      setErrorFor(userPassword, 'Пароль введён некоректно');
    } else if (userPasswordValue !== userConfPasswordValue) {
      setErrorFor(userPassword, 'Пароли не совпадают');
    } else {
      setSuccessFor(userPassword);
    }

    if (userConfPasswordValue === '') {
      setErrorFor(userConfPassword, 'Это поле не должно быть пустым');
    } else if (!isPassword(userConfPasswordValue)) {
      setErrorFor(userConfPassword, 'Пароль введён некоректно');
    } else if (userConfPasswordValue !== userPasswordValue) {
      setErrorFor(userConfPassword, 'Пароли не совпадают');
    } else {
      setSuccessFor(userConfPassword);
    }
  }

  function setErrorFor (input, message) {
    const inputBox = input.parentElement;
    const errorMsg = inputBox.querySelector('.error__msg');  

    errorMsg.innerText = message;

    inputBox.className = 'form-input__box error';
  }

  function setSuccessFor(input) {
    const inputBox = input.parentElement;

    inputBox.className = 'form-input__box success';
  }

  function isName(name) {
    return /^[a-z0-9]{3,16}$/.test(name);
  }

  function isEmail(email) {
    return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,4}))$/.test(email);
  }

  function isPassword(password) {
    return /^(?=.*[a-zA-Z0-9]).{6,}$/.test(password);
  }
})