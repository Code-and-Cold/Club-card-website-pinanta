<script setup>
import ClubIcon from './ClubIcon.vue'

defineProps({
  title: {
    type: String,
    default: 'Почему у нас классно?',
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="advantages" aria-labelledby="advantages-title">
    <div class="advantages__container">
      <h2 id="advantages-title" class="advantages__title">{{ title }}</h2>

      <div class="advantages__grid">
        <article v-for="(item, index) in items" :key="item.id ?? index" class="advantage-card">
          <ClubIcon variant="badge" />

          <div class="advantage-card__content">
            <h3 class="advantage-card__title">{{ item.title }}</h3>
            <p class="advantage-card__accent">{{ item.accent }}</p>
            <p class="advantage-card__text">{{ item.text }}</p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.advantages {
  padding-block: clamp(48px, 7vw, 108px);
}

.advantages__container {
  width: min(100% - 32px, 1180px);
  margin-inline: auto;
}

.advantages__title {
  margin-bottom: clamp(28px, 4vw, 54px);
  font-family: 'JetBrains Mono', monospace;
  font-size: clamp(28px, 4.1vw, 52px);
  font-weight: 400;
  line-height: 1.08;
  letter-spacing: 0.015em;
}

.advantages__grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(12px, 1.8vw, 24px);
  padding-inline: clamp(0px, 8vw, 95px);
}

.advantage-card {
  display: flex;
  flex-direction: column;
  min-height: clamp(210px, 23vw, 285px);
  padding: clamp(20px, 2.4vw, 30px);
  border-radius: 16px;

  background-color: var(--card-bg);
  color: var(--text-main);

  transition:
    background-color 180ms ease,
    color 180ms ease,
    transform 180ms ease,
    box-shadow 180ms ease;
}

.advantage-card:hover {
  background-color: #002f55;
  color: #ffffff;

  transform: translateY(-4px);
  box-shadow: 0 16px 35px rgb(0 47 85 / 18%);
}

.advantage-card__content {
  /* Фиксируем начало текста относительно верха карточки:
     длина описания больше не двигает заголовок вверх/вниз. */
  margin-top: clamp(18px, 2vw, 24px);
}

.advantage-card__title,
.advantage-card__accent,
.advantage-card__text {
  font-size: clamp(13px, 1.15vw, 17px);
  line-height: 1.15;
}

.advantage-card__title {
  font-family: 'JetBrains Mono', monospace;
  font-weight: 400;
}

.advantage-card__accent {
  color: var(--accent-orange);
  font-family: 'JetBrains Mono', monospace;
  font-weight: 700;
  text-transform: uppercase;
}

.advantage-card__text {
  margin-top: 7px;
  font-family: 'Inter', sans-serif;
  font-size: clamp(10px, 0.82vw, 13px);
  font-weight: 400;
  line-height: 1.25;
}

@media (max-width: 1000px) {
  .advantages__grid {
    gap: 12px;
    padding-inline: 0;
  }
}

@media (max-width: 760px) {
  .advantages__container {
    width: min(100% - 24px, 1180px);
  }

  .advantages__grid {
    grid-template-columns: 1fr;
    width: min(100%, 310px);
    margin-inline: auto;
  }

  .advantage-card {
    min-height: 190px;
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
