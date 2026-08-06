import { describe, expect, test, beforeEach, afterEach } from 'vitest'
import { render } from 'vitest-browser-vue'
import { page, userEvent } from 'vitest/browser'

import FeedbackSection from '../src/components/FeedbackSection.vue'
import { fillField } from './utils/form-helpers.ts'

import '../src/assets/styles/global.css'

const viewports = [
  { name: 'mobile', width: 375, height: 900, isMobile: true },
  { name: 'desktop', width: 1440, height: 1024, isMobile: false },
]

const viewportConfigs = [
  {
    name: 'mobile',
    width: 375,
    height: 900,
    locatorsAndStyles: [
			{
				find: (page: any) => page.getByRole('region', { name: /feedback/i }), // Section
				style: { width: '375px', height: '812px', padding: '15px', backgroundColor: '#003B6B' },
			},
			{
				find: (page: any) => page.getByRole('heading', { name: 'Хочешь тусить с нами?' }), // Title
				style: { },
			},
			{
				find: (page: any) => page.getByRole('heading', { level: 3, name: 'Заполняй анкету...' }), // Subtitle
				style: { },
			},
			{
				find: (page: any) => getByPlaceholder('ФИО'), // Text input
				style: { },
			},
			{
				find: (page: any) => page.getByRole('combobox', { name: 'Высшая школа' }), // Select input
				style: { },
			},
			{
				find: (page: any) => page.getByRole('checkbox', { name: /даю согласие/i }), // Checkbox input
				style: { },
			},
			{
				find: (page: any) => page.getByRole('button', { name: 'Вступить в клуб' }), // Submit button
				style: { },
			},
			{
				find: (page: any) => page.getByRole('link', { name: 'ivan@cold-code.ru' }), // Intext link
				style: { },
			},
			{
				find: (page: any) => page.getByAltText('Черемша'), // Image
				style: { },
			},
			{
				find: (page: any) => page.getByTestId('feedback-section__message'), // Form response message
				style: { },
			},
    ]
  },
  {
    name: 'desktop',
    width: 1440,
    height: 1024,
    locatorsAndStyles: []
  },
]

describe('FeedbackSectionLogic', () => {
  let getByTestId: any
  let getByText: any
  let getByPlaceholder: any
  let getByRole: any

	beforeEach(async () => {
		vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ success: true }),
    }))
		
		const viewport = viewports.find((viewport) => viewport.name == 'mobile')
    await page.viewport(viewport.width, viewport.height)
   
    const renderResult = render(FeedbackSection)
    getByTestId = renderResult.getByTestId
    getByText = renderResult.getByText
    getByPlaceholder = renderResult.getByPlaceholder
    getByRole = renderResult.getByRole
	})
	
	afterEach(() => {
    vi.unstubAllGlobals()
  })

	test('submits with all required fields', async () => {
		const user = userEvent.setup()

    const nameInput = getByPlaceholder('ФИО')
		const schoolSelect = getByRole('combobox', { name: 'Высшая школа' })
		const courselSelect = getByTestId('feedback-section__input--course')
		const vkInput = getByTestId('feedback-section__input--vk')
		const checkbox = getByRole('checkbox', { name: /даю согласие/i })

    await user.fill(nameInput, 'Иван Петров')
    await user.selectOptions(schoolSelect, 'ВШ 1')
    await user.selectOptions(courselSelect, '1')
    await user.fill(vkInput, 'https://vk.com/ivan_petrov')
    await user.click(checkbox)

    const submitButton = getByTestId('feedback-section__button--submit')
    await user.click(submitButton)
    
    await expect
      .poll(() => getByText('Форма успешно отправлена, скоро с тобой свяжутся!'), {
        timeout: 2000,
        interval: 100,
      })
      .toBeInTheDocument()

    expect(fetch).toHaveBeenCalledTimes(1)
    expect(fetch).toHaveBeenCalledWith('/api/apply', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: 'Иван Петров',
        department: '1',
        course: '1',
        link: 'https://vk.com/ivan_petrov'
      }),
    })

  })

	const fieldConfigs = [
    {
      name: 'fullname',
      testId: 'feedback-section__input--fullname',
      type: 'input',
      required: true,
      fillValue: 'Иван Петров',
      errorMessage: 'Пожалуйста, введите ФИО',
    },
    {
      name: 'school',
      testId: 'feedback-section__input--school',
      type: 'select',
      required: true,
      fillValue: 'ВШ 1',
      errorMessage: 'Пожалуйста, выберите высшую школу',
    },
    {
      name: 'course',
      testId: 'feedback-section__input--course',
      type: 'select',
      required: true,
      fillValue: '1',
      errorMessage: 'Пожалуйста, выберите курс',
    },
    {
      name: 'vk',
      testId: 'feedback-section__input--vk',
      type: 'input',
      required: false,
      fillValue: 'https://vk.com/ivan_petrov',
      errorMessage: '',
    },
    {
      name: 'agreement',
      testId: 'feedback-section__checkmark',
      type: 'checkbox',
      required: true,
      fillValue: true,
      errorMessage: 'Необходимо дать согласие на обработку данных',
    },
  ]

  fieldConfigs
	.filter(field => field.required)
	.forEach((field) => {
		test(`form requires ${field.name} field`, async () => {
			const user = userEvent.setup()

			for (const otherField of fieldConfigs) {
				if (otherField.name !== field.name) {
					await fillField(user, getByTestId, otherField)
				}
			}

			const submitButton = getByTestId('feedback-section__button--submit')
			await user.click(submitButton)

			expect(fetch).not.toHaveBeenCalled()

			await expect
				.poll(() => getByText(field.errorMessage), {
					timeout: 2000,
					interval: 100,
				})
				.toBeInTheDocument()
		})
	})
	
  test('not clears fields on failed submit', async () => {
		const mockFetch = vi.fn().mockRejectedValue(new Error('Network error'))
    fetch = mockFetch

		const user = userEvent.setup()

    const nameInput = getByTestId('feedback-section__input--fullname')
		const schoolSelect = getByTestId('feedback-section__input--school')
		const courselSelect = getByTestId('feedback-section__input--course')
		const vkInput = getByTestId('feedback-section__input--vk')
		const checkbox = getByTestId('feedback-section__checkmark')

    await user.fill(nameInput, 'Иван Петров')
    await user.selectOptions(schoolSelect, 'ВШ 1')
    await user.selectOptions(courselSelect, '1')
    await user.fill(vkInput, 'https://vk.com/ivan_petrov')
    await user.click(checkbox)

    const submitButton = getByTestId('feedback-section__button--submit')
    await user.click(submitButton)
    
    await expect
      .poll(() => getByText('Форма не отправлена:('), {
        timeout: 2000,
        interval: 100,
      })
      .toBeInTheDocument()
    
    expect(fetch).toHaveBeenCalledTimes(1)
    
    await expect.element(getByTestId('feedback-section__form')).toHaveFormValues({
		  name: 'Иван Петров',
      department: '1',
      course: '1',
      link: 'https://vk.com/ivan_petrov',
      agreement: true
		})
  })
  
  test('clears fields on successful submit', async () => {
		const user = userEvent.setup()

    const nameInput = getByTestId('feedback-section__input--fullname')
		const schoolSelect = getByTestId('feedback-section__input--school')
		const courselSelect = getByTestId('feedback-section__input--course')
		const vkInput = getByTestId('feedback-section__input--vk')
		const checkbox = getByTestId('feedback-section__checkmark')

    await user.fill(nameInput, 'Иван Петров')
    await user.selectOptions(schoolSelect, 'ВШ 1')
    await user.selectOptions(courselSelect, '1')
    await user.fill(vkInput, 'https://vk.com/ivan_petrov')
    await user.click(checkbox)

    const submitButton = getByTestId('feedback-section__button--submit')
    await user.click(submitButton)
    
    await expect
      .poll(() => getByText('Форма успешно отправлена, скоро с тобой свяжутся!'), {
        timeout: 2000,
        interval: 100,
      })
      .toBeInTheDocument()
    
    expect(fetch).toHaveBeenCalledTimes(1)
    expect(fetch).toHaveBeenCalledWith('/api/apply', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: 'Иван Петров',
        department: '1',
        course: '1',
        link: 'https://vk.com/ivan_petrov'
      }),
    })
    
    await expect.element(getByTestId('feedback-section__form')).toHaveFormValues({
		  name: '',
      department: '',
      course: '',
      link: '',
      agreement: false
		})
  })
})

describe.skip('FeedbackSectionRendering', () => {
  viewportConfigs.forEach(({ name, width, height, locatorsAndStyles }) => {
    describe(`${name} viewport`, () => {
      test('all elements visible', async () => {
        await page.viewport(width, height)
        const renderResult = render(FeedbackSection)

        for (const { find } of locatorsAndStyles) {
          const element = find(renderResult)
          await expect.element(element).toBeVisible()
        }
      })

      locatorsAndStyles.forEach(({ find, style }, index) => {
        test(`element ${index + 1} styles`, async () => {
          await page.viewport(width, height)
          const renderResult = render(FeedbackSection)
          
          const element = find(renderResult)
          await expect.element(element).toBeVisible()
          await expect.element(element).toHaveStyle(style)
        })
      })
    })
  })
})

describe.skip('FeedbackSection', () => {
  test.skip('должен корректно рендерить все элементы на мобильном viewport (375px)', async () => {
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
        inputType: 'input',
      },
      {
        name: 'school',
        inputType: 'select',
      },
      {
        name: 'course',
        inputType: 'select',
      },
      {
        name: 'vk',
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

  test.skip('интерактивность формы: чекбокс и кнопка', async () => {
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

  test.skip('изображение загружается без ошибки 404', async () => {
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

  test.skip('все обязательные элементы присутствуют и имеют правильную структуру', async () => {
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
