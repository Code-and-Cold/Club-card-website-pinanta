<script setup></script>

<template>
  <section class="section footer">
    <div class="footer__row footer__top">
      <div class="footer__icons-wrapper">
        <img class="footer__club-icon" src="../assets/vector/club_logo.svg" alt="Лого Клуба" />
        <img class="footer__narfu-icon" src="../assets/vector/narfu_logo.svg" alt="Лого САФУ" />
      </div>
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

  padding-bottom: 55px !important; /* FIXME: Соответствие макету, несоответствие инструкциям в беседе */
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
}

.footer__icons-wrapper {
  width: clamp(
    150px,
    calc((100vw - 375px) / (1440px - 375px) * (364px - 150px) + 150px),
    364px
  ); /* FIXME: 150px взято случайно, в макете нет mobile иконок */
  height: fit-content;

  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
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

.footer__club-icon {
  width: clamp(
    150px,
    calc((100vw - 375px) / (1440px - 375px) * (272px - 150px) + 150px),
    272px
  ); /* FIXME: 150px взято случайно, в макете нет mobile иконок */
}

.footer__narfu-icon {
  width: clamp(
    50px,
    calc((100vw - 375px) / (1440px - 375px) * (82px - 50px) + 50px),
    82px
  ); /* FIXME: 50px взято случайно, в макете нет mobile иконок */
}

.footer__vk-icon {
  width: clamp(40px, calc((100vw - 375px) / (1440px - 375px) * (60px - 40px) + 40px), 60px);
}
</style>
