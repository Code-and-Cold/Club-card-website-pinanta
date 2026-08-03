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
        <div class="team__slider-wraper">
          <div ref="team__swiper-button-prev" class="team__swiper-button-prev">
            <svg ref="team__swiper-navigation-icon-prev" class="team__swiper-navigation-icon-prev" viewBox="0 0 11 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M10.4341 20.0762C10.7053 19.805 10.7053 19.3654 10.4341 19.0942L1.61953 10.2796L10.4341 1.46497C10.7053 1.19379 10.7053 0.754138 10.4341 0.482966C10.1629 0.211794 9.72329 0.211794 9.45212 0.482966L0.38296 9.55214C-0.0188039 9.9539 -0.0188039 10.6053 0.38296 11.007L9.45212 20.0762C9.72329 20.3474 10.1629 20.3474 10.4341 20.0762Z" fill="currentColor"/>
            </svg>
          </div>
          <Swiper
            :modules="[Navigation, Pagination]"
            :breakpoints="{
              375: { slidesPerView: 1, spaceBetween: 10 },
              640: { slidesPerView: 2, spaceBetween: 10 },
              1024: { slidesPerView: 3, spaceBetween: 20 }
            }"
            :navigation="{
              prevEl: '.team__swiper-button-prev',
              nextEl: '.team__swiper-button-next'
            }"
            class="team__swiper"
          >
            <SwiperSlide v-for="(item, index) in items"
              :key="item.id ?? index"
              class="member-card"
            >
              <div class="member-card__content">
                <img :src="item.photo" class="member-card__photo" alt="Здесь могла быть ваша реклама">
                <h3 class="member-card__name">{{ item.name }}</h3>
                <p class="member-card__position">{{ item.position }}</p>
              </div>
            </SwiperSlide>
          </Swiper>
          <div ref="team__swiper-button-next" class="team__swiper-button-next">
            <svg ref="team__swiper-navigation-icon-next" class="team__swiper-navigation-icon-next" viewBox="0 0 11 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M0.38296 20.0762C0.111788 19.805 0.111788 19.3654 0.38296 19.0942L9.19758 10.2796L0.38296 1.46497C0.111788 1.19379 0.111788 0.754138 0.38296 0.482966C0.654131 0.211794 1.09379 0.211794 1.36496 0.482966L10.4341 9.55214C10.8359 9.9539 10.8359 10.6053 10.4341 11.007L1.36496 20.0762C1.09379 20.3474 0.654131 20.3474 0.38296 20.0762Z" fill="currentColor"/>
            </svg>
          </div>
        </div>
    </div>
  </section>
</template>

<style scoped>
.team {
  padding-block: clamp(48px, 7vw, 108px);
  flex: 0 0 auto;
  background-color: #F0F0F1;
}

.team__container {
  width: min(100% - 32px, 1180px);
  margin-inline: auto;
}

.team__title {
  font: 'JetBrains Mono';
  font-size: clamp(28px, 4.1vw, 52px);
  color: #002F55;
  font-weight: 350;
  line-height: 1.05;
  letter-spacing: 0;
  text-align: justify;
}

.team__slider-wraper {
  display: flex;
  align-items: center;
  position: relative;
  width: 100%;
  height: 486px;
}

.team__swiper {
  width: 100%;
  overflow: hidden;
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
  height: 100%;
}

.member-card:hover {
  transform: translateY(-4px);
  background-color: #E8E8E8;
}
.member-card__content {
  margin: auto;
}

.member-card__photo {
  width: 312px;
  height: 412px;
  object-fit: fill;
  margin: auto;
  clip-path: inset(2% 2% 2% 2% round 20px);
}

.member-card__name,
.member-card__position {
  font-size: clamp(13px, 1.15vw, 17px);
  line-height: 1.10;
  color: #002F55;
}

.member-card__name {
  font-weight: 240;
}

.member-card__position {
  margin-top: 7px;
  font-size: clamp(10px, 0.82vw, 13px);
  line-height: 1.25;
}

.team__swiper-button-prev,
.team__swiper-button-next {
  position: absolute;
  z-index: 10;
  top: 50%;
  width: 50px;
  height: 50px;
  margin-top: 0;
  flex-shrink: 0;
  color: #002F55;
  background: #E7E7E7;
  border-radius: 50%;
  transition: all 200ms ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.team__swiper-button-prev:hover,
.team__swiper-button-next:hover {
    color: #3BB0E3;
}

.team__swiper-button-prev {
  left: 10px;
  right: auto;
}

.team__swiper-button-next {
  right: 10px;
  left: auto;
}

.team__swiper-navigation-icon-prev,
.team__swiper-navigation-icon-next {
  width: 11px;
  height: 20px;
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
