<script setup>
import { computed } from 'vue'
import MemberBenefits from './MemberBenefits.vue'

const props = defineProps({
  items: {
    type: Array,
    default: () => [],
  },
  title: {
    type: String,
    default: 'Твой профит\nот участия в\nклубе',
  },
})

const memberBenefits = computed(() =>
  props.items.map((item) => ({
    id: item.id,
    icon: item.icon_url,
    title: item.title,
    text: item.description,
  })),
)

const titleLines = computed(() => props.title.split('\n'))
</script>

<template>
  <div class="member-benefits-section" id="benefits-section">
    <MemberBenefits :items="memberBenefits" :title-lines="titleLines" />
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
