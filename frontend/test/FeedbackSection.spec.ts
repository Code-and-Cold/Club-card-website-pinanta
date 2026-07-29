import { expect, test } from 'vitest'
import { render } from 'vitest-browser-vue'
import FeedbackSection from '../src/components/FeedbackSection.vue'

test('renders name and the counter', async () => {
  const { getByText, _getByRole } = render(FeedbackSection, {})

  await expect.element(getByText('Хочешь тусить с нами?')).toBeInTheDocument()
  await expect.element(getByText('Заполняй анкету и с тобой свяжутся наши шестерки')).toBeInTheDocument()
})
