import { describe, expect, test } from 'vitest'
import { render } from 'vitest-browser-vue'
import { page } from 'vitest/browser'

import PhraseSection from '../src/components/PhraseSection.vue'

import '../src/assets/styles/global.css'

describe('PhraseSection', () => {
  test('mobile viewport rendering (375px)', async () => {
    await page.viewport(375, 500) // Высота в макете - 470, запас для проверки ограничения по высоте

    const { getByTestId } = render(PhraseSection)

    // Контейнер секции
    const section = getByTestId('phrase-section')
    await expect.element(section).toBeInTheDocument()
    await expect.element(section).toBeVisible()
    await expect.element(section).toHaveStyle({
      backgroundColor: '#002F55',
      width: '375px',
      height: '470px',
      paddingLeft: '20px',
      paddingRight: '20px',
    })

    // Фраза
    const phrase = getByTestId('phrase-section__phrase')
    await expect.element(phrase).toBeInTheDocument()
    await expect.element(phrase).toBeVisible()
    await expect
      .element(phrase)
      .toHaveTextContent(
        'Хватит писать учебные лабы «в стол». Присоединяйся к нам, чтобы делать реальные проекты в кайфовой компании и прокачать резюме еще до выпуска.',
      )
    await expect.element(phrase).toHaveStyle({
      width: '335px',
      height: '224px',

      fontFamily: 'JetBrains Mono',
      fontSize: '25px',
      fontWeight: '400',
      color: '#FFFFFF',
      lineHeight: '110%',
    })

    // Выделение
    const highlightedText = getByTestId('phrase-section__highlight')
    await expect.element(highlightedText).toBeInTheDocument()
    await expect.element(highlightedText).toHaveTextContent('Присоединяйся к нам')
    await expect.element(highlightedText).toHaveStyle({
      color: '#3BB0E3',
    })
  })

  test('все обязательные элементы присутствуют и имеют правильную структуру', async () => {
    const { getByTestId } = render(PhraseSection)

    await expect.element(getByTestId('phrase-section')).toBeInTheDocument()
    await expect.element(getByTestId('phrase-section')).toBeVisible()
    await expect.element(getByTestId('phrase-section__phrase')).toBeInTheDocument()
    await expect.element(getByTestId('phrase-section__phrase')).toBeVisible()
  })
})

describe('PhraseSectionDesktop', () => {
  test('desktop viewport rendering (1440px)', async () => {
    await page.viewport(1440, 1024) // Высота в макете - 900, запас для проверки ограничения по высоте

    const { getByTestId } = render(PhraseSection)

    // Контейнер секции
    const phraseSection = getByTestId('phrase-section')
    const phraseElement = phraseSection.element()
    const phraseStyles = window.getComputedStyle(phraseElement)
    const phraseHeight = parseFloat(phraseStyles.height)

    await expect.element(phraseSection).toBeInTheDocument()
    await expect.element(phraseSection).toBeVisible()
    await expect.element(phraseSection).toHaveStyle({
      backgroundColor: '#002F55',
      width: '1440px',
      padding: '285px 73px 265px 74px', // С какого-то перепугу в макете figma на desktop статичный текст без адаптива
    })

    expect(phraseHeight).toBeGreaterThanOrEqual(895)
    expect(phraseHeight).toBeLessThanOrEqual(905)

    // Фраза
    const phrase = getByTestId('phrase-section__phrase')
    const element = phrase.element()
    const styles = window.getComputedStyle(element)
    const height = parseFloat(styles.height)

    await expect.element(phrase).toBeInTheDocument()
    await expect.element(phrase).toBeVisible()
    await expect
      .element(phrase)
      .toHaveTextContent(
        'Хватит писать учебные лабы «в стол». Присоединяйся к нам, чтобы делать реальные проекты в кайфовой компании и прокачать резюме еще до выпуска.',
      )
    await expect.element(phrase).toHaveStyle({
      width: '1293px',

      fontFamily: 'JetBrains Mono',
      fontSize: '60px',
      fontWeight: '400',
      color: '#FFFFFF',
      lineHeight: '117%',
    })

    expect(height).toBeGreaterThanOrEqual(260) // FIXME: Not an actual size
    expect(height).toBeLessThanOrEqual(355)

    // Выделение
    const highlightedText = getByTestId('phrase-section__highlight')
    await expect.element(highlightedText).toBeInTheDocument()
    await expect.element(highlightedText).toHaveTextContent('Присоединяйся к нам')
    await expect.element(highlightedText).toHaveStyle({
      color: '#3BB0E3',
    }) // TODO: Refactor those two big tests to be one with loop
  })

  test('все обязательные элементы присутствуют и имеют правильную структуру', async () => {
    const { getByTestId } = render(PhraseSection)

    await expect.element(getByTestId('phrase-section')).toBeInTheDocument()
    await expect.element(getByTestId('phrase-section')).toBeVisible()
    await expect.element(getByTestId('phrase-section__phrase')).toBeInTheDocument()
    await expect.element(getByTestId('phrase-section__phrase')).toBeVisible()
  })
})
