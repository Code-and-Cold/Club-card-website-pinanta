<script setup></script>

<template>
  <section class="section footer">
    <div class="footer__row footer__top">
      <p>ЛОГО САФУ И КЛУБА</p>
      <img class="footer__vk-icon" src="../assets/vector/ri_vk-fill.svg" alt="Группа в ВК" />
    </div>
    <hr class="footer__divider" />
    <div class="footer__row footer__bottom">
      <p>Политика конфиденциальности</p>
      <p>© Код и Холод 2026</p>
    </div>
  </section>
</template>

<style>
.footer {
  min-height: clamp(221px, 30vh, 300px);

  color: white;
  background-color: #002f55;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 20px;

  padding-top: clamp(30px, calc((100vh - 667px) / (1024px - 667px) * (90px - 30px) + 30px), 90px);
  padding-right: clamp(20px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 75px);
  padding-bottom: clamp(
    40px,
    calc((100vh - 667px) / (1024px - 667px) * (50px - 40px) + 40px),
    50px
  );
  padding-left: clamp(20px, calc((100vw - 375px) / (1440px - 375px) * (75px - 20px) + 20px), 75px);
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
  font-family: Inter;
  font-size: 16px;
  font-weight: 400; /* Regular */
  line-height: 1.05;
  color: white;

  flex-wrap: wrap;
}

.footer__vk-icon {
  width: clamp(40px, calc((100vw - 375px) / (1440px - 375px) * (60px - 40px) + 40px), 60px);
}
</style>
