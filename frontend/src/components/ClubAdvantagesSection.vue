<script setup>
import { computed } from 'vue'
import ClubAdvantages from './ClubAdvantages.vue'

const props = defineProps({
  cards: {
    type: Array,
    default: () => [],
  },
  title: {
    type: String,
    default: 'Что тебя ждет в клубе?',
  },
})

const advantages = computed(() =>
  props.cards.map((card) => ({
    id: card.id,
    icon: card.icon_url,
    title: card.title,
    text: card.description,
  })),
)
</script>

<template>
  <div class="club-advantages-section" id="advantages-section">
    <ClubAdvantages :items="advantages" :title="title" />
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
