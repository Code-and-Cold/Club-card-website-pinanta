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
      <h2 id="events-title" class="events__title title">{{ title }}</h2>
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
            350: { slidesPerView: 1.2, 
              spaceBetween: 0,
              centeredSlides: false, 
            },
            640: { slidesPerView: 2, spaceBetween: 0 },
            950: { slidesPerView: 3, spaceBetween: 0 },
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
              <h3 class="event-card__name subtitle">{{ item.title }}</h3>
              <p class="event-card__text text">{{ item.text }}</p>
              <p class="event-card__data text">{{ item.data }}</p>
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
            <div v-if="selectedItem" class="modal-body">
                <div class="modal-image-wrapper">
                <img 
                  :src="selectedItem.photo" 
                  class="modal-image" 
                  :alt="selectedItem.title"
                >
              </div>
              <div class="modal-alltext-content">
                <h2 class="modal-title">{{ selectedItem.title }}</h2>
                <div class="modal-text-content">
                    <p class="modal-text">{{ selectedItem.text }}</p>
                    <p v-if="selectedItem.fullText" class="modal-full-text">
                    {{ selectedItem.fullText }}
                    </p>
                    <p class="modal-date">{{ selectedItem.data }}</p>
                </div>
            </div>
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
  color: #002f55;
  margin-left: 5px;
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
  border-radius: 16px;
  transition:
    transform 180ms ease,
    box-shadow 180ms ease;
  height: 100%;
  cursor: pointer;
  width: 430px;
}

.event-card:hover {
  background-color: #e8e8e8;
}

.event-card__content {
  margin-top: auto;
}

.event-card__photo {
  width: 100%;
  height: 209px;
  object-fit: cover;
  margin: auto;
  clip-path: inset(2% 2% 2% 2% round 20px);
  margin-bottom: 15px;
}
.event-card__name,
.event-card__text,
.event-card__data {
  color: #002f55;
  margin-left: 5px;
  margin-bottom: 5px;  
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
  left: -50px;
  right: auto;
}

.event__swiper-button-next {
  right: -50px;
  left: auto;
}

.event__swiper-navigation-icon-prev,
.event__swiper-navigation-icon-next {
  width: 11px;
  height: 20px;
  margin: auto;
}

.swiper-button-disabled {
    opacity: 0;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 47, 85, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  padding: 20px;
}

.modal-content {
  background: white;
  border-radius: 24px;
  max-width: 1220px;
  width: 100%;
  height: 70%;
  display: flex;
  flex-direction: column;
  position: relative;
  padding: 40px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  overflow: hidden;
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
  max-height: 90vh;
  overflow: hidden;
}

.modal-alltext-content {
  display: flex;
  flex-direction: column;
  max-height: 90vh;
  overflow: hidden;
}

.modal-image-wrapper {
  flex-shrink: 0;
}

.modal-image {
  width: 100%;
  height: 208px;
  object-fit: cover;
  border-radius: 16px;
}

.modal-text-content {
  flex: 1;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  scrollbar-width: thin;
  scrollbar-color: #ccc #f5f5f5;
  margin-top: 20px;
}

.modal-text-content::-webkit-scrollbar-track {
    color: orange;
}

.modal-title {
  color: #002f55;
  margin: 8px 0 0 0;
}

.modal-date {
  color: #002f55;
  margin-top: 20px;
}

.modal-text {
  color: #002f55;
  margin: 0;
}

.modal-full-text {
  color: #002f55;
  margin: 8px 0 0 0;
}

@media (min-width: 901px) {
  .modal-body {
    flex-direction: row;
    height: 680px;
  }
  .modal-image {
    width: 485px;
    height: 326px;
    margin-right: 30px;
  }
  .modal-title {
    font-size: 60px;
  }
}

@media (max-width: 900px) {
  .modal-content {
    padding: 24px;
    margin: 10px;
  }
}

@media (max-width: 600px) {
  .events__container {
    width: 100%;
    margin-left: 10px;
  }
  .modal-content {
    padding: 20px;
  }

  .modal-close {
    top: 12px;
    right: 12px;
    width: 36px;
    height: 36px;
  }
  .event__swiper-button-prev {
    left: 0;
    top: 40%;
  }

  .event__swiper-button-next {
    right: 10px;
    top: 40%;
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

<style>
::-webkit-scrollbar-track {
    color: orange;
}
</style>