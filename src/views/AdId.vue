<template>
  <div>
      <div class="flex">
        <div class="text-blue-600 border p-3 rounded border-blue-600 transition hover:bg-blue-600 hover:text-white cursor-pointer mr-5" @click="setImg" v-if="isAdmin || username == jsonData[0].created_by">Редактировать</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="showDelete = true" v-if="(isAdmin || username == jsonData[0].created_by) && !showDelete">Удалить</div>
        <div class="text-red-600 p-3 mr-5" v-if="username && showDelete">Вы уверены?</div>
        <div class="text-red-600 border p-3 rounded border-red-600 transition hover:bg-red-600 hover:text-white cursor-pointer mr-5" @click="deleteFromDb('deleteAd', $route.params.id)" v-if="username && showDelete">Да</div>
      </div>
      <div v-for="doc in jsonData" :key="doc.id" class="flex flex-col">
        <form v-if="edit" class="flex flex-col" @submit.prevent="updateToDb('updateAd', doc.title, doc.tag, doc.date, doc.time, id)">
            <input v-if="doc.title" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.title" type="text" placeholder="Заголовок">
            <input v-if="doc.tag" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.tag" type="text" placeholder="Тег">
            <input v-if="doc.date" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.date" type="text" placeholder="Дата">
            <input v-if="doc.time" class="p-2 rounded shadow border-2 border-blue-600 ring-offset-2 mb-5" v-model="doc.time" type="text" placeholder="Время проведения">

            <div>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-if="!isLoading">Сохранить</button>
                <button class="text-left p-3 rounded text-blue-600 transition bg-white border-2 border-blue-600 hover:bg-blue-600 hover:text-white" v-else disabled>Загрузка...</button>        
            </div>
        </form>
        </div>
      <div class="mt-5 grid lg:grid-cols-4 gap-x-10">
        <div class="col-span-4 md:col-span-3 mt-2" v-for="doc in jsonData" :key="doc.id">
            <div class="col-span-4 md:col-span-3 mt-2">
                <div class="font-nova-bold mt-1 text-lg" v-html="doc.time"></div>
                <div class="font-nova-bold mt-1 text-lg" v-html="doc.title"></div>
                <div class="text-blue-600 text-xs font-nova-bold mt-2">{{ doc.tag }}</div>
                <div class="my-3 text-sm">{{ doc.date }}</div>
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
import { onMounted, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import deleteRecord from '../composables/deleteRecord'
import getAuth from '../composables/getAuth'
import readRecords from '../composables/readRecords'
import updateRecord from '../composables/updateRecord'

const route = useRoute()
const { deleteFromDb } = deleteRecord(-1)
const { isAdmin, username, uid, checkAuth } = getAuth()
const { getJson, jsonData } = readRecords() 
const { updateToDb, handleChange, fileUrl } = updateRecord()
const id = route.params.id
checkAuth()

const setImg = () => {
    edit.value = !edit.value
    let i = document.getElementById("imgPath")
    console.log(i.getAttribute("src"))
    fileUrl.value = i.getAttribute("src")
}

getJson('singleAd', id)

const showDelete = ref(false)
const edit = ref(false)
const singleNews = ref('')
</script>
