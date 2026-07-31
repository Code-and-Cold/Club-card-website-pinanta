<template>
  <section class="phrase" data-testid="phrase-section">
    <span class="phrase__text" data-testid="phrase-section__phrase">
      Хватит писать учебные лабы «в стол».
      <mark class="phrase__mark" data-testid="phrase-section__highlight">Присоединяйся к нам</mark>,
      чтобы делать реальные проекты в кайфовой компании и прокачать резюме еще до выпуска.
    </span>
  </section>
</template>
<script setup></script>
<style scoped>
* {
  margin: 0;
  padding: 0;
}
.phrase {
  margin: 0;
  /*TODO: сделать padding равный отсутпу текста от границы в блоке hero*/
  padding: 0 20px;
  display: flex;
  background-color: #002f55;
  /*TODO: сделать min-height как: "min-height: calc(100vh - 60px)"*/
  min-height: clamp(470px, 2.5vw, 100vh);
  align-items: center;
  justify-content: center;
}

.phrase__text {
  width: 100%;
  max-width: 1293px;
  min-height: 224px;

  color: white;
  font-size: 25px;
  font-weight: 400;
  line-height: 110%;
}

.phrase__mark {
  color: #3bb0e3;
  background-color: rgba(0, 0, 0, 0);
}

@media (min-width: 1024px) {
  .phrase {
    padding: 285px 73px 265px 74px;
    min-height: 900px;
  }
  .phrase__text {
    font-size: 60px;
    line-height: 117%;
  }
}
</style>
