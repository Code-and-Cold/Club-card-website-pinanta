<script setup>
defineProps({
  titleLines: {
    type: Array,
    default: () => ['Твой профит', 'от участия в', 'клубе'],
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="member-benefits section" aria-labelledby="member-benefits-title">
    <div class="member-benefits__container">
      <h2 id="member-benefits-title" class="member-benefits__title title">
        <span v-for="line in titleLines" :key="line">{{ line }}</span>
      </h2>

      <div class="member-benefits__list">
        <article v-for="(item, index) in items" :key="item.id ?? index" class="member-benefit">
          <img class="member-benefit__icon" :src="item.icon" alt="" aria-hidden="true" />

          <div class="member-benefit__content">
            <h3 class="member-benefit__title subtitle">{{ item.title }}</h3>
            <p class="member-benefit__text text">{{ item.text }}</p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.member-benefits__container {
  display: grid;
  grid-template-columns: minmax(300px, 0.9fr) minmax(480px, 1.1fr);
  align-items: center;
  gap: clamp(48px, 8vw, 120px);
  width: 100%;
}

.member-benefits__title.title {
  color: var(--text-main);
}

.member-benefits__title span {
  display: block;
}

.member-benefits__list {
  display: grid;
  gap: clamp(22px, 2.6vw, 34px);
}

.member-benefit {
  display: grid;
  grid-template-columns: 65px minmax(0, 1fr);
  align-items: start;
  gap: clamp(16px, 1.8vw, 24px);
}

.member-benefit__icon {
  display: block;
  width: 65px;
  height: 65px;
}

.member-benefit__content {
  min-width: 0;
}

.member-benefit__title.subtitle,
.member-benefit__text.text {
  color: var(--text-main);
}

.member-benefit__title.subtitle {
  margin-bottom: 6px;
}

.member-benefit__text.text {
  max-width: 620px;
}

@media (max-width: 900px) {
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
  .member-benefit {
    grid-template-columns: 55px minmax(0, 1fr);
    gap: 14px;
  }

  .member-benefit__icon {
    width: 55px;
    height: 55px;
  }
}
</style>
