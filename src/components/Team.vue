<script setup>
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Navigation, Pagination } from 'swiper/modules'
import 'swiper/css'
import 'swiper/css'
import 'swiper/css/pagination'

defineProps({
  title: {
    type: String,
    default: 'Наша команда',
  },
  items: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="team" aria-labelledby="team-title">
  <div class="team__container">
    <h2 id="team-title" class="team__title">{{ title }}</h2>
  <Swiper
    :modules="[Navigation, Pagination]"
    :slides-per-view="1"
    :space-between="10"
    pagination
  >
    <SwiperSlide v-for="(item, index) in items"
          :key="item.id ?? index"
          class="member-card">
          <div class="member-card__content">
            <img :src="item.photo" class="member-card__photo">
            <h3 class="member-card__name">{{ item.name }}</h3>
            <p class="member-card__position">{{ item.position }}</p>
          </div>
    </SwiperSlide>
  </Swiper>
  </div>
  </section>
</template>

<style scoped>
.team {
  padding-block: clamp(48px, 7vw, 108px);
  flex: 0 0 auto
}

.team__container {
  width: min(100% - 32px, 1180px);
  margin-inline: auto;
}

.team__title {
  font-size: clamp(28px, 4.1vw, 52px);
  font-weight: bold;
  line-height: 1.08;
  letter-spacing: 0.015em;
}

.team__grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(12px, 1.8vw, 24px);
  padding-inline: clamp(0px, 8vw, 95px);
}

.member-card {
  display: flex;
  flex-direction: column;
  min-height: clamp(210px, 23vw, 285px);
  padding: clamp(20px, 2.4vw, 30px);
  border-radius: 16px;
  transition: transform 180ms ease, box-shadow 180ms ease;
}

.member-card__content {
  margin: auto;
}

.member-card__photo {
  object-fit: fill;
  margin: auto;
  clip-path: inset(2% 2% 2% 2% round 20px);
}

.member-card__name,
.member-card__position {
  font-size: clamp(13px, 1.15vw, 17px);
  line-height: 1.15;
}

.member-card__name {
  font-weight: bold;
}

.member-card__position {
  margin-top: 7px;
  font-size: clamp(10px, 0.82vw, 13px);
  line-height: 1.25;
}

@media (max-width: 900px) {
  .team__grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
    padding-inline: 0;
  }
}

@media (max-width: 600px) {
  .team__container {
    width: min(100% - 24px, 1180px);
  }

  .team__grid {
    grid-template-columns: 1fr;
  }

  .member-card {
    min-height: 190px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .member-card {
    transition: none;
  }
}
</style>
