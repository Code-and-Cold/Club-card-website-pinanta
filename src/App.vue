<script setup>
import Team from './components/Team.vue'
import RecentEvents from './components/RecentEvents.vue'

const team = [
  {
    id: 1,
    photo: "src/photos/1.jpg",
    name: 'За Лупкин',
    position: 'Директор',
  },
  {
    id: 2,
    photo: "src/photos/2.jpg",
    name: 'Яша Лава',
    position: 'Уже уволен',
  },
  {
    id: 3,
    photo: "src/photos/3.jpg",
    name: 'Спайдер мент',
    position: 'Офицер',
  },
  {
    id: 4,
    photo: "src/photos/4.png",
    name: 'Пупу Пу',
    position: 'ОАО',
  },
  {
    id: 5,
    photo: "src/photos/5.png",
    name: 'ААААААААААААААА',
    position: 'ААААААААА',
  },
]

const recentEvents = [
  {
    id: 1,
    photo: "src/photos/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 2,
    photo: "src/photos/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 3,
    photo: "src/photos/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
  {
    id: 4,
    photo: "src/photos/6.jpg",
    title: 'Проекты, которыми пользуются',
    text: 'Мы не делаем калькуляторы и todo-моды. Мы берём в работу живые проекты, решаем настоящие проблемы и доводим продукт до релиза.',
    data: '23.07.2026',
  },
]
</script>

<template>
  <main>
    <Team :items="team" />
    <RecentEvents :items="recentEvents" />
  </main>
</template>
