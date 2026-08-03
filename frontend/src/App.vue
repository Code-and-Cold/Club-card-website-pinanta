<script setup>
import HeroSection from './components/HeroSection.vue'
import FeedbackSection from './components/FeedbackSection.vue'
import PhraseSection from './components/PhraseSection.vue'
import PhraseBottomSection from './components/PhraseBottomSection.vue'
import ClubAdvantagesSection from './components/ClubAdvantagesSection.vue'
import MemberBenefitsSection from './components/MemberBenefitsSection.vue'
</script>

<template>
  <head>
    <title>Club Frontend Prototype</title>
  </head>

  <body>
    <div class="main">
      <HeroSection />
      <PhraseSection />
      <ClubAdvantagesSection />
      <MemberBenefitsSection />
      <PhraseBottomSection />
      <FeedbackSection />
    </div>
  </body>
</template>

<style>
@import url('./assets/styles/global.css');

.main {
  background-image: url('./assets/images/Space-Background-Images.jpg');
  background-attachment: fixed;
}
</style>
