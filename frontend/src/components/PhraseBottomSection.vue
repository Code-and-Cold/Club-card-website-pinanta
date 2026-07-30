<template>
  <section class="bottom-phrase__wrapper">
    <div class="bottom-phrase">
      <p class="bottom-phrase__text">
        Как говорил <mark class="mark--blue">Владимир Путин</mark>: <br />
        <br />
        «Мне не стыдно перед гражданами, которые голосовали за меня дважды на президентских выборах.
        Все эти восемь лет я пахал, как раб на галерах, с утра до ночи.
        <mark class="mark--yellow">Я доволен результатами своей работы</mark>.»
      </p>
    </div>
  </section>
</template>
<script setup></script>
<style scoped>
.bottom-phrase__wrapper {
  max-height: 100vh;
  background-color: rgba(0, 47, 85, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
}

.bottom-phrase {
  width: 100%;
  min-height: calc(100vh - 60px);
  padding: 20px 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  margin: clamp(10px, 2.5vw, 30px);
  border-radius: 30px;
  /*Блок будет находится между белыми блоками,
    так что грницы тени видно быть не должно (наверное)*/
  box-shadow: 0px 0px 0px 50px #f0f0f1;
  isolation: isolate;
  font-weight: 300;
}

.bottom-phrase__text {
  font-size: clamp(25px, 3vw, 50px);
  box-sizing: content-box;
  max-width: 1024px;
  max-height: 430px;
  line-height: 1.2;
  color: white;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  font-weight: 300;
  letter-spacing: -0.5px;
  text-align: left;
}

.mark--blue {
  background-color: rgba(0, 0, 0, 0);
  color: #3bb0e3;
}

.mark--yellow {
  background-color: rgba(0, 0, 0, 0);
  color: #e3953b;
}
</style>
