<script setup>
import ClubAdvantages from './ClubAdvantages.vue'
import brainIcon from '../assets/vector/advantage_brain.svg'
import handshakeIcon from '../assets/vector/advantage_handshake.svg'
import heartIcon from '../assets/vector/advantage_heart.svg'

const advantages = [
  {
    id: 1,
    icon: brainIcon,
    title: 'Разработка реальных проектов',
    text: 'Ошибаемся, гуглим, чиним и получаем опыт командной разработки. Точно будет что рассказать на собеседовании!',
  },
  {
    id: 2,
    icon: handshakeIcon,
    title: 'Менторство и помощь',
    text: 'У нас всегда можно задать «глупый» вопрос и получить адекватный ответ. Опытные ребята подскажут, как сдать сессию или сделать свой первый проект.',
  },
  {
    id: 3,
    icon: heartIcon,
    title: 'Свое комьюнити',
    text: 'У нас можно поспорить о том, какой язык лучше, найти напарника для пет-проекта или выдохнуть после пар в компании единомышленников.',
  },
]
</script>

<template>
  <div class="club-advantages-section" id="advantages-section">
    <ClubAdvantages :items="advantages" />
  </div>
</template>

<style scoped>
.club-advantages-section {
  --page-bg: #f3f3f3;
  --card-bg: #e6e6e6;
  --text-main: #073b63;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
