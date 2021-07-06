<template>
  <div>
      <div class="flex justify-between">
        <h1 class="text-3xl font-nova-bold">Документы</h1>
        <div v-if="username" @click="showForm = !showForm" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить документ</div>
        <div></div>
      </div>

      <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb('createDocument', doc.category, doc.name, doc.format, fileUrl, username)">
        <label for="image">Выберите документ</label>
        <div class="flex m-4">
            <input @change="handleChange($event)" class="mb-5 mr-5" name="file" type="file" placeholder="Документ">
            <div v-if="!file" @click="uploadImage" class="p-2 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Сохранить</div>
            <div v-if="file" @click="uploadImage" class="p-2 transition cursor-pointer rounded-md text-blue-600">Готово</div>
        </div>
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.category" type="text" placeholder="Категория">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Название">
        <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.format" type="text" placeholder="Формат (pdf, docx ...)">
        <div>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
            <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
        </div>
      </form>

      <div class="mt-5 grid lg:grid-cols-4 gap-x-14 text-lg">
          <div class="col-span-4 lg:col-span-3">
              <p>Главные информационные ресурсы института</p>
              <div class="grid grid-cols-1 gap-8 md:gap-5 md:mt-3">
                  <div class="w-full" v-for="document in jsonData" :key="document.id" @click.alt="deleteFromDb('deleteDocument', document.id)">
                    <div class="flex h-40 flex-col md:flex-row md:h-auto justify-between rounded-lg transition duration-500 shadow-brand hover:shadow-bigger p-3">
                        <div class="font-nova-semi" v-html="document.name"></div>
                        <a :href="document.url" download>
                            <div class="text-blue-600">Скачать <span class="uppercase md:ml-5 md:mr-2 md:text-white md:bg-red-600 rounded-full md:py-1 md:px-3">{{ document.format }}</span></div>
                        </a>
                    </div>
                  </div>
              </div>
          </div>
          <div class="mt-10 lg:mt-0 col-span-4 lg:col-span-1 mb-5">
              <div class="grid sm:grid-cols-2 gap-5 lg:grid-cols-1">
                  <SingleNews class="mb-5"></SingleNews>
                  <SingleAd></SingleAd>
              </div>
          </div>
      </div>
  </div>
</template>

<script setup>
import SingleNews from '../components/SingleNews.vue'
import SingleAd from '../components/SingleAd.vue'
import { ref } from 'vue'
import readRecords from "../composables/readRecords"
import createRecord from "../composables/createRecord"
import deleteRecord from "../composables/deleteRecord"
import getAuth from "../composables/getAuth"

const del = (id) => {
    console.log(id)
}

const showForm = ref(false)
const { getJson, jsonData } = readRecords()
const { handleChange, writeToDb, uploadImage, fileUrl } = createRecord()
const { deleteFromDb } = deleteRecord()
const { isAdmin, username, uid, checkAuth } = getAuth()

checkAuth()
getJson('getDocuments')

const doc = ref({
    name: '',
    position: '',
    phone: '',
    email: '',
})

const documents = ref([
    {
        category: "Официальные документы",
        name: "Устав ИДВ РАН",
        format: "pdf",
        url: "/numpy.pdf",
        id: "1"
    },
    {
        category: "Официальные документы",
        name: "Устав ИДВ РАН обновленный",
        format: "pdf",
        url: "/Users/radiofed/Documents/numpy.pdf",
        id: "2"
    },
    {
        category: "Противодействие коррупции",
        name: "Антикоррупционная политика",
        format: "pdf",
        url: "/Users/radiofed/Documents/numpy.pdf",
        id: "3"
    },
])
</script>

<style>

</style>