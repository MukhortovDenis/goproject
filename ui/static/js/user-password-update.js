window.addEventListener('DOMContentLoaded', () => {

  let resultJSON;

  const sendData = async (url, data) => {
    const response = await fetch(url, {
      method: 'PATCH',
      body: data,
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
    })
  

    if (!response.ok) {
      throw new Error( `Ошибка по даресу ${url}, статус ошибки: ${response.status}` )
    }

    resultJSON = await response.json();

    console.log( resultJSON );

    return resultJSON;
  };

  const sendForm = () => {
    const form = document.getElementById('form-change-password');

    console.log( form );

    form.addEventListener('submit', e => {
      e.preventDefault();

      const successCount = checkInputs();
      const formData = new FormData(form);
      const data = {};

      for (const [key, value] of formData) {
        data[key] = value;
      }

      if (successCount == 2) {
        sendData('/cabinet-password-change', JSON.stringify(data))
          .then(() => {
            if (resultJSON.checkPass == true) {
              setErrorFor(userPassword, 'Текущий пароль введён неверно');
            } else {
                setSuccessFor(userPassword);
                setSuccessFor(userNewPassword);
                window.location.href = '/cabinet-password';
            }
          })
          .catch((err) => {
            console.log(err);
          });
        }
    });
  };

  function checkInputs() {
    const userPasswordValue = userPassword.value.trim();
    const userNewPasswordValue = userNewPassword.value.trim();

    let success = 0;

    if (userPasswordValue === '') {
      setErrorFor(userPassword, 'Это поле не должно быть пустым');
    } else if (userPasswordValue.length < 6 || userPasswordValue.length > 32) {
        setErrorFor(userPassword, 'Длина пароля должна быть от 6 до 32 символов');
    } else if (!isPassword(userPasswordValue)) {
        setErrorFor(userPassword, 'Пароль введён некоректно');
    } else {
        setEmptyFor(userPassword);
        success++;
    }

    if (userNewPasswordValue === '') {
      setErrorFor(userNewPassword, 'Это поле не должно быть пустым');
    } else if (userNewPasswordValue.length < 6 || userNewPasswordValue.length > 32) {
        setErrorFor(userNewPassword, 'Длина пароля должна быть от 6 до 32 символов');
    } else if (!isPassword(userNewPasswordValue)) {
        setErrorFor(userNewPassword, 'Пароль введён некоректно');
    } else if (userNewPasswordValue == userPasswordValue) {
        setErrorFor(userNewPassword, 'Новый пароль совпадает с текущим');
    } else {
        setEmptyFor(userNewPassword);
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

  function setEmptyFor(input) {
    const inputBox = input.parentElement;

    inputBox.className = 'form-input__box empty';
  }

  function isPassword(password) {
    return /^(?=.*[a-zA-Z0-9]).{6,32}$/.test(password);
  }

  const userPassword = document.getElementById('password');
  const userNewPassword = document.getElementById('new-password');

  sendForm();
});
