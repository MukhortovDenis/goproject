window.addEventListener('DOMContentLoaded', () => {

  let resultJSON;

  const sendData = async (url, data) => {
    const response = await fetch(url, {
      method: 'PATCH',
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
    const form = document.getElementById('form-user-info');

    form.addEventListener('submit', e => {
      e.preventDefault();

      const successCount = checkInputs();
      const formData = new FormData(form);
      const data = {};

      for (const [key, value] of formData) {
        data[key] = value;
      }

      if (successCount == 2) {
        sendData('/cabinet-info-reset', JSON.stringify(data))
          .then(() => {
            if (resultJSON.checkEmail == true) {
              setErrorFor(userEmail, 'Этот email уже занят');
            } else {
                setSuccessFor(userName);
                setSuccessFor(userEmail);
                window.location.href = '/cabinet-info';
            }
          })
          .catch((err) => {
            console.log(err);
          });
        }
    });
  };

  function checkInputs() {
    const userNameValue = userName.value.trim();
    const userEmailValue = userEmail.value.trim();

    console.log( userNameValue );
    console.log( userEmailValue );

    let success = 0;

    if (userNameValue === '') {
      setErrorFor(userName, 'Это поле не должно быть пустым');
    } else if (userNameValue.length < 5 || userNameValue.length > 20) {
        setErrorFor(userName, 'Длина логина должна быть от 5 до 20 символов');
    } else if (!isName(userNameValue)) {
        setErrorFor(userName, 'Введены некорректные символы');
    } else {
        setEmptyFor(userName)
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
        setEmptyFor(userEmail);
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

  function isName(name) {
    return /^[a-zA-Z0-9]{5,20}$/.test(name);
  }

  function isEmail(email) {
    return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,4}))$/.test(email);
  }

  const userName = document.getElementById('name');
  const userEmail = document.getElementById('email');

  sendForm();
});