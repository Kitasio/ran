<template>
  <div>
      <div class="flex">
        <div class="text-blue-600 border p-3 rounded border-blue-600 transition hover:bg-blue-600 hover:text-white cursor-pointer mr-5" @click="setImg" v-if="admin">Редактировать</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="showDelete = true" v-if="admin && !showDelete">Удалить</div>
        <div class="text-red-600 p-3 mr-5" v-if="admin && showDelete">Вы уверены?</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="deleteFromDb('deleteUnit', $route.params.id)" v-if="admin && showDelete">Да</div>
      </div>
      <div v-for="doc in jsonData" :key="doc.id" class="flex flex-col">
        <form v-if="edit" class="flex flex-col" @submit.prevent="updateToDb('updateUnit', doc.title, doc.name, doc.body, fileUrl, id)">
            <label for="image">Выберите картинку</label>
            <input @change="handleChange" class="mb-5" name="image" type="file" placeholder="Картинка">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
            <input class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.name" type="text" placeholder="Тег">
            <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.body" type="text" placeholder="Текст"></textarea>

            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
        </form>
        </div>
      <div class="mt-5 grid lg:grid-cols-4 gap-x-10">
        <div class="col-span-4 md:col-span-3 mt-2">
          <div v-for="doc in jsonData" :key="doc.id">
              <div class="col-span-4 md:col-span-3 mt-2">
                <div class="font-nova-bold mt-1 text-4xl">{{ doc.title }}</div>
                <div class="mt-10 text-lg" v-html="doc.body"></div>
              </div>
          </div>
          <div class="mt-10">
            TODO create another composable for subrecords to fetch this data so there
            will not be conflicts with readRecords
            <div v-if="admin" @click="showForm = !showForm" class="p-2 mb-5 text-blue-600 border-2 border-blue-600 transition cursor-pointer rounded-md hover:text-white hover:bg-blue-600">Добавить подразделение</div>
            <form v-if="showForm" class="flex flex-col border-b border-blue-600 pb-10" @submit.prevent="writeToDb()">
              <textarea class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="researchText" type="text" placeholder="Текст"></textarea>

              <div>
                  <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                  <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
              </div>
            </form>

            <h1 class="text-2xl font-nova-bold mb-5">Направления исследований</h1>
            <div class="flex justify-center items-center">
              <div class="bg-blue-600 p-2 rounded-full mr-3"></div>
              <p class="text-lg">современные российско-китайские отношения и роль "китайского фактора" в возрождении России</p>
            </div>
          </div>
        </div>
        <div class="mt-10 lg:mt-0 col-span-4 lg:col-span-1 mb-5">
            <div class="grid sm:grid-cols-2 gap-5 lg:grid-cols-1 sticky top-3">
                <SingleNews class="mb-5"></SingleNews>
                <SingleAd class="mb-5"></SingleAd>
            </div>
        </div>
      </div>
  </div>
</template>

<script setup>
import axios from 'axios'
import SingleNews from '../components/SingleNews.vue'
import SingleAd from '../components/SingleAd.vue'
import { onMounted, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import deleteRecord from '../composables/deleteRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'
import updateRecord from '../composables/updateRecord'

const route = useRoute()
const { deleteFromDb } = deleteRecord()
const { admin, checkAuth } = getAuth()
const { getJson, jsonData } = readRecords() 
const { updateToDb, handleChange, fileUrl } = updateRecord()
const id = route.params.id
console.log(id)
checkAuth()

const setImg = () => {
    edit.value = !edit.value
    let i = document.getElementById("imgPath")
    console.log(i.getAttribute("src"))
    fileUrl.value = i.getAttribute("src")
}

getJson('singleUnit', id)

const showForm = ref(false)
const researchText = ref('')

const showDelete = ref(false)
const edit = ref(false)
const singleNews = ref('')
</script>
