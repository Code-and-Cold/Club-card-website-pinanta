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
@import url('./assets/global.css');

.main {
  background-image: url('./assets/images/Space-Background-Images.jpg');
  background-attachment: fixed;
}
</style>
