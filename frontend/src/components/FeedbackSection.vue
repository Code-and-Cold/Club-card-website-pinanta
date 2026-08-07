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
const messageColor = ref('')

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
  messageColor.value = '#3bb0e3'

  try {
    const res = await fetch('/api/apply', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      // body: JSON.stringify(formData)
      body: JSON.stringify({ name: formData.name, mail: email }), // FIXME: обновить вместе с API
    })

    if (!res.ok) throw new Error()

    message.value = '✅ Заявка отправлена! С вами свяжутся'
    messageColor.value = 'green'

    Object.assign(formData, {
      name: '',
      department: '',
      course: '',
      link: '',
      agreement: false,
    })
  } catch {
    message.value = '❌ Ошибка отправки. Попробуйте ещё раз.'
    messageColor.value = 'red'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <section class="section feedback-section" data-testid="feedback-section">
    <div class="feedback-section__content">
      <h2 class="feedback-section__title" data-testid="feedback-section__title">
        Хочешь тусить с нами?
      </h2>

      <h3 class="feedback-section__subtitle" data-testid="feedback-section__subtitle">
        Заполняй анкету и с тобой свяжутся наши шестерки
      </h3>

      <form
        class="form feedback-section__form"
        data-testid="feedback-section__form"
        @submit.prevent="submitForm"
      >
        <div class="form__field" data-testid="feedback-section__field--fullname">
          <input
            class="form__input"
            :class="{ 'form__field--error': errors.name }"
            data-testid="feedback-section__input--fullname"
            v-model="formData.name"
            placeholder="ФИО"
          />

          <div
            v-if="errors.name"
            class="form__field-error"
            data-testid="feedback-section__field-error--fullname"
          >
            <span class="form__field-error-icon form__field-error-icon--error"></span>
            <span class="form__field-error-text">{{ errors.name }}</span>
          </div>
        </div>

        <div class="form__field" data-testid="feedback-section__field--school">
          <select
            class="form__select"
            :class="{ 'form__field--error': errors.department }"
            name="Высшая школа"
            v-model="formData.department"
          >
            <option class="form__select-item" value="" disabled selected>Высшая школа</option>
            <option class="form__select-item" value="1">ВШ 1</option>
            <option class="form__select-item" value="2">ВШ 2</option>
            <option class="form__select-item" value="3">ВШ 3</option>
          </select>

          <div
            v-if="errors.department"
            class="form__field-error"
            data-testid="feedback-section__field-error--department"
          >
            <span class="form__field-error-icon form__field-error-icon--error"></span>
            <span class="form__field-error-text">{{ errors.department }}</span>
          </div>
        </div>

        <div class="form__field" data-testid="feedback-section__field--course">
          <select
            class="form__select"
            :class="{ 'form__field--error': errors.course }"
            data-testid="feedback-section__input--course"
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

          <div
            v-if="errors.course"
            class="form__field-error"
            data-testid="feedback-section__field-error--course"
          >
            <span class="form__field-error-icon form__field-error-icon--error"></span>
            <span class="form__field-error-text">{{ errors.course }}</span>
          </div>
        </div>

        <div class="form__field" data-testid="feedback-section__field--vk">
          <input
            class="form__input"
            data-testid="feedback-section__input--vk"
            type="url"
            v-model="formData.link"
            placeholder="Страница Вконтакте"
          />

          <div
            v-if="errors.link"
            class="form__field-error"
            :class="{ 'form__field--error': errors.link }"
            data-testid="feedback-section__field-error--link"
          >
            <span class="form__field-error-icon form__field-error-icon--error"></span>
            <span class="form__field-error-text">{{ errors.link }}</span>
          </div>
        </div>

        <div class="form__checkbox-panel">
          <label class="form__checkbox-label" data-testid="feedback-section__checkbox-text">
            <input
              class="form__checkbox-input"
              data-testid="feedback-section__checkbox--consent"
              type="checkbox"
              v-model="formData.agreement"
            />
            <span
              class="form__checkmark"
              :class="{ 'form__field--error': errors.agreement && !formData.agreement }"
              data-testid="feedback-section__checkmark"
            ></span>
            <span class="form__checkbox-text">
              Даю <a class="form__link" href="https://www.example.com">согласие</a> на обработку
              <a class="form__link" href="https://www.example.com">персональных данных</a>
            </span>
          </label>

          <div
            v-if="errors.agreement"
            class="form__field-error"
            data-testid="feedback-section__field-error--agreement"
          >
            <span class="form__field-error-icon form__field-error-icon--error"></span>
            <span class="form__field-error-text">{{ errors.agreement }}</span>
          </div>
        </div>

        <button
          class="form__submit-button"
          data-testid="feedback-section__button--submit"
          type="submit"
          :disabled="isLoading"
        >
          {{ isLoading ? 'Отправка...' : 'Вступить в клуб' }}
        </button>

        <p
          v-if="message"
          class="form__message"
          data-testid="feedback-section__message"
          :style="{ color: messageColor }"
        >
          {{ message }}
        </p>
      </form>

      <p class="feedback-section__footer">
        Или напиши на почту руководителю: <br />
        <a
          class="feedback-section__link"
          data-testid="feedback-section__email"
          href="mailto:ivan@cold-code.ru"
          >ivan@cold-code.ru</a
        >
      </p>
    </div>

    <div class="feedback-section__media">
      <img
        class="feedback-section__image"
        data-testid="feedback-section__image--meme"
        src="/src/assets/images/cheremsha.png"
        alt="Черемша"
      />
    </div>
  </section>
</template>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;

  /* TODO: шрифты не везде такие, вынести по готовности тестов */
  font-family: 'JetBrains Mono';
  font-style: normal;
}

.feedback-section {
  background-color: #002f55;
  padding: 50px 20px;

  display: flex;
  flex-direction: column;
  gap: 30px;
}

.feedback-section__content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
}

.feedback-section__title {
  font-family: JetBrains Mono;
  font-size: 35px;
  font-weight: 400; /* Regular */
  line-height: 1.1;
  color: white;
}

.feedback-section__subtitle {
  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: white;
}

.feedback-section__form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form__field {
  display: flex;
  width: 100%;
  flex-direction: column;
  align-items: stretch;
  gap: 8px;
  padding: 0px 0px;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400;
  line-height: 1.2;
  color: white;

  transition:
    background-color 0.2s,
    border-color 0.2s;
}

.form__field-error {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-left: 0px;

  font-family: Inter;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.2;

  color: #d33434;
}

.form__field--error {
  border-color: #d33434 !important; /* FIXME */
}

.form__field-error-icon {
  width: 24px;
  height: 24px;
}

.form__field-error-icon--success {
  color: #51cf66;
}

.form__field-error-icon--error {
  background-image: url('@/assets/vector/red_cross.svg');
  background-repeat: no-repeat;
  background-position: center;
  padding-right: 0px;
}

.form__field-error-text {
  flex: 1;
}

.form__input {
  flex: 1;
  width: 100%;
  height: 100%;

  background-color: #003b6b;
  border-radius: 20px;
  border: 1px solid #3bb0e3;

  padding: 18px 30px;

  color: white;
  font: inherit;
}

.form__select {
  flex: 1;
  width: 100%;
  height: 100%;

  background-color: #003b6b;
  border-radius: 20px;
  border: 1px solid #3bb0e3;

  padding: 18px 30px;

  color: #3bb0e3;
  font: inherit;

  cursor: pointer;

  appearance: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;

  background-image: url('data:image/svg+xml,...'); /* кастомная стрелка */
  cursor: pointer;
}

.form__select-item {
  flex: 1;
  width: 100%;
  height: 100%;
  background-color: #003b6b;
}

input {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  user-select: none;
}

.form__input::placeholder {
  color: #3bb0e3; /* Forms placeholder font color */
  opacity: 1;
}

.form__checkbox-panel {
  margin-top: -5px; /* FIXME: margin не рекомендуется */
  width: 100%;

  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.form__checkbox-label {
  flex: 1;
  width: 100%;
  height: 100%;

  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 15px;

  font-family: Inter;
  font-size: 16px;
  font-weight: 400; /* Regular */
  line-height: 1.5;
  color: #3bb0e3;
}

.form__checkbox-input {
  position: absolute;
  width: 1px;
  height: 1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  border: 0;
}

.form__link {
  color: #3bb0e3;
}

@media (hover: hover) {
  .form__link:hover {
    color: #5cc5f5;
  }
}

.form__checkmark {
  flex-shrink: 0;
  width: 35px;
  height: 35px;
  border-radius: 10px;
  border: 1px solid #3bb0e3;
  background-color: #003b6b;
  background-repeat: no-repeat;
  background-position: center;
  background-size: 20px 20px;
  transition:
    background-color 0.2s,
    border-color 0.2s;
}

/* On mouse-over, add a grey background color; desktop only */
@media (hover: hover) {
  .form__checkbox-label:hover .form__checkmark {
    background-color: rgba(59, 176, 227, 0.3);
  }

  .form__checkbox-label:hover .form__checkbox-input:checked + .form__checkmark {
    background-color: #1a7fc9;
    border-color: #1a7fc9;
  }
}

.form__checkbox-input:not(:checked) + .form__checkmark {
  background-image: none;
}

.form__checkbox-input:checked + .form__checkmark {
  background-image: url('@/assets/vector/checkmark.svg');

  border-color: #2196f3;
}

.form__checkbox-text {
  flex: 1;
}

.form__submit-button {
  width: fit-content;
  height: fit-content;

  border-radius: 20px;
  border: 0px solid #e3953b;

  padding: 15px 25px;

  background-color: #e3953b;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: white;
}

.feedback-section__footer {
  height: fit-content;

  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: white;
}

.feedback-section__link {
  font-family: Inter;
  font-size: 24px;
  font-weight: 400; /* Regular */
  line-height: 1.2;
  color: #e3953b;
}

.feedback-section__media {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 30px;
}

.feedback-section__image {
  width: 335px;
  border-radius: 30px;
  border: 0px solid #e3953b;
}
</style>
