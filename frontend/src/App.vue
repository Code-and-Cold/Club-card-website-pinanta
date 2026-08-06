<script setup>
import HeroSection from './components/HeroSection.vue'
import FeedbackSection from './components/FeedbackSection.vue'
import PhraseSection from './components/PhraseSection.vue'
import PhraseBottomSection from './components/PhraseBottomSection.vue'
import ClubAdvantagesSection from './components/ClubAdvantagesSection.vue'
import MemberBenefitsSection from './components/MemberBenefitsSection.vue'
import TeamSection from './components/TeamSection.vue'
import RecentEvents from './components/RecentEvents.vue'
import FooterSection from './components/FooterSection.vue'

const team = [
  {
    id: 1,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'За Лупкин',
    position: 'Директор',
  },
  {
    id: 2,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    name: 'Яша Лава',
    position: 'Уже уволен',
  },
  {
    id: 3,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    name: 'Спайдер мент',
    position: 'Офицер',
  },
  {
    id: 4,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    name: 'Пупу Пу',
    position: 'ОАО',
  },
  {
    id: 5,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    name: 'ААААААААААААААА',
    position: 'ААААААААА',
  },
]

const recentEvents = [
  {
    id: 1,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    fullText:
      'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 2,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 3,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 4,
<<<<<<< Updated upstream
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
=======
    photo: '/src/assets/images/cheremsha.png',
>>>>>>> Stashed changes
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
]
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
      <TeamSection :items="team" />
      <MemberBenefitsSection />
      <PhraseBottomSection />
      <RecentEvents :items="recentEvents" />
      <FeedbackSection />
      <FooterSection />
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
