const cross = document.getElementById('cross');
const crossReg = document.getElementById('crossReg');
const crossFog = document.getElementById('crossFog');
const modalWindow = document.getElementById('modalwindow');
const modalWindow2 = document.getElementById('modalwindow2');
const modalWindowReg = document.getElementById('modalwindowReg');
const modalWindow3 = document.getElementById('modalwindow3');
const modalWindowFog = document.getElementById('modalwindowFog');
const modalWindow4 = document.getElementById('modalwindow4');
const overlay = document.getElementById('modalopened');
const slogan = document.getElementById('slogan');
const startBtn = document.getElementById('startbtn');
const preregBtn = document.getElementById('preregbtn');
const forBtn = document.getElementById('forgot');

function closeModal(){
    modalWindow.style.display = 'none';
    modalWindow.classList.add('closed');
    modalWindow2.classList.add('closed');
    cross.classList.add('closed');
    overlay.classList.remove('modalopened');
    slogan.classList.remove('sloganclosed');
    startBtn.classList.remove('sloganclosed');

}
cross.addEventListener('click', closeModal);
function openModal(){
    modalWindow.style.display = 'flex';
    modalWindow.classList.remove('closed');
    modalWindow2.classList.remove('closed');
    cross.classList.remove('closed');
    overlay.classList.add('modalopened');
    slogan.classList.add('sloganclosed');
    startBtn.classList.add('sloganclosed');
}
function closeModalReg(){
    modalWindowReg.style.display = 'none';
    modalWindowReg.classList.add('closed');
    modalWindow3.classList.add('closed');
    crossReg.classList.add('closed');
    overlay.classList.remove('modalopened');
    slogan.classList.remove('sloganclosed');
    startBtn.classList.remove('sloganclosed');

}
crossReg.addEventListener('click', closeModalReg);
function openModalReg(){
    closeModal();
    modalWindowReg.style.display = 'flex';
    modalWindowReg.classList.remove('closed');
    modalWindow3.classList.remove('closed');
    crossReg.classList.remove('closed');
    overlay.classList.add('modalopened');
    slogan.classList.add('sloganclosed');
    startBtn.classList.add('sloganclosed');
    
}
preregBtn.addEventListener('click', openModalReg);

function openModalFog(){
    closeModal();
    modalWindowFog.style.display = 'flex';
    modalWindowFog.classList.remove('closed');
    modalWindow4.classList.remove('closed');
    crossFog.classList.remove('closed');
    overlay.classList.add('modalopened');
    slogan.classList.add('sloganclosed');
    startBtn.classList.add('sloganclosed');
}
function closeModalFog(){
    modalWindowFog.style.display = 'none';
    modalWindowFog.classList.add('closed');
    modalWindow4.classList.add('closed');
    crossFog.classList.add('closed');
    overlay.classList.remove('modalopened');
    slogan.classList.remove('sloganclosed');
    startBtn.classList.remove('sloganclosed');
}
forBtn.addEventListener('click', openModalFog);
crossFog.addEventListener('click', closeModalFog);





startBtn.addEventListener('click', openModal);

// Получаем элементы формы
const form = document.getElementById('login-form');

// Слушаем событие отправки формы
form.addEventListener('submit', function(event) {
    event.preventDefault(); // Отключаем стандартное поведение отправки формы

    // Получаем значения полей
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Создаем объект с данными
    const userData = {
        username: username,
        password: password
    };

    // Отправляем данные на бэкенд
    console.log('Данные для отправки на бэкенд:', userData);

    // Здесь можно добавить логику отправки данных на сервер, например, с использованием fetch
    /*
    fetch('YOUR_BACKEND_URL', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(userData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Успех:', data);
    })
    .catch(error => {
        console.error('Ошибка:', error);
    });
    */
});
