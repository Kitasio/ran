<template>
    <div class="flex">
        <div v-if="admin" @click="showForm = !showForm" class="p-2 my-5 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить дополнительное образование</div>
    </div>
    <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createEducationPlus', doc.year, doc.text, props.id)">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.year" type="text" placeholder="Год">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.text" type="text" placeholder="Текст">
        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
    </form>
    <div v-if="jsonData.length" class="mb-3 font-nova-bold text-2xl">Дополнительное образование / Повышение квалификации / Стажировки</div>
    <div class="grid grid-cols-1 space-y-5">
        <div v-for="data in jsonData" :key="data.id" @click.alt="deleteFromDb('deleteEducationPlus', data.id.toString())">
            <div class="flex flex-col">
                <div class="itallic text-gray-500 font-nova-semi">{{ data.year }}</div>
                <div class="">{{ data.text }}</div>
            </div>
        </div>
    </div>
</template>

<script setup>
import createRecord from "../../composables/createRecord";
import { defineProps, onMounted, ref } from 'vue'
import getAuth from "../../composables/getAuth";
import readRecords from "../../composables/readRecords";
import deleteRecord from "../../composables/deleteRecord";


const props = defineProps(['id'])
const { writeToDb } = createRecord()
const { admin, checkAuth } = getAuth()
const { deleteFromDb } = deleteRecord()
const { jsonData, getJson } = readRecords()
getJson('getEducationPlus', props.id)
checkAuth()
const doc = ref({
    year: '',
    text: '',
})
const showForm = ref(false)
</script>

<style>

</style>