<template>
  <dialog
    :open="props.isOpen"
    class="fixed m-auto max-w-6/10 w-100% overflow-y-auto z-2 rounded-lg shadow-sm"
  >
    <div v-if="title" class="bg-gray-200 p-4 flex flex-row items-center justify-between gap-16">
      <span class="text-lg font-semibold">{{ title }}</span>
      <button v-if="showHeaderCloseButton" title="Close" @click="emit('header-close-clicked')">
        <IconClose class="transition hover:scale-125" />
      </button>
    </div>
    <div class="bg-white p-4">
      <slot />
    </div>
  </dialog>
  <div v-if="isOpen" class="absolute inset-0 bg-white-300/50 backdrop-blur-xs h-full w-full"></div>
</template>

<script setup lang="ts">
import IconClose from '~icons/mdi-outlined/close.svg'

const props = withDefaults(
  defineProps<{
    isOpen: boolean
    showHeaderCloseButton?: boolean
    title?: string
  }>(),
  {
    isOpen: false,
  },
)

const emit = defineEmits(['header-close-clicked'])
</script>

<style scoped></style>
