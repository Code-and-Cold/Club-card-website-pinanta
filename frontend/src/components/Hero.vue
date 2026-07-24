<template>
    
    <section class="hero-wrapper">
        
        <div class="hero">
            <header class="hero__header">
                <div class="hero__logo">
                    Лого сафу и клуба
                </div>
                <nav v-if="isDesktop" class="hero__nav-desktop">
                    <ul class="hero__nav-list">
                    <li v-for="item in menuItems" :key="item.id">
                            <a :href="item.href" class="hero__nav-link">
                                {{ item.label }}
                            </a>
                        </li>
                    </ul>

                    <button class="hero__register-btn hero_register-btn--desktop" @click="handleRegister">
                        Вступить в клуб
                    </button>
                </nav>
                
                <div class="hero__mobile-buttons"
                    v-if="!isDesktop"
                >
                    <button
                        class="hero__vk-link-btn"
                        @click="goToVKGroup"
                    >
                        <img src="../assets/vector/ri_vk-fill.svg" alt="Группа в ВК">
                    </button>

                    <button
                        class="hero__burger"
                        @click="toggleMobileMenu"
                        :aria-label="isMobileMenuOpen ? 'Закрыть меню' : 'Открыть меню'"
                    >
                        <span class="hero__burger-line"  :class="{ 'hero__burger-line--active': isMobileMenuOpen }"></span>
                        <span class="hero__burger-line"  :class="{ 'hero__burger-line--active': isMobileMenuOpen }"></span>
                        <span class="hero__burger-line"  :class="{ 'hero__burger-line--active': isMobileMenuOpen }"></span>
                    </button>
                </div>
                
            </header>

            <div
                    v-if="isMobileMenuOpen"
                    class="hero__mobile-overlay"
                    @click="closeMobileMenu"
                >

                    <nav class="hero__mobile-nav" @click.stop>
                        <div></div>
                        <ul class="hero__mobile-list">
                            <li v-for="item in menuItems" :key="item.id">
                                <a
                                    :href="item.href"
                                    class="hero__mobile-link"
                                    @click="closeMobileMenu"
                                    >
                                    {{ item.label }}
                                </a>
                            </li>
                        </ul>
                        <button
                        v-if="!isDesktop"
                        class="hero__vk-content-btn"
                        @click="goToVKGroup"
                    >
                        <p>Подписывайтесь на группу Вконтакте</p>
                        <img src="../assets/vector/Vector.svg">
                        
                    </button>
                    </nav>    
                </div>
            

            <div class="hero__content">
                <div class="hero__content-row">
                    <h1 class="hero__title"> Пишем код. Согреваем атмосферой. Создаём твоё портфолио. </h1>
                    <button
                        v-if="isDesktop"
                        class="hero__vk-content-btn"
                        @click="goToVKGroup"
                    >
                        <p>Подписывайтесь на группу Вконтакте</p>
                        <img src="../assets/vector/Vector.svg">
                        
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
    { id: 1, label: 'Преимущества'},
    { id: 2, label: 'Команда'},
    { id: 3, label: 'Плюшки'},
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
.hero-wrapper {
    min-height: 100vh;
    background-color:  rgba(0, 0, 0, 0);
    display: flex;
    align-items: center;
    justify-content: center;

}

.hero {
    width: 100%;
    min-height: calc(100vh - 60px);
    background: linear-gradient(180deg, rgba(0, 0, 0, 0) 0%, rgba(0, 47, 85, 0.7) 100%);
    padding: 20px 40px;
    display: flex;
    flex-direction: column;
    position: relative;
    margin: clamp(10px, 2.5vw, 30px);
    border-radius: 30px;
    box-shadow: 0px 0px 0px 50px #002f55;   
    isolation: isolate;
     
}

.hero__header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0;
    position: relative;
    z-index: 101;

}

.hero__logo {
    /* Пока настройки текста*/
    color: white;
    font-size: 26px;
    font-weight: regular;
}

.hero__nav-desktop {
    display: flex;
    align-items: center;
    gap: 30px;
}

.hero__nav-list {
    display: flex;
    list-style: none;
    gap: 25px;
    margin: 0;
    padding: 0;
}

.hero__nav-link {
    color: white;
    text-decoration: none;
    font-size: 24px;
    font-weight: 500;
    position: relative;
    padding: auto;

}

.hero__register-btn {
    background: #E3953B;
    color: white;
    border: none;
    padding: 25px 15px;
    border-radius: 15px;
    box-sizing: content-box;
    font-size: 24px;
    height: 100%;

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
  animation: fadeIn 0.3s ease;
}

.hero__mobile-nav {
  background: #002f55;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  width: 100%;
  height: 46%;
  padding: 100px 30px 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  animation: slideIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);
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
  font-size: 20px;
  font-weight: 500;
  padding: 15px 20px;
  border-radius: 12px;
  transition: all 0.3s ease;
  display: block;
  position: relative;
  overflow: hidden;
}

.hero__content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: end;
  align-items: center;
  text-align: left;
  padding: 40px 20px;
  position: relative;
  z-index: 1;
}

.hero__content-row {
    display: flex;
    flex-direction: row;
    box-sizing: content-box;
    justify-content: space-between;
    align-items: end;
}

.hero__title {
  font-size: clamp(32px, 5vw, 60px);
  margin: 0 0 20px;
  line-height: 1.2;
  color: white;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  font-weight: regular;
  letter-spacing: -0.5px;
}

.hero__vk-content-btn {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    background: rgba(0, 47, 85, 0.5);
    border: 1px solid #3BB0E3;
    border-radius: 20px;
    color: white;
    font-size: 24px;
    font-weight: 500;
    gap: 10px;
    height: 100%;
    padding: 25px 15px;
    max-width: 400px;
    max-height: 100px;

}

@media (max-width: 950px) {
  .hero-wrapper {
    padding: 15px;
  }
  
  .hero {
    min-height: calc(100vh - 30px);
    padding: 15px 20px;
    border-radius: 20px;
  }
  
  .hero__logo-text {
    font-size: 24px;
  }
  
  .hero__title {
    font-size: clamp(28px, 6vw, 40px);
  }
}


@media (max-width: 480px) {
  .hero-wrapper {
    padding: 10px;
  }
  
  .hero {
    min-height: calc(100vh - 20px);
    padding: 10px 15px;
    border-radius: 15px;
  }
  
  .hero__logo-text {
    font-size: 20px;
  }
  
  .hero__title {
    font-size: 24px;
  }
  
  .hero__subtitle {
    font-size: 14px;
  }
  
  .hero__register-btn--mobile {
    min-width: 160px;
    padding: 12px 30px;
    font-size: 16px;
  }
  
  .hero__mobile-nav {
    width: 85%;
    padding: 80px 20px 20px;
  }
  
  .hero__mobile-link {
    font-size: 18px;
    padding: 12px 16px;
  }
}

</style>