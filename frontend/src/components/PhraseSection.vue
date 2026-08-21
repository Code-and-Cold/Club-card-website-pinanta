<template>
  <section class="phrase section" data-testid="phrase-section">
    <span class="phrase__text title" data-testid="phrase-section__phrase">
      Устал от сухой теории, заброшенных пет-проектов и мыслей о том, кому ты нужен без опыта?
      <mark class="phrase__mark" data-testid="phrase-section__highlight">Присоединяйся к нам,</mark>
      чтобы решать боевые задачи и собирать портфолио в команде таких же, как ты.
    </span>
  </section>
</template>
<script setup></script>
<style scoped>
.phrase {
  /*TODO: сделать padding равный отсутпу текста от границы в блоке hero*/
  display: flex;
  background-color: #002f55;
  /*TODO: сделать min-height как: "min-height: calc(100vh - 60px)"*/
  align-items: center;
  justify-content: center;
}

.phrase__text {
  width: 100%;
  margin: 15% 0;
}

.phrase__mark {
  color: #3bb0e3;
  background-color: rgba(0, 0, 0, 0);
}
</style>
