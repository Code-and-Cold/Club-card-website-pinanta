<script setup>
import HeroSection from './components/HeroSection.vue'
import FeedbackSection from './components/FeedbackSection.vue'
</script>

<template>
  <head>
    <title>Club Frontend Prototype</title>
  </head>

  <body>
    <div class="main">
      <HeroSection />
      <FeedbackSection />
    </div>
  </body>
</template>

<style>
/* Стиль для всей страницы, поэтому не scoped. В компонентах scoped. */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;

  /* TODO: шрифты не везде такие, вынести по готовности тестов */
  font-family: 'JetBrains Mono';
  font-style: normal;
}

/* TODO: Вынести в HeroSection как будут готовы тесты */
h1 {
  text-align: center;
  color: white;
}

.main {
  background-image: url('./assets/images/Space-Background-Images.jpg');
  background-attachment: fixed;
}

/* TODO: Вынести в компоненты или отдельный стилевой файл как будут готовы тесты */
.section {
  width: 100%; /* Немного адаптива, проверяйте mobile размеры через инструменты разработчика */
  min-width: 330px; /* При таком размере остаётся читаемым текст формы feedback, остальное не проверено */
  padding: 50px 20px;
}
</style>
