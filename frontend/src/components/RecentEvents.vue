<script setup>
import { ref } from 'vue'
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

const selectedItem = ref(null)
const isModalOpen = ref(false)

const openModal = (item) => {
  selectedItem.value = item
  isModalOpen.value = true
  document.body.style.overflow = 'hidden'
}

const closeModal = () => {
  isModalOpen.value = false
  document.body.style.overflow = ''
  setTimeout(() => {
    selectedItem.value = null
  }, 300)
}

const handleOverlayClick = (e) => {
  if (e.target === e.currentTarget) {
    closeModal()
  }
}

const handleKeydown = (e) => {
  if (e.key === 'Escape' && isModalOpen.value) {
    closeModal()
  }
}
</script>

<template>
  <section class="events" aria-labelledby="events-title">
    <div class="events__container">
      <h2 id="events-title" class="events__title">{{ title }}</h2>
      <div class="events__slider-wraper">
        <div ref="event__swiper-button-prev" class="event__swiper-button-prev">
          <svg
            ref="event__swiper-navigation-icon-prev"
            class="event__swiper-navigation-icon-prev"
            viewBox="0 0 11 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M10.4341 20.0762C10.7053 19.805 10.7053 19.3654 10.4341 19.0942L1.61953 10.2796L10.4341 1.46497C10.7053 1.19379 10.7053 0.754138 10.4341 0.482966C10.1629 0.211794 9.72329 0.211794 9.45212 0.482966L0.38296 9.55214C-0.0188039 9.9539 -0.0188039 10.6053 0.38296 11.007L9.45212 20.0762C9.72329 20.3474 10.1629 20.3474 10.4341 20.0762Z"
              fill="currentColor"
            />
          </svg>
        </div>
        <Swiper
          :modules="[Navigation, Pagination]"
          :breakpoints="{
            375: { slidesPerView: 1, spaceBetween: 10 },
            640: { slidesPerView: 2, spaceBetween: 10 },
            1024: { slidesPerView: 3, spaceBetween: 20 },
          }"
          :navigation="{
            prevEl: '.event__swiper-button-prev',
            nextEl: '.event__swiper-button-next',
          }"
          class="event__swiper"
        >
          <SwiperSlide
            v-for="(item, index) in items"
            :key="item.id ?? index"
            class="event-card"
            @click="openModal(item)"
          >
            <div class="event-card__content">
              <img :src="item.photo" class="event-card__photo" :alt="item.title" />
              <h3 class="event-card__name">{{ item.title }}</h3>
              <p class="event-card__text">{{ item.text }}</p>
              <p class="event-card__data">{{ item.data }}</p>
            </div>
          </SwiperSlide>
        </Swiper>
        <div ref="event__swiper-button-next" class="event__swiper-button-next">
          <svg
            ref="event__swiper-navigation-icon-next"
            class="event__swiper-navigation-icon-next"
            viewBox="0 0 11 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M0.38296 20.0762C0.111788 19.805 0.111788 19.3654 0.38296 19.0942L9.19758 10.2796L0.38296 1.46497C0.111788 1.19379 0.111788 0.754138 0.38296 0.482966C0.654131 0.211794 1.09379 0.211794 1.36496 0.482966L10.4341 9.55214C10.8359 9.9539 10.8359 10.6053 10.4341 11.007L1.36496 20.0762C1.09379 20.3474 0.654131 20.3474 0.38296 20.0762Z"
              fill="currentColor"
            />
          </svg>
        </div>
      </div>
    </div>

    <Teleport to="body">
      <Transition name="modal">
        <div
          v-if="isModalOpen"
          class="modal-overlay"
          @click="handleOverlayClick"
          @keydown="handleKeydown"
          tabindex="0"
        >
          <div class="modal-content" role="dialog" aria-modal="true">
            <button class="modal-close" @click="closeModal" aria-label="Закрыть новость">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
                <path
                  d="M18 6L6 18M6 6L18 18"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
            </button>

            <div v-if="selectedItem" class="modal-body">
              <img :src="selectedItem.photo" class="modal-image" :alt="selectedItem.title" />
              <h2 class="modal-title">{{ selectedItem.title }}</h2>
              <p v-if="selectedItem.fullText" class="modal-full-text">
                {{ selectedItem.fullText }}
              </p>
              <p v-else class="modal-text">{{ selectedItem.text }}</p>
              <p class="modal-date">{{ selectedItem.data }}</p>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </section>
</template>

<style scoped>
.events {
  padding-block: clamp(48px, 7vw, 108px);
  flex: 0 0 auto;
  background-color: #f0f0f1;
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
  color: #002f55;
}

.events__slider-wraper {
  display: flex;
  align-items: center;
  position: relative;
  width: 100%;
  height: 486px;
}

.event__swiper {
  width: 100%;
  overflow: hidden;
}

.event-card {
  display: flex;
  flex-direction: column;
  min-height: clamp(210px, 23vw, 285px);
  padding: clamp(20px, 2.4vw, 30px);
  border-radius: 16px;
  transition:
    transform 180ms ease,
    box-shadow 180ms ease;
  height: 100%;
  cursor: pointer;
}

.event-card:hover {
  transform: translateY(-4px);
  background-color: #e8e8e8;
}

.event-card__content {
  margin-top: auto;
}

.event-card__photo {
  width: 312px;
  height: 209px;
  margin: auto;
  height: 480px;
  clip-path: inset(2% 2% 2% 2% round 20px);
}
.event-card__name,
.event-card__text,
.event-card__data {
  font-size: clamp(13px, 1.15vw, 17px);
  line-height: 1.15;
  color: #002f55;
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

.event__swiper-button-prev,
.event__swiper-button-next {
  position: absolute;
  z-index: 10;
  top: 50%;
  width: 50px;
  height: 50px;
  margin-top: 0;
  flex-shrink: 0;
  color: #002f55;
  background: #e7e7e7;
  border-radius: 50%;
  transition: all 200ms ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.event__swiper-button-prev:hover,
.event__swiper-button-next:hover {
  color: #3bb0e3;
}
.event__swiper-button-prev {
  left: 10px;
  right: auto;
}

.event__swiper-button-next {
  right: 10px;
  left: auto;
}

.event__swiper-navigation-icon-prev,
.event__swiper-navigation-icon-next {
  width: 11px;
  height: 20px;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  padding: 20px;
}

.modal-content {
  background: white;
  border-radius: 24px;
  max-width: 700px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  position: relative;
  padding: 40px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.modal-close {
  position: absolute;
  top: 16px;
  right: 16px;
  width: 40px;
  height: 40px;
  border: none;
  background: rgba(0, 0, 0, 0.05);
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 200ms ease;
  color: #333;
  z-index: 1;
}

.modal-close:hover {
  background: rgba(0, 0, 0, 0.1);
  transform: rotate(90deg);
}

.modal-body {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.modal-image {
  width: 100%;
  max-height: 400px;
  object-fit: cover;
  border-radius: 16px;
}

.modal-title {
  font-size: clamp(24px, 3vw, 32px);
  font-weight: bold;
  color: #002f55;
  margin: 8px 0 0 0;
}

.modal-date {
  font-size: 14px;
  color: #002f55;
  margin: 0;
}

.modal-text {
  font-size: 16px;
  line-height: 1.6;
  color: #002f55;
  margin: 0;
}

.modal-full-text {
  font-size: 16px;
  line-height: 1.8;
  color: #002f55;
  margin: 8px 0 0 0;
  padding-top: 16px;
  border-top: 1px solid #eee;
}

.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

.modal-enter-to,
.modal-leave-from {
  opacity: 1;
  transform: scale(1);
}

.modal-content::-webkit-scrollbar {
  width: 6px;
}

.modal-content::-webkit-scrollbar-track {
  background: transparent;
}

.modal-content::-webkit-scrollbar-thumb {
  background: #ccc;
  border-radius: 3px;
}

.modal-content::-webkit-scrollbar-thumb:hover {
  background: #aaa;
}

@media (max-width: 900px) {
  .modal-content {
    padding: 24px;
    margin: 10px;
  }
}

@media (max-width: 600px) {
  .modal-content {
    padding: 20px;
  }

  .modal-close {
    top: 12px;
    right: 12px;
    width: 36px;
    height: 36px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .event-card,
  .modal-close,
  .modal-enter-active,
  .modal-leave-active {
    transition: none;
  }
}
</style>
