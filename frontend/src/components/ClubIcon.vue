<script setup>
defineProps({
  variant: {
    type: String,
    default: 'badge',
  },
})
</script>

<template>
  <span class="club-icon" :class="`club-icon--${variant}`" aria-hidden="true">
    <svg v-if="variant === 'badge'" viewBox="0 0 48 48" fill="none">
      <rect x="3" y="3" width="42" height="42" rx="9" stroke="currentColor" stroke-width="3" />
      <path
        d="M14 18.5C14 13.806 17.806 10 22.5 10h3C30.194 10 34 13.806 34 18.5v5.742c0 6.33-4.286 11.857-10.412 13.42a1.8 1.8 0 0 1-.888 0C16.286 36.099 12 30.572 12 24.242V20.5c0-1.105.895-2 2-2Z"
        stroke="currentColor"
        stroke-width="2.5"
      />
      <path
        d="m19 24 3.2 3.2L29 20.4"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>

    <svg v-else viewBox="0 0 48 48" fill="none">
      <path
        d="M13 15 8 8v16c0 8.837 7.163 16 16 16s16-7.163 16-16V8l-5 7"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <path d="M16 15h16v12a8 8 0 1 1-16 0V15Z" stroke="currentColor" stroke-width="3" />
      <path
        d="M20 23h.01M28 23h.01"
        stroke="currentColor"
        stroke-width="4"
        stroke-linecap="round"
      />
      <path
        d="M20 30c1.1 1.1 2.43 1.65 4 1.65S26.9 31.1 28 30"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
      />
    </svg>
  </span>
</template>

<style scoped>
.club-icon {
  display: inline-grid;
  place-items: center;
  flex: 0 0 auto;
  color: var(--accent-blue);
}

.club-icon svg {
  display: block;
  width: 100%;
  height: 100%;
}

.club-icon--badge {
  width: 46px;
  height: 46px;
}

.club-icon--cat {
  width: 42px;
  height: 42px;
}
</style>
