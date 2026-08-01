<script setup>
import ClubAdvantages from './ClubAdvantages.vue'
import MemberBenefits from './MemberBenefits.vue'

const advantages = [
  {
    id: 1,
    title: 'У нас есть черешня',
    accent: 'Ммного черешни',
    text: 'Потому что у нас много своей черешни — только у нас.',
  },
  {
    id: 2,
    title: 'У нас есть черешня',
    accent: 'Очень много черешни',
    text: 'Потому что у нас много своей черешни — только у нас.',
  },
  {
    id: 3,
    title: 'У нас есть Пиво',
    accent: 'Очень много',
    text: 'Что вам еще нужно?',
  },
]

const memberBenefits = [
  {
    id: 1,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
  },
  {
    id: 2,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
  },
  {
    id: 3,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
  },
  {
    id: 4,
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
  },
]
</script>

<template>
  <div class="advantages-benefits-sections">
    <ClubAdvantages :items="advantages" />
    <MemberBenefits :items="memberBenefits" />
  </div>
</template>

<style scoped>
.advantages-benefits-sections {
  --page-bg: #f3f3f3;
  --card-bg: #e6e6e6;
  --text-main: #073b63;
  --accent-blue: #27b7ed;
  --accent-orange: #e9931d;

  color: var(--text-main);
  background-color: var(--page-bg);
}
</style>
