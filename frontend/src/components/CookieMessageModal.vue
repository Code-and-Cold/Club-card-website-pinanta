<template>
    <div v-if="showMessage" class="cookie-message">
        <p class="cookie-message__text">
            Продолжая использовать сайт, вы принимаете использование
            файлов для сохранения настроек и статистики посещений (куки-файлы)
            и метрических программ. Подробнее в <a class="text-anchor">Политике 
            конфиденциальности</a> и <a class="text-anchor">Пользовательском соглашении</a>.
        </p>
        <button class="cookie-message__accept-button" @click="acceptCookies">
            Принять
        </button>
    </div>
</template>

<script setup>
    import { ref, onMounted } from 'vue';
    const showMessage = ref(true)

    const STORAGE_KEY = 'cookie_consent'

    const checkConsent = () => {
        const consent = localStorage.getItem(STORAGE_KEY)
        if (consent) {
            try {
                const data = JSON.parse(consent)

                const isExpired = Date.now() - data.timestamp > 30 * 24 * 60 * 60 * 1000
                if (!isExpired) {
                    showMessage.value = false
                    return
                }
            } catch (e) {
                localStorage.removeItem(STORAGE_KEY)
            } 
        }
        showMessage.value = true
    }

    const saveConsent = (accepted) => {
        const consentData = {
            accepted,
            timestamp: Date.now()
        }
        localStorage.setItem(STORAGE_KEY, JSON.stringify(consentData))
        showMessage.value = false
    }

    const acceptCookies = () => saveConsent() 

    onMounted(() => {
        checkConsent()
    })

</script>

<style scoped>
    .cookie-message {
        position: fixed;
        bottom: clamp(33px, 2vw, 71px);
        right: 0;

        margin-left: clamp(36px, 2vw, 74px);
        margin-right: clamp(36px, 2vw, 74px);
        z-index: 9999;
        padding: 15px 25px;

        background-color: #002F55;

        border: 1px solid #3BB0E3;
        border-radius: 20px;

        max-width: 450px;

        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        gap: 15px;
    }

    .cookie-message__text {
        font-family: 'Inter';
        font-style: normal;
        font-weight: 400;
        font-size: 16px;
        line-height: 120%;

        color: #FFFFFF;
    }

    .text-anchor {
        color: #E3953B;
        
    }

    .text-anchor:hover {
        text-decoration-line: underline;
    }

    .cookie-message__accept-button {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        padding: 10px 25px;
        gap: 10px;

        font-family: 'Inter';
        font-style: normal;
        font-weight: 400;
        font-size: 16px;
        line-height: 120%;

        color: #FFFFFF;

        background: #3BB0E3;
        border-radius: 15px;
        border-width: 0;
    }
</style>