let siteDataPromise = null

export function getSiteData() {
  if (!siteDataPromise) {
    siteDataPromise = fetch('/api/site').then(async (response) => {
      if (!response.ok) {
        throw new Error(`Failed to load site data: ${response.status}`)
      }

      return response.json()
    })
  }

  return siteDataPromise
}
