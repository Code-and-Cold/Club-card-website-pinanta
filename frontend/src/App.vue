<script setup>
import HeroSection from './components/HeroSection.vue'
import FeedbackSection from './components/FeedbackSection.vue'
import PhraseSection from './components/PhraseSection.vue'
import PhraseBottomSection from './components/PhraseBottomSection.vue'
import AdvantagesBenefitsSections from './components/AdvantagesBenefitsSections.vue'
import Team from './components/Team.vue'
import RecentEvents from './components/RecentEvents.vue'

const team = [
  {
    id: 1,
    photo: "src/assets/images/1.png",
    name: 'За Лупкин',
    position: 'Директор',
  },
  {
    id: 2,
    photo: "src/assets/images/1.png",
    name: 'Яша Лава',
    position: 'Уже уволен',
  },
  {
    id: 3,
    photo: "src/assets/images/3.jpg",
    name: 'Спайдер мент',
    position: 'Офицер',
  },
  {
    id: 4,
    photo: "src/assets/images/4.png",
    name: 'Пупу Пу',
    position: 'ОАО',
  },
  {
    id: 5,
    photo: "src/assets/images/5.png",
    name: 'ААААААААААААААА',
    position: 'ААААААААА',
  },
]

const recentEvents = [
  {
    id: 1,
    photo: "src/assets/images/1.png",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    fullText: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 2,
    photo: "src/assets/images/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 3,
    photo: "src/assets/images/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 4,
    photo: "src/assets/images/6.jpg",
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
      <AdvantagesBenefitsSections />
      <Team :items="team" />
      <PhraseBottomSection />
      <RecentEvents :items="recentEvents" />
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
