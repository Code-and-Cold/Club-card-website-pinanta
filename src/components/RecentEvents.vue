<script setup>
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Navigation, Pagination } from 'swiper/modules'
import 'swiper/css'
import 'swiper/css'
import 'swiper/css/pagination'

defineProps({
  title: {
    type: String,
    default: 'Недавние события',
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="events" aria-labelledby="events-title">
  <div class="events__container">
    <h2 id="events-title" class="events__title">{{ title }}</h2>
  <Swiper
    :modules="[Navigation, Pagination]"
    :slides-per-view="1"
    :space-between="10"
    pagination
  >
    <SwiperSlide v-for="(item, index) in items"
          :key="item.id ?? index"
          class="event-card">
          <div class="event-card__content">
            <img :src="item.photo" class="event-card__photo">
            <h3 class="event-card__name">{{ item.title }}</h3>
            <p class="event-card__text">{{ item.text }}</p>
            <p class="event-card__data">{{ item.data }}</p>
          </div>
    </SwiperSlide>
  </Swiper>
  </div>
  </section>
</template>

<style scoped>
.events {
  padding-block: clamp(48px, 7vw, 108px);
  flex: 0 0 auto
}

.events__container {
  width: min(100% - 32px, 1180px);
  margin-inline: auto;
}

.events__title {
  margin-bottom: clamp(28px, 4vw, 54px);
  font-size: clamp(28px, 4.1vw, 52px);
  font-weight: bold;
  line-height: 1.08;
  letter-spacing: 0.015em;
}

.events__grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(12px, 1.8vw, 24px);
  padding-inline: clamp(0px, 8vw, 95px);
}

.event-card {
  display: flex;
  flex-direction: column;
  min-height: clamp(210px, 23vw, 285px);
  padding: clamp(20px, 2.4vw, 30px);
  border-radius: 16px;
  transition: transform 180ms ease, box-shadow 180ms ease;
}

.event-card__content {
  margin-top: auto;
}

.event-card__photo {
  margin: auto;
  height: 480px;
  clip-path: inset(2% 2% 2% 2% round 20px);
}
.event-card__name,
.event-card__text,
.event-card__data {
  font-size: clamp(13px, 1.15vw, 17px);
  line-height: 1.15;
}

.event-card__name {
  font-weight: bold;
}

.event-card__text,
.event-card__data {
  margin-top: 7px;
  font-size: clamp(10px, 0.82vw, 13px);
  line-height: 1.25;
}

@media (max-width: 900px) {
  .events__grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
    padding-inline: 0;
  }
}

@media (max-width: 600px) {
  .events__container {
    width: min(100% - 24px, 1180px);
  }

  .events__grid {
    grid-template-columns: 1fr;
  }

  .event-card {
    min-height: 190px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .event-card {
    transition: none;
  }
}
</style>