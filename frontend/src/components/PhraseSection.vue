<template>
    <section class="phrase">
        <p class="phrase__text">
            Хватит писать учебные лабы «в стол». <mark class="phrase__mark">Присоединяйся к нам</mark>, чтобы делать реальные проекты в кайфовой компании и прокачать резюме еще до выпуска.
        </p>
    </section>
</template>
<script setup></script>
<style scoped>
    .phrase {
        margin: 0;
        padding: 0 30px;
        display: flex;
        background-color: #002f55;
        min-height: clamp(470px, 2.5vw, 100vh);
        align-items: center;
        justify-content: center;
    }    

    .phrase__text {
        color: white;
        font-size: clamp(25px, 2.5vw, 60px);
        max-width: 1290px;
        font-weight: 300;
        min-height: 225px;
    }

    .phrase__mark {
        color: #3BB0E3;
        background-color: rgba(0, 0, 0, 0);
    }
</style>