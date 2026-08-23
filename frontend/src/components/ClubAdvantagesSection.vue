<script setup>
import { onMounted, ref } from 'vue'
import { getSiteData } from '../api.js'
import ClubAdvantages from './ClubAdvantages.vue'
import networkIntelIcon from '../assets/icons/advantages/network-intel-node.svg'
import handshakeIcon from '../assets/icons/advantages/handshake-rounded.svg'
import heartSmileIcon from '../assets/icons/advantages/heart-smile-outline-rounded.svg'

const fallbackIcons = [networkIntelIcon, handshakeIcon, heartSmileIcon]

const title = ref('Что тебя ждет в клубе?')

const advantages = ref([
  {
    id: 1,
    icon: networkIntelIcon,
    title: 'Разработка реальных проектах',
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
    icon: heartSmileIcon,
    title: 'Свое комьюнити',
    text: 'У нас можно поспорить о том, какой язык лучше, найти напарника для пет-проекта или выдохнуть после пар в компании единомышленников.',
  },
])

onMounted(async () => {
  try {
    const siteData = await getSiteData()

    if (siteData?.why_us?.title) {
      title.value = siteData.why_us.title
    }

    if (Array.isArray(siteData?.why_us_cards) && siteData.why_us_cards.length) {
      advantages.value = siteData.why_us_cards.map((card, index) => ({
        id: card.id ?? index,
        icon: card.icon_url || fallbackIcons[index] || networkIntelIcon,
        title: card.title ?? '',
        text: card.description ?? '',
      }))
    }
  } catch (error) {
    console.warn('Не удалось загрузить блок преимуществ с backend.', error)
  }
})
</script>

<template>
  <div class="club-advantages-section" id="advantages-section">
    <ClubAdvantages :title="title" :items="advantages" />
  </div>
</template>

<style scoped>
.club-advantages-section {
  --page-bg: #f0f0f1;
  --card-bg: #e7e7e7;
  --text-main: #002f55;
  --accent-blue: #3bb0e3;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
