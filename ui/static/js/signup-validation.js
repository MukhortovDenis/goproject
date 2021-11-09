window.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('form-register');
  const userName = document.getElementById('name');
  const userEmail = document.getElementById('email');
  const userPassword = document.getElementById('password');
  const userConfPassword = document.getElementById('confirn-password');
  const requestURL = '/save_user';

  form.onsubmit = async (e) => {
    e.preventDefault();

    const successCount = checkInputs();

    const formData = new FormData(form);

    formData.delete('password-check');

    const plainFormData = Object.fromEntries(formData.entries());
    const fromDataJSON = JSON.stringify(plainFormData);

    if (successCount === 4) {
      let response = fetch(requestURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: fromDataJSON
        });
       
      // let result = response.json();
  
      form.reset();
  
      // window.location.href = '/';

      setTimeout(() => window.location.href = '/', 1000);
    }
  }

  function checkInputs() {
    const userNameValue = userName.value.trim();
    const userEmailValue = userEmail.value.trim();
    const userPasswordValue = userPassword.value.trim();
    const userConfPasswordValue = userConfPassword.value.trim();

    let success = 0;

    if (userNameValue === '') {
      setErrorFor(userName, 'Это поле не должно быть пустым');
    } else if (userNameValue.length < 5 || userNameValue.length > 20) {
        setErrorFor(userName, 'Длина логина должна быть от 5 до 20 символов');
    } else if (!isName(userNameValue)) {
        setErrorFor(userName, 'Введены некорректные символы');
    } else {
        setSuccessFor(userName)
        success++;
    }

    if (userEmailValue === '') {
      setErrorFor(userEmail, 'Это поле не должно быть пустым');
    } else if (userEmailValue.length < 8) {
        setErrorFor(userEmail, 'Email слишком короткий');
    } else if (userEmailValue.length > 32) {
        setErrorFor(userEmail, 'Email слишком длинный');
    } else if (!isEmail(userEmailValue)) {
        setErrorFor(userEmail, 'Email введён некоректно');
    } else {
        setSuccessFor(userEmail);
        success++;
    }

    if (userPasswordValue === '') {
      setErrorFor(userPassword, 'Это поле не должно быть пустым');
    } else if (userPasswordValue.length < 6 || userPasswordValue.length > 32) {
        setErrorFor(userPassword, 'Длина пароля должна быть от 6 до 32 символов');
    } else if (!isPassword(userPasswordValue)) {
        setErrorFor(userPassword, 'Пароль введён некоректно');
    } else if (userPasswordValue !== userConfPasswordValue) {
        setErrorFor(userPassword, 'Пароли не совпадают');
    } else {
        setSuccessFor(userPassword);
        success++;
    }

    if (userConfPasswordValue === '') {
      setErrorFor(userConfPassword, 'Это поле не должно быть пустым');
    } else if (userConfPasswordValue.length < 6 || userConfPasswordValue.length > 32) {
        setErrorFor(userConfPassword, 'Длина пароля должна быть от 6 до 32 символов');
    } else if (!isPassword(userConfPasswordValue)) {
        setErrorFor(userConfPassword, 'Пароль введён некоректно');
    } else if (userConfPasswordValue !== userPasswordValue) {
        setErrorFor(userConfPassword, 'Пароли не совпадают');
    } else {
        setSuccessFor(userConfPassword);
        success++;
    }

    return success;
  }

  function setErrorFor(input, message) {
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
    return /^[a-zA-Z0-9]{5,20}$/.test(name);
  }

  function isEmail(email) {
    return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,4}))$/.test(email);
  }

  function isPassword(password) {
    return /^(?=.*[a-zA-Z0-9]).{6,32}$/.test(password);
  }
})