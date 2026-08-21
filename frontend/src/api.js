const API_BASE_URL = import.meta.env.VITE_API_URL || '/api'

export const api = {
  async submitFeedback(data) {
    const response = await fetch(`${API_BASE_URL}/apply`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    return response.json()
  },
}
