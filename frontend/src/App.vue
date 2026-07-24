<script setup>
  import Hero from './components/Hero.vue';
  import FeedbackSection from './components/FeedbackSection.vue';
</script>

<template>
  <head>
    <title>Club Frontend Prototype</title>
  </head>

  <body>
    <div class="main">
      <Hero />
      <FeedbackSection />
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
  margin: 0;
  padding: 0;
  box-sizing: border-box;

  font-family: 'JetBrains Mono';
  font-style: normal;
}

h1 {
  text-align: center;
  color: white;
}

.main {
  background-image: url('./assets/images/Space-Background-Images.jpg');
  background-attachment: fixed;
}
.section {
  width: 100%; /* Немного адаптива, проверяйте mobile размеры через инструменты разработчика */
  min-width: 330px; /* При таком размере остаётся читаемым текст формы feedback, остальное не проверено */
  padding: 50px 20px;
}
</style>
