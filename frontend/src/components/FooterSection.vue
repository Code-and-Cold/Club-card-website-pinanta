<script setup></script>

<template>
  <section class="section footer">
    <div class="footer__row footer__top">
      <p>ЛОГО САФУ И КЛУБА</p>
      <img class="footer__vk-icon" src="../assets/vector/ri_vk-fill.svg" alt="Группа в ВК" />
    </div>
    <hr class="footer__divider" />
    <div class="footer__row footer__bottom text">
      <a class="text" href="https://www.example.com">Политика конфиденциальности</a>
      <p>© Код и Холод 2026</p>
    </div>
  </section>
</template>

<style>
.footer {
  min-height: clamp(
    221px,
    calc((100vh - 667px) / (1024px - 667px) * (300px - 221px) + 221px),
    300px
  );

  background-color: #002f55;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 20px;
}

.footer__row {
  width: 100%;

  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  gap: 25px;
}

.footer__top {
  /* Здесь в макете шрифт, отличный от title, subtitle, text */
  font-family: Inter;
  font-size: clamp(25px, calc((100vw - 375px) / (1440px - 375px) * (35px - 25px) + 25px), 35px);
  font-weight: 400; /* Regular */
  line-height: 1.05;
  color: white;
}

.footer__divider {
  width: 100%;
  height: 1px;
}

.footer__bottom {
  flex-wrap: wrap;
}

@media (min-width: 510px) {
  .footer__bottom {
    flex-direction: row-reverse;
  }
}

.footer__vk-icon {
  width: clamp(40px, calc((100vw - 375px) / (1440px - 375px) * (60px - 40px) + 40px), 60px);
}
</style>
