import { beforeAll } from 'vitest';

// ✅ Просто добавляем CSS через document (без page)
beforeAll(() => {
  // Скрываем скроллбар через DOM-стили
  const style = document.createElement('style');
  style.id = 'test-scrollbar-hide';
  style.textContent = `
    ::-webkit-scrollbar {
      width: 0px !important;
      height: 0px !important;
      background: transparent !important;
    }
    * {
      scrollbar-width: none !important;
    }
    body {
      margin: 0 !important;
      padding: 0 !important;
      overflow: hidden !important;
    }
  `;
  document.head.appendChild(style);
});
