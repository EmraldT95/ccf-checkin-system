<template>
  <div class="flex flex-col gap-2">
    <div class="flex flex-row justify-between items-center">
      <span class="text-2xl font-semibold">Persons</span>
      <button class="group inline-flex gap-2 align-center" @click="openEditDialog">
        <IconAddCircle class="transition group-hover:fill-yellow-500" />
        <span class="transition group-hover:text-yellow-500">Add Person</span>
      </button>
    </div>
    <hr />
    <div :key="entry.id" v-for="entry in people" class="flex flex-row justify-between">
      <div>{{ entry.Name }}</div>
      <div class="inline-flex gap-2">
        <button title="Show details" @click="editPerson(entry)">
          <IconEdit class="transition hover:scale-125" />
        </button>
        <button title="Remove" @click="removePerson(entry.id)">
          <IconClose class="hover:fill-red-500 transition hover:scale-125" />
        </button>
      </div>
    </div>
  </div>
  <EditPersonDialog
    :open="dialogIsOpen"
    :edit-mode="editting"
    :model-value="person"
    @add="addPerson"
    @save="savePersonDetails"
    @close="dialogIsOpen = false"
  />
</template>

<script setup lang="ts">
import { ref, type Ref } from 'vue'

import { Person } from '@/classes/Person'

import IconAddCircle from '~icons/mdi-outlined/add_circle_outline.svg'
import IconClose from '~icons/mdi-outlined/close.svg'
import IconEdit from '~icons/mdi-outlined/edit.svg'
import EditPersonDialog from '@/components/EditPersonDialog.vue'

const people: Ref<Array<Person>> = ref([new Person({ id: '232', Name: 'test' })])
const dialogIsOpen = ref(false)
const editting = ref(false)
const person = ref(new Person())

const openEditDialog = () => {
  person.value = new Person()
  editting.value = false
  dialogIsOpen.value = true
}

const addPerson = (value: Person) => {
  people.value.push(value)
  dialogIsOpen.value = false
}

const editPerson = (entry: Person) => {
  editting.value = true
  person.value = entry
  dialogIsOpen.value = true
}

const savePersonDetails = (value: Person) => {
  for (let i = 0; i < people.value.length; i++) {
    if (people.value[i]?.id === value.id) {
      people.value[i] = value
      break
    }
  }
  dialogIsOpen.value = false
}

const removePerson = (id: string) => {
  for (let i = 0; i < people.value.length; i++) {
    if (people.value[i]?.id === id) {
      people.value.splice(i, 1)
      break
    }
  }
  dialogIsOpen.value = false
}
</script>

<style scoped></style>
