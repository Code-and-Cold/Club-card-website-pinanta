<script setup>
import MemberBenefits from './MemberBenefits.vue'
import documentIcon from '../assets/vector/benefit_document.svg'
import lightningIcon from '../assets/vector/benefit_lightning.svg'
import badgeIcon from '../assets/vector/benefit_badge.svg'
import partyIcon from '../assets/vector/benefit_party.svg'

const memberBenefits = [
  {
    id: 1,
    icon: documentIcon,
    title: 'Строчка в резюме',
    text: 'Получишь опыт участия в проектах полного цикла — от проектирования до релиза,  — и это выделит тебя среди других джунов на рынке.',
  },
  {
    id: 2,
    icon: lightningIcon,
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
    icon: partyIcon,
    title: 'Сохраненные нервы',
    text: 'Тебе не придется часами гуглить непонятную ошибку в одиночку. В клубе есть у кого спросить совет по архитектуре, попросить код-ревью или просто узнать, как сдать сложный предмет.',
  },
]
</script>

<template>
  <div class="member-benefits-section" id="benefits-section">
    <MemberBenefits :items="memberBenefits" />
  </div>
</template>

<style scoped>
.member-benefits-section {
  --page-bg: #f3f3f3;
  --text-main: #073b63;

  display: flow-root;
  width: 100%;
  color: var(--text-main);
  background-color: var(--page-bg);
  isolation: isolate;
}
</style>
