<script setup>
import MemberBenefits from './MemberBenefits.vue'

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
  <div class="member-benefits-section">
    <MemberBenefits :items="memberBenefits" />
  </div>
</template>

<style scoped>
.member-benefits-section {
  --page-bg: #f3f3f3;
  --text-main: #073b63;
  --accent-blue: #27b7ed;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
