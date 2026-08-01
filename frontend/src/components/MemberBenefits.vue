<script setup>
import ClubIcon from './ClubIcon.vue'

defineProps({
  titleLines: {
    type: Array,
    default: () => ['Что ты получишь,', 'вступив к нам', 'и круто с нами', 'поработав'],
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="member-benefits" aria-labelledby="member-benefits-title">
    <div class="member-benefits__container">
      <h2 id="member-benefits-title" class="member-benefits__title">
        <span v-for="line in titleLines" :key="line">{{ line }}</span>
      </h2>

      <div class="member-benefits__list">
        <article
          v-for="(item, index) in items"
          :key="item.id ?? index"
          class="member-benefit"
        >
          <ClubIcon variant="cat" />

          <div>
            <h3 class="member-benefit__title">{{ item.title }}</h3>
            <p class="member-benefit__text">{{ item.text }}</p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.member-benefits {
  padding-block: clamp(56px, 8vw, 118px);
}

.member-benefits__container {
  display: grid;
  grid-template-columns: minmax(280px, 0.9fr) minmax(420px, 1.1fr);
  align-items: center;
  gap: clamp(48px, 8vw, 120px);
  width: min(100% - 40px, 1180px);
  margin-inline: auto;
}

.member-benefits__title {
  font-size: clamp(31px, 4.6vw, 58px);
  font-weight: 400;
  line-height: 1.03;
  letter-spacing: 0.015em;
}

.member-benefits__title span {
  display: block;
}

.member-benefits__list {
  display: grid;
  gap: clamp(22px, 3vw, 38px);
}

.member-benefit {
  display: grid;
  grid-template-columns: 48px minmax(0, 1fr);
  align-items: start;
  gap: 18px;
}

.member-benefit__title {
  margin-bottom: 6px;
  font-size: clamp(16px, 1.55vw, 21px);
  font-weight: 500;
  line-height: 1.15;
}

.member-benefit__text {
  max-width: 620px;
  font-size: clamp(11px, 0.95vw, 14px);
  line-height: 1.4;
}

@media (max-width: 850px) {
  .member-benefits__container {
    grid-template-columns: 1fr;
    align-items: start;
    gap: 42px;
  }

  .member-benefits__title {
    max-width: 700px;
  }
}

@media (max-width: 520px) {
  .member-benefits__container {
    width: min(100% - 24px, 1180px);
  }

  .member-benefit {
    grid-template-columns: 38px minmax(0, 1fr);
    gap: 13px;
  }
}
</style>
