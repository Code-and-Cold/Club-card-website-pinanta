import { describe, expect, test } from 'vitest'
import { render } from 'vitest-browser-vue'
import { page } from 'vitest/browser'

import HeroSection from '../src/components/HeroSection.vue'

import '../src/assets/styles/global.css'

const viewports = [
  { name: 'mobile', width: 375, height: 900, isMobile: true },
  { name: 'desktop', width: 1440, height: 1024, isMobile: false },
]

describe('HeroSectionLogic', () => {
  const burgerMenuItems = ['Преимущества', 'Команда', 'Плюшки']

  test('burger menu opens and closes on click', async () => {
    const viewport = viewports.find((viewport) => viewport.name == 'mobile')
    await page.viewport(viewport.width, viewport.height)
    const { getByTestId, getByText } = render(HeroSection)

    await expect.element(getByTestId('hero-section__burger-menu')).not.toBeInTheDocument()
    for (const item of burgerMenuItems) {
      await expect.element(getByText(item)).not.toBeInTheDocument()
    }

    const burgerButton = getByTestId('hero-section__burger')

    await expect.element(burgerButton).toBeVisible()
    await burgerButton.click()

    const burgerMenu = await getByTestId('hero-section__burger-menu')
    await expect.element(burgerMenu).toBeVisible()

    for (const item of burgerMenuItems) {
      await expect.element(getByText(item)).toBeVisible()
    }

    await expect.element(burgerButton).toBeVisible()
    await burgerButton.click()

    await expect.element(getByTestId('hero-section__burger-menu')).not.toBeInTheDocument()
  })

  test('navigation links work', async () => {
    const viewport = viewports.find((viewport) => viewport.name == 'mobile')
    await page.viewport(viewport.width, viewport.height)
    const { getByTestId, getByText } = render(HeroSection)

    // TODO: Implement proper redirection check
    for (const item of burgerMenuItems) {
      await expect.element(getByTestId('hero-section__burger-menu')).not.toBeInTheDocument()
      await getByTestId('hero-section__burger').click()
      await expect.element(getByTestId('hero-section__burger-menu')).toBeVisible()
      await getByText(item).click()
      await expect.element(getByTestId('hero-section__burger-menu')).not.toBeInTheDocument()
    }
  })
})

describe('HeroSectionRendering', () => {
  viewports.forEach(({ name, width, height, isMobile }) => {
    test(`${name} viewport rendering (${width}px)`, async () => {
      await page.viewport(width, height)

      const { getByTestId } = render(HeroSection)

      await expect.element(getByTestId('hero-section')).toBeVisible()
      await expect.element(getByTestId('hero-section__header')).toBeVisible()
      await expect.element(getByTestId('hero-section__logo--club')).toBeVisible()
      await expect.element(getByTestId('hero-section__logo--vk')).toBeVisible()
      await expect.element(getByTestId('hero-section__title')).toBeVisible()
      await expect.element(getByTestId('hero-section__button--cta')).toBeVisible()
      await expect.element(getByTestId('hero-section__burger-menu')).not.toBeInTheDocument() // v-if hides it in both modes

      if (isMobile) {
        await expect.element(getByTestId('hero-section__burger')).toBeVisible()
        await expect.element(getByTestId('hero-section__nav')).not.toBeInTheDocument() // Desktop links at the top
      } else {
        await expect.element(getByTestId('hero-section__burger')).not.toBeInTheDocument()
        await expect.element(getByTestId('hero-section__nav')).toBeVisible()
      }
    })
  })
})

describe('mobile viewport styling', () => {
  const viewport = viewports.find((v) => v.name === 'mobile')
  const heroSectionElements = [
    {
      testID: 'hero-section',
      style: {
        width: '375px',
        height: '812px',

        padding: '15px',

        backgroundColor: '#003B6B',
      },
    },
    {
      testID: 'hero-section__header',
      style: {
        width: '375px',
        height: '75px',

        paddingTop: '20px',
        paddingBottom: '20px',
        gap: '23px', // spacing из фигмы
      },
    },
    {
      testID: 'hero-section__logo--vk',
      style: {
        width: '40px',
        height: '40px',
      },
    },
    {
      testID: 'hero-section__title',
      style: {
        // Auto height
        width: '305px',
        paddingLeft: '20px',
        paddingRight: '20px',

        fontFamily: 'JetBrains Mono',
        fontSize: '35px',
        color: '#FFFFFF',
        lineHeight: '110%',
      },
    },
    {
      testID: 'hero-section__button--cta',
      style: {
        // Hug content properties override height and width
        padding: '25px 15px',

        backgroundColor: '#E3953B',

        fontFamily: 'JetBrains Mono',
        fontSize: '35px',
        color: '#FFFFFF',
        lineHeight: '110%',
      },
    },
    {
      testID: 'hero-section__burger', // Button
      style: {
        height: '35px',
        width: '35px',
      },
    },
  ]

  const burgerMenuElements = [
    {
      testID: 'hero-section__burger-menu',
      style: {
        // Auto height
        width: '375px',
        paddingTop: '15px', // Там (в макете) какие-то очень страшные padding-и с других сторон

        backgroundColor: '#003B6B',
      },
    },
    {
      testID: 'hero-section__burger-nav', // Список ссылок (преимущества, команда..)
      style: {
        // Hug content properties override height and width
        gap: '40px',
        padding: '0px',

        fontFamily: 'Inter',
        fontSize: '22px',
        color: '#FFFFFF',
        lineHeight: '105%',
      },
    },
    {
      testID: 'hero-section__burger-vk',
      style: {
        // Hug content properties override height and width
        // padding: '11px 19px' // на глаз, там страшные цифры, возможно из-за статических размеров контейнера с текстом

        borderRadius: '15px',

        backgroundColor: '#003B6B',

        fontFamily: 'Inter',
        fontSize: '22px',
        color: '#FFFFFF',
        lineHeight: '120%',
      },
    },
    {
      testID: 'hero-section__burger-vk-icon',
      style: {
        height: '45px',
        width: '45px',
      },
    },
  ]

  heroSectionElements.forEach(({ testID, style }) => {
    test(`HeroSection: ${testID} styled`, async () => {
      await page.viewport(viewport.width, viewport.height)
      const { getByTestId } = render(HeroSection)

      const element = getByTestId(testID)
      await expect.element(element).toHaveStyle(style)
    })
  })

  burgerMenuElements.forEach(({ testID, style }) => {
    test(`Burger menu: ${testID} styled`, async () => {
      await page.viewport(viewport.width, viewport.height)
      const { getByTestId } = render(HeroSection)

      const burgerButton = getByTestId('hero-section__burger')
      await expect.element(burgerButton).toBeVisible()
      await burgerButton.click()

      const element = getByTestId(testID)
      await expect.element(element).toHaveStyle(style)
    })
  })
})
