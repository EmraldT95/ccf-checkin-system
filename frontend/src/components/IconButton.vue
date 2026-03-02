<template>
  <button
    class="group gap-2 align-center cursor-pointer"
    :class="{ 'inline-flex': inline, flex: !inline }"
    :title="props.tooltip"
  >
    <component
      :is="props.icon"
      class="icon transition"
      :class="{
        [`group-hover:fill-${getColorValue()}`]: true,
        'hover:scale-125': !props.text,
      }"
    />
    <span class="transition" :class="{ [`group-hover:text-${getColorValue()}`]: true }">
      {{ props.text }}
    </span>
  </button>
</template>

<script setup lang="ts">
import { type FunctionalComponent } from 'vue'

import { Colors } from '@/classes/Colors'

const props = withDefaults(
  defineProps<{
    hoverColor?: Colors | string
    hoverColorWeight?: 50 | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | 950
    tooltip?: string
    text?: string
    icon: FunctionalComponent
    inline?: boolean
  }>(),
  {
    hoverColor: Colors.yellow,
    hoverColorWeight: 500,
  },
)

const getColorValue = () => {
  if (Object.values(Colors).includes(props.hoverColor as Colors)) {
    return `${props.hoverColor}-${props.hoverColorWeight}`
  }

  return `[${props.hoverColor}]`
}
</script>

<style scoped></style>
