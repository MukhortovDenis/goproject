window.addEventListener('DOMContentLoaded', () => {
  
  const sendData = async (url, data) => {
    const response = await fetch(url, {
      method: 'POST',
      body: data,
    });

    if (!response.ok) {
      throw new Error( `Ошибка по даресу ${url}, статус ошибки: ${response.status}` )
    }

    return await response.json();
  };

  const sendForm = () => {
    const form = document.getElementById('form-login');

    form.addEventListener('submit', e => {
      e.preventDefault();

      const successCount = checkInputs();

      const formData = new FormData(form);

      const data = {};

      for (const [key, value] of formData) {
        data[key] = value;
      }

      if (successCount == 2) {
        sendData('/check_user', JSON.stringify(data))
          .then(() => {
            window.location.href = '/check';
            // setTimeout(() => window.location.href = '/', 2000); // на случай "медленных интернетов"
            form.reset();
          })
          .catch((err) => {
            console.log(err);
          });
        }
    });
  };

  // form.onsubmit = async (e) => {
  //   e.preventDefault();

  //   const successCount = checkInputs();

  //   const formData = new FormData(form);
  //   const plainFormData = Object.fromEntries(formData.entries());

  //   if (successCount === 2) {
  //     let response = await fetch(requestURL, {
  //       method: 'POST',
  //       // headers: {
  //       //   'Content-Type': 'application/json'
  //       // },
  //       body: JSON.stringify(plainFormData)
  //     });

  //     let formDataJSON = JSON.stringify(plainFormData);

  //     console.log( plainFormData );
  //     console.log( formDataJSON );

  //     // form.reset();

  //     // window.location.href = '/';

  //     setTimeout(() => window.location.href = '/', 10000); // на случай "медленных интернетов"

  //     return await response.json();
  //   }
  // };

  function checkInputs() {
    const userEmailValue = userEmail.value.trim();
    const userPasswordValue = userPassword.value.trim();

    let success = 0;

    if (userEmailValue === '') {
      setErrorFor(userEmail, 'Это поле не должно быть пустым');
    } else if (userEmailValue.length < 8) {
        setErrorFor(userEmail, 'Email слишком короткий');
    } else if (userEmailValue.length > 32) {
        setErrorFor(userEmail, 'Email слишком длинный');
    } else {
        setSuccessFor(userEmail);
        success++;
    }

    if (userPasswordValue === '') {
      setErrorFor(userPassword, 'Это поле не должно быть пустым');
    } else if (userPasswordValue.length < 6 || userPasswordValue.length > 32) {
        setErrorFor(userPassword, 'Длина пароля должна быть от 6 до 32 символов');
    } else {
        setSuccessFor(userPassword);
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


  const userEmail = document.getElementById('email');
  const userPassword = document.getElementById('password');

  sendForm();
});