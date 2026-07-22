<script setup></script>

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
            <input class="form__input" type="text" id="name" placeholder="ФИО" />
          </div>

          <select class="form__select" id="department">
            <option class="form__select-item" value="" disabled selected>Высшая школа</option>
            <option class="form__select-item" value="1">ВШ 1</option>
            <option class="form__select-item" value="2">ВШ 2</option>
            <option class="form__select-item" value="3">ВШ 3</option>
          </select>

          <select class="form__select" id="course">
            <option class="form__select-item" value="" disabled selected>Курс</option>
            <option class="form__select-item" value="1">1</option>
            <option class="form__select-item" value="2">2</option>
            <option class="form__select-item" value="3">3</option>
            <option class="form__select-item" value="4">4</option>
            <option class="form__select-item" value="5">5</option>
            <option class="form__select-item" value="6">6</option>
          </select>

          <label class="form__checkbox-label">
            <input class="form__checkbox" type="checkbox" />
            Даю <a class="form__link" href="https://www.example.com">согласие</a> на обработку
            <a class="form__link" href="https://www.example.com">персональных данных</a>
          </label>

          <div class="form__field">
            <input class="form__input" type="text" id="link" placeholder="Страница Вконтакте" />
          </div>

          <div class="form__field">
            <input class="form__input" type="checkbox" id="name" placeholder="ФИО" />
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

    <br /><br /><br /><br /><br />

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

.form__input {
  padding: 30px 18px;
  height: 65px;

  border-radius: 20px;
  border: 2px solid #3bb0e3;

  background-color: #003b6b;
  color: white;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
}

.form__select {
  padding: 30px 18px;
  height: 65px;

  border-radius: 20px;
  border: 2px solid #3bb0e3;

  background-color: #003b6b;
  color: white;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
}

::placeholder {
  color: #3bb0e3; /* Forms placeholder */
  opacity: 1; /* Firefox */
}
</style>
