const ASSET_URL_KEYS = new Set(['photo_url', 'image_url', 'icon_url'])

/** Приводит /assets/... к относительному пути (без хоста :8080 из docker-dev). */
export function resolveAssetUrl(url) {
  if (!url) return ''
  if (url.startsWith('/assets/')) return url

  try {
    const { pathname } = new URL(url, window.location.origin)
    if (pathname.startsWith('/assets/')) return pathname
  } catch {
    // не URL — оставляем как есть
  }

  return url
}

export function normalizeSiteData(data) {
  if (Array.isArray(data)) {
    return data.map(normalizeSiteData)
  }

  if (data && typeof data === 'object') {
    return Object.fromEntries(
      Object.entries(data).map(([key, value]) => {
        if (ASSET_URL_KEYS.has(key) && typeof value === 'string') {
          return [key, resolveAssetUrl(value)]
        }
        return [key, normalizeSiteData(value)]
      }),
    )
  }

  return data
}
