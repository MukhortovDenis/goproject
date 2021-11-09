window.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('form-login');
  const userEmail = document.getElementById('email');
  const userPassword = document.getElementById('password');
  const requestURL = '/check_user';


  form.onsubmit = async (e) => {
    e.preventDefault();

    const successCount = checkInputs();

    const formData = new FormData(form);
    const plainFormData = Object.fromEntries(formData.entries());
    const fromDataJSON = JSON.stringify(plainFormData);

    if (successCount === 2) {
      let response = await fetch(requestURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: fromDataJSON
        });

      let result = await response.json();

      form.reset();

      window.location.href = '/';

      // setTimeout(() => window.location.href = '/', 500); на случай "медленных интернетов"
    }

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
  };
});