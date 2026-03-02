<template>
  <ModalDialog
    title="Enter person details"
    :is-open="props.open"
    show-header-close-button
    @header-close-clicked="emit('close')"
  >
    <div class="flex flex-col gap-4">
      <label for="name">Name:</label>
      <input
        ref="nameInput"
        type="text"
        v-model="currentEntry.Name"
        class="min-h-10 w-[20rem] rounded-sm outline-2 outline-black focus:outline-blue-500"
      />
    </div>
    <div class="bg-white mt-6 flex gap-4 justify-end">
      <button
        v-if="!editMode"
        class="bg-green-600 text-white font-semibold px-4 py-2 rounded-md"
        @click="addPerson"
      >
        Add
      </button>
      <button
        v-if="editMode"
        class="border-2 font-semibold px-4 py-2 rounded-md"
        @click="emit('save', currentEntry)"
      >
        Save
      </button>
      <button
        class="bg-red-700 text-white font-semibold px-4 py-2 rounded-md"
        @click="emit('close')"
      >
        Cancel
      </button>
    </div>
  </ModalDialog>
</template>

<script setup lang="ts">
import { Person } from '@/classes/Person'
import ModalDialog from './ModalDialog.vue'
import { ref, useTemplateRef, watch } from 'vue'
import { Utils } from '@/classes/Utils'

const nameInput = useTemplateRef('nameInput')

const props = withDefaults(
  defineProps<{
    open: boolean
    editMode?: boolean
    modelValue: Person
  }>(),
  {
    modelValue: () => new Person(),
  },
)

const currentEntry = ref(new Person())

const emit = defineEmits<{
  (event: 'close'): void
  (event: 'add', value: Person): void
  (event: 'save', value: Person): void
}>()

const addPerson = () => {
  currentEntry.value.id = Utils.generateRandomString(10)
  console.log(currentEntry.value.id)

  emit('add', currentEntry.value)
}

const debounce = Utils.debounce(() => nameInput.value?.focus(), 300)

watch(
  () => props.open,
  (newVal: boolean) => {
    if (newVal) {
      currentEntry.value = new Person(props.modelValue)
      debounce()
    }
  },
)
</script>

<style scoped></style>
