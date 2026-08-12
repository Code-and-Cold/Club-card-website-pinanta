<script setup>
defineProps({
  title: {
    type: String,
    default: 'Что тебя ждет в клубе?',
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="advantages section" aria-labelledby="advantages-title">
    <div class="advantages__container">
      <h2 id="advantages-title" class="advantages__title title">{{ title }}</h2>

      <div class="advantages__grid">
        <article v-for="(item, index) in items" :key="item.id ?? index" class="advantage-card">
          <img class="advantage-card__icon" :src="item.icon" alt="" aria-hidden="true" />

          <div class="advantage-card__content">
            <h3 class="advantage-card__title subtitle">{{ item.title }}</h3>
            <p class="advantage-card__text text">{{ item.text }}</p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.advantages__container {
  width: 100%;
}

.advantages__title.title {
  margin-bottom: clamp(36px, 4vw, 58px);
  color: var(--text-main);
}

.advantages__grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(12px, 1.8vw, 24px);
  width: min(100%, 1100px);
  margin-inline: auto;
}

.advantage-card {
  display: flex;
  flex-direction: column;
  min-width: 0;
  min-height: clamp(300px, 23vw, 340px);
  padding: clamp(24px, 2.2vw, 32px);
  border-radius: 16px;
  background-color: var(--card-bg);
  transition:
    background-color 180ms ease,
    transform 180ms ease,
    box-shadow 180ms ease;
}

.advantage-card:hover {
  background-color: #002f55;
  transform: translateY(-4px);
  box-shadow: 0 16px 35px rgb(0 47 85 / 18%);
}

.advantage-card__icon {
  display: block;
  width: 60px;
  height: 60px;
  flex: 0 0 60px;
}

.advantage-card__content {
  display: grid;
  gap: 10px;
  margin-top: clamp(18px, 2vw, 26px);
}

.advantage-card__title.subtitle,
.advantage-card__text.text {
  color: var(--text-main);
}

.advantage-card:hover .advantage-card__title.subtitle,
.advantage-card:hover .advantage-card__text.text {
  color: #ffffff;
}

@media (max-width: 1000px) {
  .advantages__grid {
    gap: 12px;
  }
}

@media (max-width: 760px) {
  .advantages__grid {
    grid-template-columns: 1fr;
    width: min(100%, 310px);
  }

  .advantage-card {
    min-height: 310px;
  }
}

@media (min-width: 601px) and (max-width: 760px) {
  .advantages__grid {
    width: min(100%, 360px);
  }
}

@media (prefers-reduced-motion: reduce) {
  .advantage-card {
    transition: none;
  }
}
</style>
