<script setup>
import { ref, reactive } from 'vue'

const formData = reactive({
  name: '',
  department: '',
  course: '',
  link: '',
  agreement: false,
})

const errors = reactive({
  name: '',
  department: '',
  course: '',
  link: '',
  agreement: false,
})

const errorMessages = {
  name: 'Это поле обязательно к заполнению.',
  department: 'Пожалуйста, выберите высшую школу.',
  link: '',
  course: 'Пожалуйста, выберите курс.',
  agreement: 'Необходимо дать согласие на обработку данных.',
}

const requiredFields = ['name', 'department', 'course', 'agreement']

const message = ref('')
const isLoading = ref(false)
const messageType = ref('')

const email = 'ivan@example.com' // FIXME: Только для текущего API, удалить по изменению

async function submitForm() {
  Object.assign(errors, { name: '', department: '', course: '', link: '', agreement: false })

  let hasError = false

  for (const field of requiredFields) {
    if (!formData[field] || (typeof formData[field] === 'string' && !formData[field].trim())) {
      errors[field] = errorMessages[field]
      hasError = true
    }
  }

  if (hasError) return

  isLoading.value = true
  message.value = 'Отправка...'
  messageType.value = ''

  try {
    const res = await fetch('/api/apply', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      // body: JSON.stringify(formData)
      body: JSON.stringify({ name: formData.name, mail: email }), // FIXME: обновить вместе с API
    })

    if (!res.ok) throw new Error()

    message.value = 'Форма успешно отправлена, скоро с тобой свяжутся!'
    messageType.value = 'success'

    Object.assign(formData, {
      name: '',
      department: '',
      course: '',
      link: '',
      agreement: false,
    })
  } catch {
    message.value = 'Форма не отправлена:('
    messageType.value = 'error'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <section class="section feedback" id="feedback-section">
    <div class="feedback__content">
      <h1 class="title">Хочешь тусить с нами?</h1>

      <div class="feedback__form-wrapper">
        <h2 class="subtitle">
          Заполняй форму ниже. Это бесплатно и ни к чему не обязывает, но точно сделает твою
          студенческую жизнь круче и полезнее.
        </h2>

        <form class="form text text--form" @submit.prevent="submitForm" id="feedbackForm">
          <input
            class="form__field"
            :class="{ 'form__field--danger': errors.name }"
            v-model="formData.name"
            placeholder="ФИО"
          />
          <div v-if="errors.name" class="form__error-wrapper">
            <span class="icon icon--danger"></span>
            <span class="text--danger">{{ errors.name }}</span>
          </div>

          <select
            class="form__field"
            :class="{ 'form__field--danger': errors.department }"
            name="Высшая школа"
            v-model="formData.department"
          >
            <option class="form__select-item" value="" disabled selected>Высшая школа</option>
            <option class="form__select-item" value="ВШСГНиМК">ВШСГНиМК</option>
            <option class="form__select-item" value="ВШИТАС">ВШИТАС</option>
            <option class="form__select-item" value="ВШППиФК">ВШППиФК</option>
            <option class="form__select-item" value="ВИШ">ВИШ</option>
          </select>
          <div v-if="errors.department" class="form__error-wrapper">
            <span class="icon icon--danger"></span>
            <span class="text--danger">{{ errors.name }}</span>
          </div>

          <select
            class="form__field"
            :class="{ 'form__field--danger': errors.course }"
            v-model="formData.course"
          >
            <option class="form__select-item" value="" disabled selected>Курс</option>
            <option class="form__select-item" value="1">1</option>
            <option class="form__select-item" value="2">2</option>
            <option class="form__select-item" value="3">3</option>
            <option class="form__select-item" value="4">4</option>
            <option class="form__select-item" value="5">5</option>
            <option class="form__select-item" value="6">6</option>
          </select>
          <div v-if="errors.course" class="form__error-wrapper">
            <span class="icon icon--danger"></span>
            <span class="text--danger">{{ errors.course }}</span>
          </div>

          <input
            class="form__field"
            :class="{ 'form__field--danger': errors.link }"
            v-model="formData.link"
            placeholder="Страница Вконтакте"
          />
          <div v-if="errors.link" class="form__error-wrapper">
            <span class="icon icon--danger"></span>
            <span class="text--danger">{{ errors.link }}</span>
          </div>

          <label class="form__checkbox-wrapper">
            <label class="checkbox">
              <input type="checkbox" v-model="formData.agreement" />
              <span class="checkmark" :class="{ 'form__field--danger': errors.agreement }"></span>
            </label>

            <span class="">
              Даю <a class="link text--attention" href="https://www.example.com">согласие</a> на
              обработку
              <a class="link text--attention" href="https://www.example.com">персональных данных</a>
            </span>
          </label>
          <div v-if="errors.agreement" class="form__error-wrapper">
            <span class="icon icon--danger"></span>
            <span class="text--danger">{{ errors.agreement }}</span>
          </div>
        </form>

        <button class="button subtitle" type="submit" :disabled="isLoading" form="feedbackForm">
          {{ isLoading ? 'Отправка...' : 'Хочу в клуб!' }}
        </button>

        <div v-if="message" class="form__error-wrapper">
          <span
            class="icon"
            :class="{
              'icon--info': messageType === 'success',
              'icon--danger': messageType === 'error',
            }"
          ></span>
          <span
            class="text"
            :class="{
              'text--form': messageType === 'success',
              'text--danger': messageType === 'error',
            }"
          >
            {{ message }}
          </span>
        </div>
      </div>

      <p class="text">
        Или напиши на почту руководителю: <br />
        <a class="link text--attention" href="mailto:ivan@cold-code.ru">ivan@cold-code.ru</a>
      </p>
    </div>

    <img class="feedback__image" src="/src/assets/images/cheremsha.png" alt="Черемша" />
  </section>
</template>

<style scoped>
@import url('@/assets/styles/checkbox.css');

.text--form {
  color: #3bb0e3;
}

.text--attention {
  color: #e3953b;
}

a {
  text-decoration: none;
}

.link:hover {
  text-decoration: underline;
}

.text--danger {
  color: #d33434;
}

.icon {
  width: 25px;
  height: 25px;
  display: inline-block;

  background-repeat: no-repeat;
  background-position: center;
}

.icon--danger {
  background-image: url('@/assets/vector/red_cross.svg');
}

.icon--info {
  background-image: url('@/assets/vector/blue_checkmark.svg'); /* FIXME: blue_checkmark упоминается в другом месте */
}

.feedback {
  min-height: clamp(
    221px,
    calc((100vh - 667px) / (1024px - 667px) * (300px - 221px) + 221px),
    300px
  );

  background-color: #002f55;

  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  gap: clamp(20px, calc((100vw - 375px) / (950px - 375px) * (50px - 20px) + 20px), 50px);
}

.feedback__content {
  flex: 1;
  min-width: clamp(335px, calc((100vw - 375px) / (950px - 375px) * (550px - 335px) + 335px), 550px);
  max-width: 950px;

  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 15px;
}

.feedback__form-wrapper {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: clamp(25px, calc((100vw - 375px) / (950px - 375px) * (30px - 25px) + 25px), 30px);
}

.form {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 15px;
}

.form__field {
  background-color: #003b6b;
  border-radius: 20px;
  border: 1px solid #3bb0e3;

  padding: 18px 30px;

  appearance: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;

  font-family: inherit;
  font-size: inherit;
  font-weight: inherit;
  line-height: inherit;
  color: inherit;
}

.form__field::placeholder {
  color: #3bb0e3;
}

.form__field--danger {
  border-color: #d33434 !important;
}

.form__error-wrapper {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  gap: 10px;
}

.form__checkbox-wrapper {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  gap: 15px;
}

.feedback__image {
  flex-shrink: 0;

  height: clamp(356px, calc((100vw - 375px) / (950px - 375px) * (730px - 356px) + 335px), 730px);
  width: clamp(335px, calc((100vw - 375px) / (950px - 375px) * (686px - 335px) + 335px), 686px);

  border-radius: 30px;
  border: 0px solid #e3953b;
}

.horizontial-overcompensator {
  padding-top: 10px;
}
</style>
