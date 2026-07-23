<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  document.querySelector('.feedback-section__form').addEventListener('submit', async (e) => {
    e.preventDefault()

    // Собираем данные
    const name = document.getElementById('name').value
    //const department = document.getElementById('department').value;
    //const course = document.getElementById('course').value;
    //const link = document.getElementById('link').value;
    const agreement = document.querySelector('.form__checkbox-input').checked

    const email = 'ivan@example.com' // FIXME: Только для текущего API, удалить по изменению

    // Элемент для сообщения
    let msg = document.querySelector('.form__message')
    if (!msg) {
      msg = document.createElement('p')
      msg.className = 'form__message'
      document.querySelector('.form').appendChild(msg)
    }

    // Валидация
    if (!name.trim()) {
      msg.style.color = 'red'
      msg.textContent = 'Пожалуйста, введите ФИО'
      return
    }

    if (!agreement) {
      msg.style.color = 'red'
      msg.textContent = 'Необходимо дать согласие на обработку данных'
      return
    }

    msg.textContent = 'Отправка...'
    msg.style.color = '#3bb0e3'

    try {
      /* FIXME: использовать полные данные не позволяет API, но вот пример запроса со всеми данными:
      const res = await fetch('/api/apply', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
          name, 
          department, 
          course, 
          link: link || '',
          agreement 
        })
      }); */

      // FIXME: удалить по обновлнию API:
      const res = await fetch('/api/apply', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email }),
      })

      if (!res.ok) throw new Error()

      msg.style.color = 'green'
      msg.textContent = '✅ Заявка отправлена! С вами свяжутся'
      e.target.reset()

      // Сброс чекбокса
      document.querySelector('.form__checkbox-input').checked = false
    } catch {
      msg.style.color = 'red'
      msg.textContent = '❌ Ошибка отправки. Попробуйте ещё раз.'
    }
  })
})
</script>

<template>
  <head>
    <title>Club Frontend Prototype</title>
  </head>

  <body>
    <section class="section feedback-section">
      <div class="feedback-section__content">
        <h2 class="feedback-section__title">Хочешь тусить с нами?</h2>

        <h3 class="feedback-section__subtitle">Заполняй анкету и с тобой свяжутся наши шестерки</h3>

        <form class="form feedback-section__form" action="#" method="POST">
          <div class="form__field">
            <input
              class="form__input"
              type="text"
              id="name"
              placeholder="ФИО"
              required
              pattern="^[А-ЯЁ][а-яё]+\s[А-ЯЁ][а-яё]+(?:\s[А-ЯЁ][а-яё]+)?$"
            />
          </div>

          <div class="form__field">
            <select class="form__select" id="department" required>
              <option class="form__select-item" value="" disabled selected>Высшая школа</option>
              <option class="form__select-item" value="1">ВШ 1</option>
              <option class="form__select-item" value="2">ВШ 2</option>
              <option class="form__select-item" value="3">ВШ 3</option>
            </select>
          </div>

          <div class="form__field">
            <select class="form__select" id="course" required>
              <option class="form__select-item" value="" disabled selected>Курс</option>
              <option class="form__select-item" value="1">1</option>
              <option class="form__select-item" value="2">2</option>
              <option class="form__select-item" value="3">3</option>
              <option class="form__select-item" value="4">4</option>
              <option class="form__select-item" value="5">5</option>
              <option class="form__select-item" value="6">6</option>
            </select>
          </div>

          <div class="form__field">
            <input
              class="form__input"
              type="url"
              id="link"
              placeholder="Страница Вконтакте"
              pattern="^(https?:\/\/)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(:\d+)?(\/.*)?$"
            />
            <!-- Здесь нужно прописать required для работы pattern -->
          </div>

          <div class="form__checkbox-panel">
            <label class="form__checkbox-label">
              <input class="form__checkbox-input" type="checkbox" id="policies" />
              <span class="form__checkmark"></span>
              <span class="form__checkbox-text">
                Даю <a class="form__link" href="https://www.example.com">согласие</a> на обработку
                <a class="form__link" href="https://www.example.com">персональных данных</a>
              </span>
            </label>
          </div>

          <button class="form__submit-button" type="submit">Вступить в клуб</button>
        </form>

        <p class="feedback-section__footer">
          Или напиши на почту руководителю: <br />
          <a class="feedback-section__link" href="mailto:ivan@cold-code.ru">ivan@cold-code.ru</a>
        </p>
      </div>

      <div class="feedback-section__media">
        <img class="feedback-section__image" src="/src/assets/images/cheremsha.png" alt="Черемша" />
      </div>
    </section>

    <!-- API -->
    <!--
	const res = await fetch('/api/apply', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, email })
    });-->
  </body>
</template>

<style>
/* FIXME: используйте <style scoped> как только перейдёте на отдельные css файлы */

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

h1 {
  text-align: center;
  color: white;
}

.section {
  width: 100%; /* Немного адаптива, проверяйте mobile размеры через инструменты разработчика */
  min-width: 330px; /* При таком размере остаётся читаемым текст формы feedback, остальное не проверено */
  padding: 50px 20px;
}

.feedback-section {
  background-color: #002f55;

  color: white;

  display: flex;
  flex-direction: column;
  align-items: justify;
  gap: 30px;
}

.feedback-section__content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
}

.feedback-section__title {
  font-family: JetBrains Mono;
  font-size: 35px;
  font-weight: 400; /* Regular */
  line-height: 1.1;
}

.feedback-section__subtitle {
  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form__field {
  width: 100%;
  height: 65px;

  border-radius: 20px;
  border: 2px solid #3bb0e3;

  background-color: #003b6b;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400;
  line-height: 1.2;
  color: white;

  display: flex;
  align-items: center;
  padding: 0 18px;
}

.form__input {
  flex: 1;
  width: 100%;
  height: 100%;

  background: transparent;
  border: none;
  outline: none;

  color: white;
  font: inherit;
}

.form__select {
  flex: 1;
  width: 100%;
  height: 100%;

  background: transparent;
  border: none;
  outline: none;

  color: white;
  font: inherit;

  cursor: pointer;

  appearance: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;

  background-image: url('data:image/svg+xml,...'); /* кастомная стрелка */
  cursor: pointer;
}

.form__select-item {
  flex: 1;
  width: 100%;
  height: 100%;
  background-color: #003b6b;
}

input {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  user-select: none;
}

.form__input::placeholder {
  color: #3bb0e3; /* Forms placeholder font color */
  opacity: 1;
}

.form__checkbox-panel {
  width: 100%;
  height: 65px;
}

.form__checkbox-label {
  flex: 1;
  width: 100%;
  height: 100%;

  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 15px;

  font-family: Inter;
  font-size: 16px;
  font-weight: 400; /* Regular */
  line-height: 1.5;
  color: #3bb0e3;
}

.form__checkbox-element {
  height: 35px;
  width: 35px;

  background-color: red;
}

.form__checkbox-input {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  border: 0;
}

.form__link {
  color: #3bb0e3;
}

@media (hover: hover) {
  .form__link:hover {
    color: #5cc5f5;
  }
}

.form__checkmark {
  flex-shrink: 0;
  width: 35px;
  height: 35px;
  border-radius: 10px;
  border: 2px solid #3bb0e3;
  background-color: #003b6b;
  display: flex;
  align-items: center;
  justify-content: center;
  transition:
    background-color 0.2s,
    border-color 0.2s;
}

.form__checkmark::after {
  content: '✓';
  font-size: 24px;
  font-weight: 700;
  color: white;
  opacity: 0;
  transform: scale(0.5);
  transition:
    opacity 0.2s,
    transform 0.2s;
}

/* On mouse-over, add a grey background color; desktop only */
@media (hover: hover) {
  .form__checkbox-label:hover .form__checkmark {
    background-color: rgba(59, 176, 227, 0.3);
  }

  .form__checkbox-label:hover .form__checkbox-input:checked + .form__checkmark {
    background-color: #1a7fc9;
    border-color: #1a7fc9;
  }
}

/* When the checkbox is checked, add a blue background */
.form__checkbox-input:checked + .form__checkmark {
  background-color: #2196f3;
  border-color: #2196f3;
}

/* Create the checkmark/indicator (hidden when not checked) */
.checkmark:after {
  /*
  content: '';
  position: absolute;
  display: none;
  */
}

/* Show the checkmark when checked */
.form__checkbox-input:checked + .form__checkmark::after {
  opacity: 1;
  transform: scale(1);
}

.form__checkbox-text {
  flex: 1;
}

.form__submit-button {
  width: fit-content;
  height: fit-content;

  border-radius: 20px;
  border: 0px solid #e3953b;

  padding: 15px 25px;

  background-color: #e3953b;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: white;
}

.feedback-section__footer {
  height: fit-content;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: white;
}

.feedback-section__link {
  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: #e3953b;
}

.feedback-section__media {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 30px;
}

.feedback-section__image {
  width: 335px;
  border-radius: 30px;
  border: 0px solid #e3953b;
}
</style>
