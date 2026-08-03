<script setup>
import ClubAdvantages from './ClubAdvantages.vue'

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
</script>

<template>
  <div class="club-advantages-section">
    <ClubAdvantages :items="advantages" />
  </div>
</template>

<style scoped>
.club-advantages-section {
  --page-bg: #f3f3f3;
  --card-bg: #e6e6e6;
  --text-main: #073b63;
  --accent-blue: #27b7ed;
  --accent-orange: #e9931d;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
