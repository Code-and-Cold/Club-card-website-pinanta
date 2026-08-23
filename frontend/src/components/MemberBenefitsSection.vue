<script setup>
import { onMounted, ref } from 'vue'
import { getSiteData } from '../api.js'
import MemberBenefits from './MemberBenefits.vue'
import docsIcon from '../assets/icons/benefits/docs-outline.svg'
import boltIcon from '../assets/icons/benefits/bolt-boost-rounded.svg'
import badgeIcon from '../assets/icons/benefits/badge-outline-rounded.svg'
import celebrationIcon from '../assets/icons/benefits/celebration-outline-rounded.svg'

const fallbackIcons = [docsIcon, boltIcon, badgeIcon, celebrationIcon]

const titleLines = ref(['Твой профит', 'от участия в клубе'])

const memberBenefits = ref([
  {
    id: 1,
    icon: docsIcon,
    title: 'Строчка в резюме',
    text: 'Получишь опыт участия в проектах полного цикла — от проектирования до релиза, — и это выделит тебя среди других джунов на рынке.',
  },
  {
    id: 2,
    icon: boltIcon,
    title: 'Прокачка своих навыков',
    text: 'Освоишь современные инструменты и технологии, которые используются в IT-компаниях. А для тех, кому хочется большего — это идеальная среда, чтобы тестировать гипотезы, писать низкоуровневый код и оптимизировать сложные алгоритмы.',
  },
  {
    id: 3,
    icon: badgeIcon,
    title: 'Полезные связи',
    text: 'Познакомишься со специалистами из индустрии, старшекурсниками, которые уже где-то стажируются, и просто амбициозными ребятами из твоего потока.',
  },
  {
    id: 4,
    icon: celebrationIcon,
    title: 'Сохраненные нервы',
    text: 'Тебе не придется часами гуглить непонятную ошибку в одиночку. В клубе есть у кого спросить совет по архитектуре, попросить код-ревью или просто узнать, как сдать сложный предмет.',
  },
])

onMounted(async () => {
  try {
    const siteData = await getSiteData()

    if (siteData?.benefits?.title) {
      titleLines.value = siteData.benefits.title.split('\n')
    }

    if (Array.isArray(siteData?.benefits_items) && siteData.benefits_items.length) {
      memberBenefits.value = siteData.benefits_items.map((item, index) => ({
        id: item.id ?? index,
        icon: item.icon_url || fallbackIcons[index] || docsIcon,
        title: item.title ?? '',
        text: item.description ?? '',
      }))
    }
  } catch (error) {
    console.warn('Не удалось загрузить блок преимуществ участника с backend.', error)
  }
})
</script>

<template>
  <div class="member-benefits-section" id="benefits-section">
    <MemberBenefits :title-lines="titleLines" :items="memberBenefits" />
  </div>
</template>

<style scoped>
.member-benefits-section {
  --page-bg: #f0f0f1;
  --text-main: #002f55;
  --accent-blue: #3bb0e3;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
