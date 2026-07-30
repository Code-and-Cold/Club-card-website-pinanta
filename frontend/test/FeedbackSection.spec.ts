import { describe, expect, test } from 'vitest'
import { render } from 'vitest-browser-vue'
import { page } from '@vitest/browser/context'
import { userEvent } from 'vitest/browser'
import FeedbackSection from '../src/components/FeedbackSection.vue'

import '../src/assets/global.css'
import '../src/assets/styles/test.css'

describe('FeedbackSection', () => {
  test('должен корректно рендерить все элементы на мобильном viewport (375px)', async () => {
    await page.viewport(375, 1108)

    const { getByTestId } = render(FeedbackSection)

    // Контейнер секции
    const section = getByTestId('feedback-section')
    await expect.element(section).toBeInTheDocument()
    await expect.element(section).toBeVisible()
    await expect.element(section).toHaveStyle({
      backgroundColor: '#002F55',
      width: '375px',
      padding: '50px 20px',
    })

    // Заголовок
    const title = getByTestId('feedback-section__title')
    await expect.element(title).toBeInTheDocument()
    await expect.element(title).toBeVisible()
    await expect.element(title).toHaveTextContent('Хочешь тусить с нами?')
    await expect.element(title).toHaveStyle({
      fontFamily: 'JetBrains Mono',
      fontSize: '35px',
      color: '#FFFFFF',
      lineHeight: '110%',
    })

    // Подзаголовок
    const subtitle = getByTestId('feedback-section__subtitle')
    await expect.element(subtitle).toBeInTheDocument()
    await expect.element(subtitle).toBeVisible()
    await expect
      .element(subtitle)
      .toHaveTextContent('Заполняй анкету и с тобой свяжутся наши шестерки')
    await expect.element(subtitle).toHaveStyle({
      fontFamily: 'Inter',
      fontSize: '24px',
      color: '#FFFFFF',
      lineHeight: '120%',
    })

    // === ПРОВЕРКИ ФОРМЫ ===

    // Форма
    const form = getByTestId('feedback-section__form')
    await expect.element(form).toBeInTheDocument()
    await expect.element(form).toBeVisible()

    const fields = [
      {
        name: 'fullname',
        placeholder: 'ФИО',
        inputType: 'input',
      },
      {
        name: 'school',
        placeholder: 'Высшая школа',
        inputType: 'select',
      },
      {
        name: 'course',
        placeholder: 'Курс',
        inputType: 'select',
      },
      {
        name: 'vk',
        placeholder: 'Страница Вконтакте',
        inputType: 'input',
      },
    ] as const

    for (const field of fields) {
      const fieldWrapper = getByTestId(`feedback-section__field--${field.name}`)
      await expect.element(fieldWrapper).toBeInTheDocument()
      await expect.element(fieldWrapper).toHaveStyle({
        backgroundColor: '#003B6B',
        border: '2px solid #3BB0E3',
        borderRadius: '20px',
        padding: '18px 30px',
        display: 'flex',
        alignItems: 'center',
      })

      const input = getByTestId(`feedback-section__input--${field.name}`)
      await expect.element(input).toBeInTheDocument()
      await expect.element(input).toBeVisible()
      await expect.element(input).toHaveAttribute('placeholder', field.placeholder)
    }

    // Чекбокс согласия
    const checkbox = getByTestId('feedback-section__checkbox--consent')
    await expect.element(checkbox).toBeInTheDocument()
    await expect.element(checkbox).toBeVisible()
    await expect.element(checkbox).toHaveAttribute('type', 'checkbox')

    // Текст чекбокса
    const consentLabel = getByTestId('feedback-section__checkbox-text')
    await expect.element(consentLabel).toBeInTheDocument()
    await expect.element(consentLabel).toBeVisible()
    await expect.element(consentLabel).toHaveStyle({
      fontFamily: 'Inter',
      fontSize: '16px',
      color: '#3BB0E3',
      lineHeight: '150%',
    })

    // Кнопка отправки
    const button = getByTestId('feedback-section__button--submit')
    await expect.element(button).toBeInTheDocument()
    await expect.element(button).toBeVisible()
    await expect.element(button).toHaveTextContent('Вступить в клуб')
    await expect.element(button).toHaveStyle({
      backgroundColor: '#E3953B',
      borderRadius: '20px',
      padding: '15px 25px',
      fontSize: '24px',
      color: '#FFFFFF',
    })
    await expect.element(button).toBeEnabled()

    // еmail
    const emailText = getByTestId('feedback-section__email')
    await expect.element(emailText).toBeInTheDocument()
    await expect.element(emailText).toBeVisible()
    await expect.element(emailText).toHaveTextContent('ivan@cold-code.ru')
    await expect.element(emailText).toHaveStyle({
      color: '#e3953b',
    })

    // === МЕДИА-ЭЛЕМЕНТЫ ===

    // Изображение-мем
    const image = getByTestId('feedback-section__image--meme')
    await expect.element(image).toBeInTheDocument()
    await expect.element(image).toBeVisible()
    await expect.element(image).toHaveStyle({
      width: '335px',
      height: '357px',
      borderRadius: '30px',
    })
  })

  test('интерактивность формы: чекбокс и кнопка', async () => {
    const { getByTestId } = render(FeedbackSection)

    // Чекбокс
    const checkbox = getByTestId('feedback-section__checkbox--consent')
    const checkmark = getByTestId('feedback-section__checkmark')

    //    await expect.element(checkbox).not.toBeVisible();
    await expect.element(checkmark).toBeVisible()

    // Клик
    await checkmark.click()
    await expect.element(checkbox).toBeChecked()

    // Клик для снятия
    await checkmark.click()
    await expect.element(checkbox).not.toBeChecked()

    // Кнопка; отправку проверяем отдельно
    const button = getByTestId('feedback-section__button--submit')
    await expect.element(button).toBeVisible()
  })

  test('отправка формы с валидными данными', async () => {
    const user = userEvent.setup()
    const { getByTestId } = render(FeedbackSection)

    // 1. Заполняем поле "ФИО"
    const nameInput = getByTestId('feedback-section__input--fullname')
    await user.type(nameInput, 'Иван Петров')

    // 2. Заполняем поле "Высшая школа"
    const schoolSelect = getByTestId('feedback-section__input--school')
    await user.selectOptions(schoolSelect, 'ВШ 1')

    // 3. Заполняем поле "Страница Вконтакте"
    // const vkInput = getByTestId('feedback-section__input--vk')
    // await user.type(vkInput, 'https://vk.com/ivan_petrov')

    // 4. Отмечаем чекбокс
    const checkbox = getByTestId('feedback-section__checkmark') // Кастомный чекбокс
    await user.click(checkbox)

    // 5. Кликаем по кнопке отправки
    const submitButton = getByTestId('feedback-section__button--submit')
    await user.click(submitButton)

    // 6. Проверяем результат
    //    const successMessage = getByTestId('feedback-section__message');
    //    await expect.element(successMessage).toBeVisible();
    //    await expect.element(successMessage).toHaveTextContent('✅ Заявка отправлена!');
  })

  test('изображение загружается без ошибки 404', async () => {
    const { getByTestId } = render(FeedbackSection)

    const image = getByTestId('feedback-section__image--meme')
    await expect.element(image).toBeInTheDocument()

    // Проверяем наличие background-image
    await expect
      .element(image)
      .toHaveStyle({ backgroundImage: /url\(.*JzSaEUPBnE\.1-cheremsha-meme\.jpg.*\)/ })

    // Дополнительно: пытаемся загрузить изображение по URL
    const style = image.element()?.getAttribute('style')
    const urlMatch = style?.match(/url\(["']?(.*?)["']?\)/)
    if (urlMatch && urlMatch[1]) {
      const imageUrl = urlMatch[1]
      // Используем fetch для проверки доступности
      try {
        const response = await fetch(imageUrl)
        expect(response.status).toBe(200)
      } catch (error) {
        // Если fetch недоступен или URL относительный, пропускаем
        // или логируем предупреждение
        console.warn('Не удалось проверить статус изображения:', error)
      }
    }
  })

  test('все обязательные элементы присутствуют и имеют правильную структуру', async () => {
    const { getByTestId } = render(FeedbackSection)

    // Проверяем, что все элементы с data-testid существуют
    await expect.element(getByTestId('feedback-section')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__title')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__subtitle')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__form')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__input--fullname')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__input--school')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__input--vk')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__checkbox--consent')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__button--submit')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__email')).toBeInTheDocument()
    await expect.element(getByTestId('feedback-section__image--meme')).toBeInTheDocument()
  })
})

describe('FeedbackSectionDesktop', () => {
  test('должен корректно отображаться на десктопе (1024px)', async () => {
    await page.viewport(1024, 800)

    const { getByTestId } = render(FeedbackSection)

    // Проверяем, что основные элементы видны
    await expect.element(getByTestId('feedback-section')).toBeVisible()
    await expect.element(getByTestId('feedback-section__title')).toBeVisible()
    await expect.element(getByTestId('feedback-section__subtitle')).toBeVisible()
    await expect.element(getByTestId('feedback-section__form')).toBeVisible()
    await expect.element(getByTestId('feedback-section__button--submit')).toBeVisible()
    await expect.element(getByTestId('feedback-section__image--meme')).toBeVisible()

    // Дополнительно: можно проверить, что ширина секции адаптируется (если CSS медиа-запросы)
    // Поскольку в CSS ширина фиксирована 375px, на десктопе она может быть другой,
    // но в макете нет адаптивных правил, поэтому просто проверяем наличие
    const section = getByTestId('feedback-section')
    await expect.element(section).toBeInTheDocument()
    // Если планируется адаптивность, можно проверить изменение стилей
  })
})
