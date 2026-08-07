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
import CookieMessageModal from './components/CookieMessageModal.vue'

const team = [
  {
    id: 1,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'За Лупкин',
    position: 'Директор',
  },
  {
    id: 2,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'Яша Лава',
    position: 'Уже уволен',
  },
  {
    id: 3,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'Спайдер мент',
    position: 'Офицер',
  },
  {
    id: 4,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'Пупу Пу',
    position: 'ОАО',
  },
  {
    id: 5,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    name: 'ААААААААААААААА',
    position: 'ААААААААА',
  },
]

const recentEvents = [
  {
    id: 1,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    fullText:
      'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза. Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 2,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 3,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 4,
    photo: new URL('./assets/images/cheremsha.png', import.meta.url).href,
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
      <CookieMessageModal />
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
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  background-attachment: fixed;
}
</style>
