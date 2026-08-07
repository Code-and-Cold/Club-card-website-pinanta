<template>
  <section class="hero-wrapper">
    <div class="hero">
      <header class="hero__header">
        <div class="hero__logo">Лого сафу и клуба</div>
        <nav v-if="isDesktop" class="hero__nav-desktop">
          <ul class="hero__nav-list">
            <li v-for="item in menuItems" :key="item.id">
              <a :href="item.href" class="hero__nav-link subtitle">
                {{ item.label }}
              </a>
            </li>
          </ul>
        </nav>

        <button
          v-if="isDesktop"
          class="hero__register-btn hero_register-btn--desktop"
          @click="handleRegister"
        >
          Вступить в клуб
        </button>

        <div class="hero__mobile-buttons" v-if="!isDesktop">
          <button v-if="!isMobileMenuOpen" class="hero__vk-link-btn" @click="goToVKGroup">
            <img src="../assets/vector/ri_vk-fill.svg" alt="Группа в ВК" />
          </button>

          <button
            class="hero__burger"
            @click="toggleMobileMenu"
            :aria-label="isMobileMenuOpen ? 'Закрыть меню' : 'Открыть меню'"
          >
            <span
              class="hero__burger-line"
              :class="{ 'hero__burger-line--active': isMobileMenuOpen }"
            ></span>
            <span
              class="hero__burger-line"
              :class="{ 'hero__burger-line--active': isMobileMenuOpen }"
            ></span>
            <span
              class="hero__burger-line"
              :class="{ 'hero__burger-line--active': isMobileMenuOpen }"
            ></span>
          </button>
        </div>
      </header>

      <div v-if="isMobileMenuOpen" class="hero__mobile-overlay" @click="closeMobileMenu">
        <nav class="hero__mobile-nav" @click.stop>
          <div></div>
          <ul class="hero__mobile-list">
            <li v-for="item in menuItems" :key="item.id">
              <a :href="item.href" class="hero__mobile-link" @click="closeMobileMenu">
                {{ item.label }}
              </a>
            </li>
          </ul>
          <button v-if="!isDesktop" class="hero__vk-content-btn" @click="goToVKGroup">
            <p>Подписывайтесь на группу Вконтакте</p>
            <img src="../assets/vector/Vector.svg" />
          </button>
        </nav>
      </div>

      <div class="hero__content">
        <div class="hero__content-row">
          <h1 class="hero__title title">Пишем код. Согреваем атмосферой. Создаём твоё портфолио.</h1>
          <button v-if="isDesktop" class="hero__vk-content-btn" @click="goToVKGroup">
            <p>Подписывайтесь на группу Вконтакте</p>
            <img class="vk-logo-blue" src="../assets/vector/Vector.svg" />
          </button>
        </div>
        <button
          v-if="!isDesktop"
          class="hero__register-btn hero__register-btn--mobile"
          @click="handleRegister"
        >
          Вступить в клуб
        </button>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

const menuItems = [
  { id: 1, label: 'Преимущества', href: '#advantages-section' },
  { id: 2, label: 'Команда', href: '#team-section' },
  { id: 3, label: 'Плюшки', href: '#benefits-section' },
]

const isMobileMenuOpen = ref(false)
const windowWidth = ref(window.innerWidth)

const isDesktop = computed(() => windowWidth.value > 950)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value

  if (isMobileMenuOpen.value) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
}

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false
  document.body.style.overflow = ''
}

const handleRegister = () => {
  // Пока ничего
}

const updateWidth = () => {
  windowWidth.value = window.innerWidth

  if (windowWidth.value > 950 && isMobileMenuOpen.value) {
    closeMobileMenu()
  }
}

onMounted(() => {
  window.addEventListener('resize', updateWidth)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateWidth)
  document.body.style.overflow = ''
})
</script>

<style lang="css" scoped>
/* TODO: Вынести в HeroSection как будут готовы тесты */
h1 {
  text-align: center;
  color: white;
}

.hero-wrapper {
  max-height: 100vh;
  background-color: rgba(0, 0, 0, 0);
  display: flex;
  align-items: center;
  justify-content: center;

  padding: clamp(15px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 30px);
}

.hero {
  width: 100%;
  min-height: calc(100vh - 60px);
  background: linear-gradient(180deg, rgba(0, 0, 0, 0) 0%, rgba(0, 47, 85, 0.7) 100%);
  display: flex;
  flex-direction: column;
  position: relative;
  
  margin: 0;
  padding-left: clamp(20px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 40px);
  padding-right: clamp(20px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 40px);
  padding-bottom: clamp(20px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 40px);
  padding-top: 0;

  border-radius: 30px;
  box-shadow: 0px 0px 0px 50px #002f55;
  isolation: isolate;
  font-weight: 300;
}

.hero__header {
  height: 100px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 30px;
  padding: 0;
  position: relative;
  z-index: 101;
}

.hero__logo {
  /* Пока настройки текста*/
  color: white;
  font-size: 26px;
}

.hero__nav-desktop {
  display: flex;
  align-items: center;
  gap: 50px;
}

.hero__nav-list {
  display: flex;
  list-style: none;
  gap: 50px;
  margin: 0;
  padding: 0;
}

.hero__nav-link {
  color: white;
  text-decoration: none;
  position: relative;
  padding: auto;
  transition: all 0.3s ease;  
}

.hero__register-btn {
  background: #e3953b;
  color: white;
  border: none;
  padding: 15px 25px;
  border-radius: 15px;
  box-sizing: content-box;
  font-size: 24px;
  line-height: 105%;

  transition: background-color 0.4s ease;
}

.hero__register-btn:hover {
  background: #3BB0E3;
}

.hero__register-btn--desktop {
}

.hero__register-btn--mobile {
  width: fit-content;
}

.hero__mobile-buttons {
  background: none;
  display: flex;
  flex-direction: row;
  gap: 10px;
}

.hero__vk-link-btn {
  background: none;
  border: none;
  border-radius: 10px;
}

.hero__vk-link-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.hero__burger {
  background: none;
  border: none;
  cursor: pointer;
  padding: 10px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  z-index: 20;
  position: relative;
  border-radius: 10px;
  transition: background 0.3s ease;
}

.hero__burger:hover {
  background: rgba(255, 255, 255, 0.1);
}

.hero__burger-line {
  width: 28px;
  height: 3px;
  background: white;
  border-radius: 2px;
  transition: all 0.3s ease;
  transform-origin: center;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.hero__burger-line--active:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.hero__burger-line--active:nth-child(2) {
  opacity: 0;
  transform: scaleX(0);
}

.hero__burger-line--active:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

.hero__mobile-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  z-index: 100;
  display: flex;
  justify-content: flex-end;
}

.hero__mobile-nav {
  background: #002f55;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  width: 100%;
  max-height: 60%;
  padding: 100px 0px 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 30px;
  box-shadow: -10px 0 40px rgba(0, 0, 0, 0.3);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  z-index: 100;
}

.hero__mobile-list {
  list-style: none;
  padding: 0;
  margin: 0;
  color: white;
  text-align: center;
  font-size: 22px;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  gap: 10px;
}

.hero__mobile-link {
  color: white;
  text-decoration: none;
  padding: 15px 20px;
  transition: all 0.3s ease;
  display: block;
  position: relative;
  overflow: hidden;

}

.hero__nav-link:hover, .hero__mobile-link:hover {
  color: #3BB0E3;
}

.hero__content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: end;
  align-items: center;
  text-align: left;
  padding-top: 30px;
  position: relative;
  z-index: 1;
  gap: 18px;
}

.hero__content-row {
  display: flex;
  flex-direction: row;
  box-sizing: content-box;
  justify-content: space-between;
  align-items: end;
  gap: 30px;
}

.hero__title {
  margin: 0;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  text-align: left;
}

.hero__vk-content-btn {
  display: flex;
  flex-direction: row;

  box-sizing: border-box;
  align-items: center;
  text-align: left;
  background: rgba(0, 47, 85, 0.5);
  border: 1px solid #3bb0e3;
  border-radius: 20px;
  color: white;
  font-size: 22px;
  font-weight: 300;
  gap: 15px;
  padding: 15px 25px;
  max-width: 400px;
}

@media (min-width: 950px) {
  .vk-logo-blue {
    height: 60px;
    width: 60px;
  }
}

@media (max-width: 950px) {

  .hero__header {
    height: 75px;
  }

  .hero__logo-text {
    font-size: 24px;
  }
}

@media (max-width: 480px) {
  .hero__logo-text {
    font-size: 20px;
  }
  .hero__mobile-nav {
    padding: 80px 20px 20px;
  }

}
</style>
