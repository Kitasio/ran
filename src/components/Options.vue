<template>
  <div v-if="edit">
    <form class="flex flex-col" @submit.prevent="update">
        <label for="image">Выберите картинку</label>
        <!-- <input @change="handleChange" class="mb-5" name="image" type="file" placeholder="Картинка"> -->
        <input v-if="doc.title" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
        <input v-if="doc.tag" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.tag" type="text" placeholder="Тег">
        <input v-if="doc.date" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.date" type="text" placeholder="Дата">
        <textarea v-if="doc.body" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.body" type="text" placeholder="Текст"></textarea>

        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
    </form>
  </div>
</template>

<script setup>
import { defineProps, defineEmit } from 'vue'
import deleteRecord from "../composables/deleteRecord"

const props = defineProps({
    doc: 'doc',
    edit: 'edit',
})
const { deleteFromDb } = deleteRecord()
</script>

<style>

</style>