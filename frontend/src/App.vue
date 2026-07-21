<script setup></script>

<template>
  <head>
    <title>Club Frontend Prototype</title>
  </head>

  <body>
    <div class="test">
      <h1>Здесь будет наш свёрстанный сайт, нужно только подождать...</h1>
    </div>

    <div class="feedback-section">
      <div class="feedback-section__content">
        <h2 class="feedback-section__title">Хочешь тусить с нами?</h2>

        <h3 class="feedback-section__subtitle">Заполняй анкету и с тобой свяжутся наши шестерки</h3>

        <form class="form feedback-section__form" action="#" method="POST">
          <div class="form__field">
            <input class="feedback-section__form__input" type="text" id="name" placeholder="ФИО" />
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
          Или напиши на почту руководителю:
          <a class="feedback-section__link" href="mailto:ivan@cold-code.ru">ivan@cold-code.ru</a>
        </p>
      </div>

      <div class="feedback-section__media">
        <img class="feedback-section__image" src="" alt="Черемша" />
      </div>
    </div>

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

<style scoped>
* {
  padding: 0;
  margin: 0;
}

h1 {
  text-align: center;
  color: white;
}

div.test {
  background-color: #002a4c;
  width: 315px;
  height: 300px;
}
</style>
