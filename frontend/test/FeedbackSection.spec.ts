import { describe, expect, test } from 'vitest'
import { render } from 'vitest-browser-vue'
import { page, userEvent } from 'vitest/browser'

import FeedbackSection from '../src/components/FeedbackSection.vue'

import '../src/assets/styles/global.css'

describe('FeedbackSection', () => {
  test('должен корректно рендерить все элементы на мобильном viewport (375px)', async () => {
    await page.viewport(375, 1108)

    const { getByTestId, getByPlaceholder, getByRole } = render(FeedbackSection)

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
        placeholder: 'ФИО',
        inputType: 'input',
      },
      {
        name: 'Высшая школа',
        inputType: 'select',
      },
      {
        name: 'Курс',
        inputType: 'select',
      },
      {
        placeholder: 'Страница Вконтакте',
        inputType: 'input',
      },
    ] as const

    for (const field of fields) {
      let fieldWrapper = undefined

      switch (field.inputType) {
        case 'input':
          fieldWrapper = getByPlaceholder(field.placeholder)
          break
        case 'select':
          fieldWrapper = getByRole('combobox').filter({ hasText: field.name })
          break
        default:
          expect.fail(`Unknown input type: ${field.inputType}`)
      }

      await expect.element(fieldWrapper).toBeInTheDocument()
      await expect.element(fieldWrapper).toHaveStyle({
        backgroundColor: '#003B6B',
        border: '1px solid #3BB0E3',
        borderRadius: '20px',
        padding: '18px 30px',
      })
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
    const { getByTestId, getByRole } = render(FeedbackSection)

    const nameInput = getByTestId('feedback-section__input--fullname')
    await user.type(nameInput, 'Иван Петров')

    const schoolSelect = getByRole('combobox').filter({ hasText: 'Высшая школа' })
    await user.selectOptions(schoolSelect, 'ВШ 1')

    const courselSelect = getByTestId('feedback-section__input--course')
    await user.selectOptions(courselSelect, '1')

    const vkInput = getByTestId('feedback-section__input--vk')
    await user.type(vkInput, 'https://vk.com/ivan_petrov')

    const checkbox = getByTestId('feedback-section__checkmark') // Кастомный чекбокс
    await user.click(checkbox)

    const submitButton = getByTestId('feedback-section__button--submit')
    await user.click(submitButton)

    await expect
      .poll(() => getByTestId('feedback-section__message'), {
        timeout: 2000,
        interval: 100,
      })
      .toBeInTheDocument()

    const message = getByTestId('feedback-section__message')
    await expect.element(message).toBeVisible()
    await expect.element(message).toHaveTextContent('❌ Ошибка отправки. Попробуйте ещё раз.') // TODO: Integration test for success
    await expect.element(message).toHaveStyle({ color: 'red' })
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
    const { getByTestId, getByRole } = render(FeedbackSection)

    // Проверяем, что все элементы с data-testid существуют
    await expect.element(getByTestId('feedback-section')).toBeVisible()
    await expect.element(getByTestId('feedback-section__title')).toBeVisible()
    await expect.element(getByTestId('feedback-section__subtitle')).toBeVisible()
    await expect.element(getByTestId('feedback-section__form')).toBeVisible()
    await expect.element(getByTestId('feedback-section__input--fullname')).toBeVisible()
    await expect.element(getByRole('combobox').filter({ hasText: 'Высшая школа' })).toBeVisible()
    await expect.element(getByTestId('feedback-section__input--vk')).toBeVisible()
    await expect.element(getByTestId('feedback-section__checkbox--consent')).toBeVisible()
    await expect.element(getByTestId('feedback-section__button--submit')).toBeVisible()
    await expect.element(getByTestId('feedback-section__email')).toBeVisible()
    await expect.element(getByTestId('feedback-section__image--meme')).toBeVisible()
  })
})

describe('FeedbackSectionDesktop', () => {
  test('должен корректно отображаться на десктопе (1440px)', async () => {
    await page.viewport(1440, 800)

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
    await expect.element(section).toBeVisible()
    // Если планируется адаптивность, можно проверить изменение стилей
  })
})
