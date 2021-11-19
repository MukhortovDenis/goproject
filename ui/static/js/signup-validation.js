window.addEventListener('DOMContentLoaded', () => {

  "use strict";

  const form = document.getElementById('form-register');
  const userName = document.getElementById('name');
  const userEmail = document.getElementById('email');
  const userPassword = document.getElementById('password');
  const userConfirmPassword = document.getElementById('confirn-password');
  const passwords = document.getElementsByClassName('password');
  const submitButton = document.querySelector('.form__btn');


  const checkUserName = () => {
    let valid = false;

    let invalidSymbols = ['`', '~', '!', '@', '"', '#', '№', '$', ';', '%', '^', ':', '&', '?', '*', '(', ')',
                          '+', '=', '{', '[', ']', '}', ';', `'`, `\\`, '|', '/', ',', '<', '>'];
    let invalidSymbol = '';

    outer: 
    for (let i = 0; i < userName.value.length; i++) {
      for (let n = 0; n < invalidSymbols.length; n++) {
        if (userName.value[i] == invalidSymbols[n]) {
          if (!invalidSymbol.includes(invalidSymbols[n])) {
            invalidSymbol += `${invalidSymbols[n]} `;
          }
          continue outer;
        } 
      }
    }

    let dotCount = 0;
    let underliningCount = 0;
    let dashCount = 0;
    let foundSymbol;

    for (let i = 0; i < userName.value.length; i++) {

      if (userName.value[i] == '.' && userName.value[i + 1] == '.') {
        dotCount++;
        foundSymbol = '.';
      } else if (userName.value[i] == '-' && userName.value[i + 1] == '-') {
          dashCount++;
          foundSymbol = '-';
      } else if (userName.value[i] == '_' && userName.value[i + 1] == '_') {
          underliningCount++;
          foundSymbol = '_';
      }

    }

    if (userName.value.length != 0) {

      if (invalidSymbol != '') {
        setError(userName, `Логин содержит недопустимые символы: ${invalidSymbol}`);
      } else if (dotCount >= 1 || dashCount >= 1 || underliningCount >= 1) {
          setError(userName, `Логин не должен содержать два или более \nсимвола ' ${foundSymbol} ' подряд`);
      } else if (userName.value.length < 6 || userName.value.length > 20) {
          setError(userName, 'Логин должен содержать от 6 до 20 символов');
      } else {
          setSuccess(userName);
          valid = true;
      }

    } else {
        setEmpty(userName);
    }
      
    return valid;
  };

  const checkUserEmail = () => {
    let valid = false;

    if (userEmail.value.length != 0) {

      if (userEmail.value.length < 8 || userEmail.value.length > 32) {
        setError(userEmail, 'Email должен содержать от 8 до 32 символов');
      } else {
          setSuccess(userEmail);
          valid = true;
      }

    } else {
        setEmpty(userEmail);
    }

    return valid;
  };

  const checkUserPassword = () => {
    let valid = false;

    for (let i = 0; i < passwords.length; i++) {
      if (userPassword.value.length != 0) {

        if (userPassword.value.length < 6 || userPassword.value.length > 32) {
          setError(userPassword, 'Пароль должен содержать от 6 до 32 символов');
          setEmpty(userConfirmPassword);
        } else if ((userPassword.value != userConfirmPassword.value) && userConfirmPassword.value.length != 0) {
            setError(passwords[i], 'Пароли не совпадают');
        } else if (userPassword.value.length == 0) {
            setEmpty(userConfirmPassword);
        } else if (userPassword.value == userConfirmPassword.value) {
            setSuccess(passwords[i]);
            valid = true;
        } else {
            setEmpty(userPassword);
        }

      } else {
          setEmpty(userPassword);
          setEmpty(userConfirmPassword);
      }
    }
        
    return valid;
  };

  const checkUserConfirmPassword = () => {
    let valid = false;

    for (let i = 0; i < passwords.length; i++) {
      if (userConfirmPassword.value.length != 0) {

        if (userPassword.value.length < 6 || userPassword.value.length > 32) {
          setEmpty(userConfirmPassword);
        } else if ((userConfirmPassword.value != userPassword.value) && userPassword.value.length != 0) {
            setError(passwords[i], 'Пароли не совпадают');         
        } else if (userPassword.value.length == 0) {
            setEmpty(userConfirmPassword);
        } else if (userConfirmPassword.value == userPassword.value) {
            setSuccess(passwords[i]);
            valid = true;
        } else {
            setEmpty(userConfirmPassword);
        }

      } else {
          setEmpty(userConfirmPassword);
      }
    } 

    return valid;
  };

  function setError(input, message) {
    const inputBox = input.parentElement;
    const errorMessage = inputBox.querySelector('.error__msg');

    errorMessage.innerText = message;

    inputBox.className = 'form-input__box error';
  }

  const setSuccess = input => {
    const inputBox = input.parentElement;

    inputBox.className = 'form-input__box success';

    return true;
  }

  const setEmpty = input => {
    const inputBox = input.parentElement;

    inputBox.className = 'form-input__box';
  }

  const showHidePassword = () => {
    const passButton = document.getElementsByClassName('showPassword');
    const passConfirmButton = document.getElementsByClassName('showConfirmPassword');

    for (let button of passButton) {
      button.addEventListener('click', () => {

        if (userPassword.getAttribute('type') === 'password') {
          userPassword.setAttribute('type', 'text');
        } else {
            userPassword.setAttribute('type', 'password');
        }

      });
    }
    
    for (let button of passConfirmButton) {
      button.addEventListener('click', () => {

        if (userConfirmPassword.getAttribute('type') === 'password') {
          userConfirmPassword.setAttribute('type', 'text');
        } else {
            userConfirmPassword.setAttribute('type', 'password');
        }

      });
    }
  };

  form.addEventListener('input', (e) => {
    let input = e.target;

    if (input.id == 'name') {
      checkUserName();
    } else if (input.id == 'email') {
        checkUserEmail();
    } else if (input.id == 'password') {
        checkUserPassword();
    } else if (input.id == 'confirn-password') {
        checkUserConfirmPassword();
    }

    if (checkUserName() && checkUserEmail() && checkUserPassword() && checkUserConfirmPassword()) {
      submitButton.removeAttribute('disabled');
    } else {
        submitButton.setAttribute('disabled', 'true');
    }

  });

  let resultJSON;

  const sendData = async (url, data) => {
    const response = await fetch(url, {
      method: 'POST',
      body: data,
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
        },
    });

    if (!response.ok) {
      throw new Error( `Ошибка по даресу ${url}, статус ошибки: ${response.status}` )
    }

    resultJSON = await response.json();

    console.log( resultJSON );

    return resultJSON;
  };

  const sendForm = () => {
    form.addEventListener('submit', e => {
      e.preventDefault();

      let isUserNameValid = checkUserName();
      let isEmailValid = checkUserEmail();
      let isPasswordValid = checkUserPassword();
      let isConfirmPasswordValid = checkUserConfirmPassword();

      let isFormValid = isUserNameValid && isEmailValid && isPasswordValid && isConfirmPasswordValid;

      setEmpty(userName);
      setEmpty(userEmail);
      setEmpty(userPassword);
      setEmpty(userConfirmPassword);

      const formData = new FormData(form);
      const data = {};

      formData.delete('password-check');

      for (const [key, value] of formData) {
        data[key] = value;
      }

      if (isFormValid) {
        sendData('/save_user', JSON.stringify(data))
          .then(() => {
            if (resultJSON.checkEmail == true) {
              setError(userEmail, 'Этот Email уже занят');
              checkUserName();
              checkUserPassword();
              checkUserConfirmPassword();
            } else {
                setSuccess(userName);
                setSuccess(userEmail);
                setSuccess(userPassword);
                setSuccess(userConfirmPassword);
                window.location.href = '/';
            }
          })
          .catch((err) => {
            console.log(err);
          });
        }
    });
  };

  showHidePassword();
  sendForm();
});