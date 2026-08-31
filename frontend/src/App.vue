<script setup>
import { ref, onMounted } from 'vue'
import { api } from '@/api.js'

import HeroSection from './components/HeroSection.vue'
import FeedbackSection from './components/FeedbackSection.vue'
import PhraseSection from './components/PhraseSection.vue'
import PhraseBottomSection from './components/PhraseBottomSection.vue'
import ClubAdvantagesSection from './components/ClubAdvantagesSection.vue'
import MemberBenefitsSection from './components/MemberBenefitsSection.vue'
import TeamSection from './components/TeamSection.vue'
import RecentEvents from './components/RecentEvents.vue'
import FooterSection from './components/FooterSection.vue'
import CookieMessageModal from './components/CookieMessageModal.vue'

const siteData = ref(null)
const isLoading = ref(true)
const error = ref(null)

onMounted(async () => {
  try {
    const data = await api.getSiteData()
    siteData.value = data
  } catch (err) {
    error.value = err.message
    console.error('Failed to load site data:', err)
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <div v-if="isLoading" class="loading">Загрузка...</div>
  <div v-else-if="error" class="error">Ошибка: {{ error }}</div>
  <div v-else class="main">
    <CookieMessageModal />
    <HeroSection :data="siteData.hero" />
    <PhraseSection :text="siteData.warmup?.text" />
    <ClubAdvantagesSection :cards="siteData.why_us_cards" :title="siteData.why_us?.title" />
    <TeamSection :items="siteData.team_members" :title="siteData.team?.title" />
    <MemberBenefitsSection :items="siteData.benefits_items" :title="siteData.benefits?.title" />
    <PhraseBottomSection :text="siteData.quote?.text" />
    <RecentEvents :items="siteData.news" :title="siteData.news_block?.title" />
    <!--
    ВРЕМЕННО!!!

    Убрали FeedbackSection, воизбежание юридических проблем, пока не готовы документы по 152-ФЗ

    <FeedbackSection
      :schools="siteData.schools"
      :courses="siteData.courses"
      :apply-block="siteData.apply_block"
    />
    -->
    
    <FooterSection />
  </div>
</template>

<style>
@import url('./assets/styles/global.css');

.loading,
.error {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  color: white;
  font-family: Inter, sans-serif;
  font-size: 24px;
}

.error {
  color: #d33434;
}

.main {
  background-image: url('./assets/images/Artboard\ 1.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  background-attachment: fixed;
  overflow-x: hidden;
}

@media (max-width: 1204px) {
  .main {
    background-image: url('./assets/images/Artboard\ 2.png');
  }
}
</style>
